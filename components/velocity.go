package components

const velocityComponentName = "Velocity"

type Velocity struct {
	X int
	Y int
}

func (v Velocity) GetComponentName() string { return velocityComponentName}