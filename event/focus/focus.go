package focus

import (
	"fmt"
)

// Event is a focus event.
type Event struct {
	// In indicates if the window was focused in or out (In=true -> focused in, In=false -> focused out)
	In bool
}

func (e Event) String() string {
	str := "FocusedIn"
	if !e.In {
		str = "FocusedOut"
	}
	return fmt.Sprintf("focus.Event{(%s)}", str)
}
