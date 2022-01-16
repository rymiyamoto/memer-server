package response

import "github.com/rymiyamoto/memer-server/model"

type (
	// HelloWorld struct
	HelloWorld struct {
		Text string `json:"text"`
	}

	// HelloWorlds struct list
	HelloWorlds []HelloWorld
)

// NewHelloWorld initialize
func NewHelloWorld(m *model.HelloWorld) *HelloWorld {
	if m == nil {
		return nil
	}

	return &HelloWorld{
		Text: m.Text,
	}
}

// NewHelloWorlds initialize
func NewHelloWorlds(m *model.HelloWorlds) *HelloWorlds {
	ret := HelloWorlds{}

	if m == nil || len(*m) == 0 {
		return &ret
	}

	for i := range *m {
		if r := NewHelloWorld(&(*m)[i]); r != nil {
			ret = append(ret, *r)
		}
	}

	return &ret
}
