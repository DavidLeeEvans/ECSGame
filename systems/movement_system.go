package systems

import (
	"BasicECS/entities"
	"BasicECS/components"
)

type movementSystem struct {}

func (s movementSystem) Update(dt float64, entityManager *entities.EntityManager) {

	entityIds := entityManager.GetAllEntitiesPossessingComponentsOfClass(components.GetMovableComponentName())

	for _, entityId := range entityIds {

		movableComponent := entityManager.GetComponentOfClass(
			components.GetMovableComponentName(),
			entityId).(*components.Movable)

		if !movableComponent.Movable { continue }

		positionComponent := entityManager.GetComponentOfClass(
			components.GetPositionComponentName(),
			entityId).(*components.Position)

		if positionComponent == nil { continue }

		speedComponent := entityManager.GetComponentOfClass(
			components.GetSpeedComponentName(),
			entityId).(*components.Speed)

		if speedComponent == nil { continue }

		inputComponent := entityManager.GetComponentOfClass(
			components.GetInputComponentName(),
			entityId).(*components.Input)

		if inputComponent == nil {
			aiMovement()
		} else {
			controllerMovement(inputComponent, positionComponent, speedComponent)
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
}

func aiMovement() {

}

func CreateMovementSystem() System { return movementSystem{} }