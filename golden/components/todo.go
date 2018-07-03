package components

// Note: this entire file is autogenerated

import (
	"github.com/bketelsen/factor/golden/models"
)

// Todo is how factor binds a model to a component. It holds the an autogenerated
// TodoClient that holds a single Todo in memory and also provides a typesafe
// interface to the RPC calls to modify the TODO on the server.
//
// Don't create this struct directly, use NewTodo instead
type Todo struct {
	// this is the
	Client models.TodoClient
}

// Get is a convenience function that gets the current in-memory Todo
func (t TodoBinding) Get() models.Todo {
	return t.Client.GetCurrent()
}

// Permalink is a convenience function to get the permalink for this todo.
func (t TodoBinding) Permalink() string {
	return fmt.Sprintf("/todos/%s", t.Get().ID.String())
}

// NewTodo creates a new Todo with a fully populated client
func NewTodo(t models.Todo) Todo {
	return Todo{
		Client: models.NewTodoClient(t)
	}
}
