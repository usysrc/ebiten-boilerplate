package systems

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/usysrc/chaoshell/internal/game/component"
	"github.com/usysrc/chaoshell/internal/game/entity"
)

const DT = 0.016

type InputSystem struct {
	Components *component.Manager
	ShipEntity component.Entity
}

func (is *InputSystem) Update() {
	velocity, exists := is.Components.Velocities[is.ShipEntity]
	if !exists {
		log.Fatal("entity does not have velocity component")
		return
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		velocity.Y -= 1.0 * DT * 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		velocity.Y += 1.0 * DT * 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		velocity.X -= 1.0 * DT * 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		velocity.X += 1.0 * DT * 5
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		bullet := entity.CreateBullet(is.Components)
		bulletPos, exists := is.Components.Positions[bullet]
		if !exists {
			log.Fatal("entity does not have position component")
		}
		shipPos, exists := is.Components.Positions[is.ShipEntity]
		if !exists {
			log.Fatal("entity does not have position component")
		}
		bulletPos.X = shipPos.X
		bulletPos.Y = shipPos.Y

		bulletVelo, exists := is.Components.Velocities[bullet]
		if !exists {
			log.Fatal("entity does not have velocity component")
		}
		bulletVelo.Y = -10
	}

	// apply friction
	velocity.X *= 0.95
	velocity.Y *= 0.95

	// limit to normalized values
	velocity.X = math.Max(math.Min(velocity.X, 1.0), -1.0)
	velocity.Y = math.Max(math.Min(velocity.Y, 1.0), -1.0)

}