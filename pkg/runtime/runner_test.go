package runtime

import (
	"testing"

	"github.com/commander-cli/commander/pkg/suite"
	"github.com/stretchr/testify/assert"
)

func Test_RunnerExexcute(t *testing.T) {
	s := getExampleTestCases()
	r := Runner{
		Nodes: getExampleNodes(),
	}

	got := r.Run(s)

	assert.IsType(t, make(<-chan TestResult), got)

	count := 0
	for r := range got {
		assert.Equal(t, "Output hello", r.TestCase.Title)
		assert.True(t, r.ValidationResult.Success)
		count++
	}
	assert.Equal(t, 1, count)
}

func Test_getExecutor(t *testing.T) {
	r := Runner{
		Nodes: getExampleNodes(),
	}

	exec := r.getExecutor("ssh")
	assert.IsType(t, SSHExecutor{}, exec)
	exec = r.getExecutor("local")
	assert.IsType(t, LocalExecutor{}, exec)
	exec = r.getExecutor("docker")
	assert.IsType(t, DockerExecutor{}, exec)
}

func getExampleNodes() []suite.Node {
	n1 := suite.Node{
		Name: "ssh",
		Type: "ssh",
	}
	n2 := suite.Node{
		Name: "local",
		Type: "local",
	}
	n3 := suite.Node{
		Name: "docker",
		Type: "docker",
	}

	nodes := []suite.Node{
		n1, n2, n3,
	}
	return nodes
}
