package jobttlseconds

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"golang.stackrox.io/kube-linter/pkg/lintcontext/mocks"
	"golang.stackrox.io/kube-linter/pkg/templates"
)

const (
	job_no_ttl    = "job_no_ttl"
	job_ttl       = "job_ttl"
	job_large_ttl = "job_large_ttl-matches-none"
)

func TestJobTTLSeconds(t *testing.T) {
	suite.Run(t, new(JobTTLSecondsTestSuite))
}

type JobTTLSecondsTestSuite struct {
	templates.TemplateTestSuite

	ctx *mocks.MockLintContext
}

func (s *JobTTLSecondsTestSuite) SetupTest() {
	s.Init(templateKey)
	s.ctx = mocks.NewMockContext()
}
