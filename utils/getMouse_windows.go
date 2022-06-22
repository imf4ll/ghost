package utils

// go:build windows

import (
	"os"
	"sync"

	"github.com/go-vgo/robotgo"
)

var wg sync.WaitGroup

type MousePosition struct {
    X1,
    X2,
    Y1,
    Y2 int
}

func GetMouse() MousePosition {
	var mousePos MousePosition
	
	mouseFirstClick := robotgo.AddMouse("left")
	
	if mouseFirstClick {
		mousePos.X1, mousePos.Y1 = robotgo.GetMousePos()
		
	}
	
	mouseLastClick := robotgo.AddMouse("left")
	
	if mouseLastClick {
		mousePos.X2, mousePos.Y2 = robotgo.GetMousePos()
	
	}
	
	return mousePos
}

func GetMouseAndKeyboard() MousePosition {
	var mousePos MousePosition
	
	go func() {
		wg.Add(1)
		
		key := robotgo.AddEvent("esc")
		
		if key {
			os.Exit(3)
			
		}
	}()
	
	mouseFirstClick := robotgo.AddMouse("left")
	
	if mouseFirstClick {
		mousePos.X1, mousePos.Y1 = robotgo.GetMousePos()
	
	}
	
	mouseLastClick := robotgo.AddMouse("left")
	
	if mouseLastClick {
		mousePos.X2, mousePos.Y2 = robotgo.GetMousePos()
	
	}
	
	wg.Done()
	wg.Wait()
	
	return mousePos
}
