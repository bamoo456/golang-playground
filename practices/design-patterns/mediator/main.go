package main

import "errors"
import "fmt"

// The mediator pattern defines an object that encapsulates how a set of objects interact.
// This pattern is considered to be a behavioral pattern due to the way it can alter the program's running behavior.

type ButtonMediator interface {
	ShowButton(Button)
	HideButton(Button)
}

type Button interface {
	Name() string
	IsEnabled() bool
	Enable(bool)
	Click()
}

// implement the Button
type button struct {
	name     string
	enabled  bool
	mediator ButtonMediator
}

func (b *button) Name() string {
	return b.name
}

func (b *button) Enable(enableFlag bool) {
	b.enabled = enableFlag
}

func (b *button) IsEnabled() bool {
	return b.enabled
}

func (b *button) Click() {
	b.mediator.ShowButton(b)
}

// implement the Button
type radioButton struct {
	name     string
	enabled  bool
	checked  bool
	mediator ButtonMediator
}

func (r *radioButton) Name() string {
	return r.name
}

func (r *radioButton) Enable(enableFlag bool) {
	r.enabled = enableFlag
}

func (r *radioButton) IsEnabled() bool {
	return r.enabled
}

func (r *radioButton) Click() {
	r.checked = true
	r.mediator.ShowButton(r)
}

// Implement the ButtonMediator
type buttonGroup struct {
	buttons map[string]Button
}

func (g *buttonGroup) ShowButton(b Button) {
	for _, btn := range g.buttons {
		btn.Enable(btn.Name() == b.Name())
	}
}

func (g *buttonGroup) HideButton(b Button) {
	b.Enable(false)
}

func (g *buttonGroup) AddButton(b Button) error {
	if _, ok := g.buttons[b.Name()]; ok {
		return errors.New("Button already exists")
	}
	g.buttons[b.Name()] = b
	return nil
}

func (g *buttonGroup) String() string {
	state := ""
	for _, btn := range g.buttons {
		state += fmt.Sprintf("button: [%s], enabled: [%v]\n", btn.Name(), btn.IsEnabled())
	}
	return state
}

func NewButtonGroup() *buttonGroup {
	return &buttonGroup{
		buttons: map[string]Button{},
	}
}

func main() {
	group := NewButtonGroup()
	baseBtn := &button{
		name:     "loginButton",
		enabled:  false,
		mediator: group,
	}
	checkedBtn := &radioButton{
		name:     "checkedButton",
		enabled:  false,
		mediator: group,
	}

	group.AddButton(baseBtn)
	group.AddButton(checkedBtn)

	fmt.Printf("=====group=====\n%s\n", group)

	baseBtn.Click()
	fmt.Printf("=====group=====\n%s\n", group)

	checkedBtn.Click()
	fmt.Printf("=====group=====\n%s\n", group)
}
