package jobttlseconds

import (
	"fmt"

	"golang.stackrox.io/kube-linter/pkg/check"
	"golang.stackrox.io/kube-linter/pkg/config"
	"golang.stackrox.io/kube-linter/pkg/diagnostic"
	"golang.stackrox.io/kube-linter/pkg/lintcontext"
	"golang.stackrox.io/kube-linter/pkg/objectkinds"
	"golang.stackrox.io/kube-linter/pkg/templates"
	"golang.stackrox.io/kube-linter/pkg/templates/jobttlseconds/internal/params"
	batchV1 "k8s.io/api/batch/v1"
)

const templateKey = "no-job-ttl-seconds"
const maxTTL int32 = 100

func init() {
	templates.Register(check.Template{
		HumanName:   "No Job TTL Service",
		Key:         templateKey,
		Description: "Flag jobs that do not set spec.ttlSecondsAfterFinished",
		SupportedObjectKinds: config.ObjectKindsDesc{
			ObjectKinds: []string{objectkinds.DeploymentLike},
		},
		Parameters:             params.ParamDescs,
		ParseAndValidateParams: params.ParseAndValidate,
		Instantiate: params.WrapInstantiateFunc(func(_ params.Params) (check.Func, error) {
			return func(lintCtx lintcontext.LintContext, object lintcontext.Object) []diagnostic.Diagnostic {
				job, ok := object.K8sObject.(*batchV1.Job)
				if !ok {
					return nil
				}
				if job.Spec.TTLSecondsAfterFinished == nil {
					return []diagnostic.Diagnostic{{Message: fmt.Sprintf("Job not specifying ttlSecondsAfterFinished")}}
				}
				if *job.Spec.TTLSecondsAfterFinished > maxTTL {
					return []diagnostic.Diagnostic{{Message: fmt.Sprintf("Job specifying too large ttlSecondsAfterFinished")}}
				}
				return nil
			}, nil
		}),
	})
}
