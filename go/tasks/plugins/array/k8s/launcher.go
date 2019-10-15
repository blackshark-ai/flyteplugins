package k8s

import (
	"context"
	"fmt"
	arrayCore "github.com/lyft/flyteplugins/go/tasks/plugins/array/core"
	"strconv"
	"strings"

	arraystatus2 "github.com/lyft/flyteplugins/go/tasks/plugins/array/arraystatus"
	errors2 "github.com/lyft/flytestdlib/errors"

	"github.com/lyft/flytestdlib/logger"

	"github.com/lyft/flyteplugins/go/tasks/pluginmachinery/utils"

	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/lyft/flyteplugins/go/tasks/pluginmachinery/core"
)

const (
	ErrBuildPodTemplate       errors2.ErrorCode = "POD_TEMPLATE_FAILED"
	ErrReplaceCmdTemplate     errors2.ErrorCode = "CMD_TEMPLATE_FAILED"
	ErrSubmitJob              errors2.ErrorCode = "SUBMIT_JOB_FAILED"
	JobIndexVarName           string            = "BATCH_JOB_ARRAY_INDEX_VAR_NAME"
	FlyteK8sArrayIndexVarName string            = "FLYTE_K8S_ARRAY_INDEX"
)

var arrayJobEnvVars = []corev1.EnvVar{
	{
		Name:  JobIndexVarName,
		Value: FlyteK8sArrayIndexVarName,
	},
}

func formatSubTaskName(_ context.Context, parentName, suffix string) (subTaskName string) {
	return fmt.Sprintf("%v-%v", parentName, suffix)
}

func ApplyPodPolicies(_ context.Context, cfg *Config, pod *corev1.Pod) *corev1.Pod {
	if len(cfg.DefaultScheduler) > 0 {
		pod.Spec.SchedulerName = cfg.DefaultScheduler
	}

	return pod
}

// Launches subtasks
func LaunchSubTasks(ctx context.Context, tCtx core.TaskExecutionContext, kubeClient core.KubeClient,
	config *Config, currentState *arrayCore.State) (newState *arrayCore.State, err error) {
	podTemplate, _, err := FlyteArrayJobToK8sPodTemplate(ctx, tCtx)
	if err != nil {
		return currentState, errors2.Wrapf(ErrBuildPodTemplate, err, "Failed to convert task template to a pod template for task")
	}

	var command []string
	if len(podTemplate.Spec.Containers) > 0 {
		command = append(podTemplate.Spec.Containers[0].Command, podTemplate.Spec.Containers[0].Args...)
		podTemplate.Spec.Containers[0].Args = []string{}
	}

	size := currentState.GetExecutionArraySize()
	// TODO: Respect parallelism param
	for i := 0; i < size; i++ {
		pod := podTemplate.DeepCopy()
		indexStr := strconv.Itoa(i)
		pod.Name = formatSubTaskName(ctx, tCtx.TaskExecutionMetadata().GetTaskExecutionID().GetGeneratedName(), indexStr)
		pod.Spec.Containers[0].Env = append(pod.Spec.Containers[0].Env, corev1.EnvVar{
			Name:  FlyteK8sArrayIndexVarName,
			Value: indexStr,
		})

		pod.Spec.Containers[0].Env = append(pod.Spec.Containers[0].Env, arrayJobEnvVars...)

		pod.Spec.Containers[0].Command, err = utils.ReplaceTemplateCommandArgs(ctx, command, arrayJobInputReader{tCtx.InputReader()}, tCtx.OutputWriter())
		if err != nil {
			return currentState, errors2.Wrapf(ErrReplaceCmdTemplate, err, "Failed to replace cmd args")
		}

		pod = ApplyPodPolicies(ctx, config, pod)

		err = kubeClient.GetClient().Create(ctx, pod)
		if err != nil && !k8serrors.IsAlreadyExists(err) {
			if k8serrors.IsForbidden(err) {
				if strings.Contains(err.Error(), "exceeded quota") {
					// TODO: Quota errors are retried forever, it would be good to have support for backoff strategy.
					logger.Warnf(ctx, "Failed to launch job, resource quota exceeded. Err: %v", err)
					return currentState, nil
				}

				currentState = currentState.SetPhase(arrayCore.PhaseRetryableFailure, 0)
				currentState = currentState.SetReason(err.Error())
				return currentState, nil
			}

			return currentState, errors2.Wrapf(ErrSubmitJob, err, "Failed to submit job")
		}
	}

	logger.Infof(ctx, "Successfully submitted Job(s) with Prefix:[%v], Count:[%v]", tCtx.TaskExecutionMetadata().GetTaskExecutionID().GetGeneratedName(), size)

	arrayStatus := arraystatus2.ArrayStatus{
		Summary:  arraystatus2.ArraySummary{},
		Detailed: arrayCore.NewPhasesCompactArray(uint(size)),
	}

	currentState.SetPhase(arrayCore.PhaseCheckingSubTaskExecutions, 0)
	currentState.SetArrayStatus(arrayStatus)

	return currentState, nil
}
