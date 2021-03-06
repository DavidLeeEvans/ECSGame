package systems

import (
	"BasicECS/components"
	"BasicECS/core"
)

type movementSystem struct {}

func (s movementSystem) Update(dt float64, entityManager *core.EntityManager) {

	entityIds := entityManager.GetAllEntitiesPossessingComponentsOfClass(components.GetMovableComponentName())

	for _, entityId := range entityIds {

		movableComponent, _ := entityManager.GetComponentOfClass(
			components.GetMovableComponentName(),
			entityId).(*components.Movable)

		if !movableComponent.Movable { continue }

		positionComponent, _ := entityManager.GetComponentOfClass(
			components.GetPositionComponentName(),
			entityId).(*components.Position)

		if positionComponent == nil { continue }

		speedComponent, _ := entityManager.GetComponentOfClass(
			components.GetSpeedComponentName(),
			entityId).(*components.Speed)

		if speedComponent == nil { continue }

		inputComponent := entityManager.GetComponentOfClass(
			components.GetInputComponentName(),
			entityId)

		if inputComponent == nil {
			aiMovement(positionComponent, speedComponent)
		} else {
			controllerMovement(inputComponent.(*components.Input), positionComponent, speedComponent)
		}
	}
}

func controllerMovement(input *components.Input, position *components.Position, speed *components.Speed) {
	if input.ForwardKey {
		position.Y -= speed.Speed
	}
	if input.BackwardsKey {
		position.Y += speed.Speed
	}
	if input.LeftKey {
		position.X -= speed.Speed
	}
	if input.RightKey {
		position.X += speed.Speed
	}
}

func aiMovement(position *components.Position, speed *components.Speed) {
	position.X += speed.Speed
}

func CreateMovementSystem() System { return movementSystem{} }