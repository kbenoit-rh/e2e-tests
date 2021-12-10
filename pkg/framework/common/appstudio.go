package common

import (
	"github.com/argoproj/gitops-engine/pkg/health"
	g "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/redhat-appstudio/e2e-tests/pkg/framework"
	commonCtrl "github.com/redhat-appstudio/e2e-tests/pkg/framework/common/controller"
)

var (
	// Pipelines names from https://github.com/redhat-appstudio/infra-deployments/tree/main/components/build/build-templates
	AppStudioPipelinesNames      = []string{"devfile-build", "java-builder", "docker-build", "nodejs-builder", "noop"}
	AppStudioComponents          = []string{"all-components-staging", "authentication", "build", "gitops", "has"}
	AppStudioComponentsNamespace = "openshift-gitops"
	PipelinesNamespace           = "build-templates"
)

var _ = framework.CommonSuiteDescribe("Red Hat App Studio common E2E", func() {
	defer g.GinkgoRecover()
	commonController, err := commonCtrl.NewCommonSuiteController()
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	g.Context("Argo CD", func() {
		for _, component := range AppStudioComponents {
			g.It(component+" status", func() {
				componentStatus, err := commonController.GetAppStudioComponentStatus(component, AppStudioComponentsNamespace)
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(componentStatus.Health.Status).To(gomega.Equal(health.HealthStatusHealthy))

			})
		}
	})

	g.Context("Pipelines:", func() {
		for _, pipelineName := range AppStudioPipelinesNames {
			g.It("Check if "+pipelineName+" pipeline is pre-created", func() {
				p, err := commonController.GetPipeline(pipelineName, PipelinesNamespace)
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(p.Name).To(gomega.Equal(pipelineName))
			})
		}
	})
})