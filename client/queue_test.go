package client

import (
	"testing"

	"github.com/kilgaloon/leprechaun/daemon"
)

func TestBuildQueue(t *testing.T) {
	Agent.BuildQueue()

	Agent.Lock()
	if len(Agent.Queue.Stack) != 6 {
		t.Errorf("Queue stack length expected to be 6, got %d", len(Agent.Queue.Stack))
	}
	Agent.Unlock()

}

func TestQueue(t *testing.T) {
	// reset queue to 0 to test AddToQueue
	Agent.Lock()
	q := &Agent.Queue
	q.Stack = q.Stack[:0]
	Agent.Unlock()

	Agent.AddToQueue(Agent.GetConfig().GetRecipesPath() + "/schedule.yml")

	Agent.Lock()
	if len(Agent.Queue.Stack) != 1 {
		t.Errorf("Queue stack length expected to be 1, got %d", len(Agent.Queue.Stack))
	}
	Agent.Unlock()

	Agent.AddToQueue(Agent.GetConfig().GetRecipesPath() + "/hook.yml")

	Agent.Lock()
	if len(Agent.Queue.Stack) != 1 {
		t.Errorf("Queue stack length expected to be 0, got %d", len(Agent.Queue.Stack))
	}
	Agent.Unlock()

	Agent.Pause()
	Agent.ProcessQueue()

	Agent.SetStatus(daemon.Started)
	Agent.ProcessQueue()
}

func TestFindInRecipe(t *testing.T) {
	// reset queue to 0 to test AddToQueue
	if Agent.FindRecipe("schedule") == nil {
		t.Fail()
	}

	if Agent.FindRecipe("random_name") != nil {
		t.Fail()
	}
}
