package game

import (
	// color "image/color"
	"math"
	_ "image/png"
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const DT = 0.016

type Ship struct{
	x, y float64
	vx, vy float64
	speed float64
	img *ebiten.Image
	op *ebiten.DrawImageOptions
	scale float64
	state *State
}

func (s *Ship) SetPos (x, y float64) {
	s.x, s.y = x, y
	s.op.GeoM.Reset()
	s.op.GeoM.Scale(s.scale, s.scale)
	s.op.GeoM.Translate(s.x, s.y)
}

func (s *Ship) Init(myState *State) {
	s.state = myState
	s.speed = 400
	var err error
	s.img, _, err = ebitenutil.NewImageFromFile("ship.png")
	if err != nil {
		log.Fatal(err)
	}
	s.scale = 4
	s.op = &ebiten.DrawImageOptions{}
	s.op.GeoM.Scale(s.scale, s.scale)
}

func (s *Ship) Update() {
	var x, y = 0.0, 0.0
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		y -= 1.0
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		y += 1.0
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		x -= 1.0
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		x += 1.0
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		{
			var bullet = new(Bullet)
			bullet.Init(s.state)
			bullet.SetPos(s.x, s.y)
			s.state.AddEntity(bullet)
		}
		{
			var bullet = new(Bullet)
			bullet.Init(s.state)
			bullet.SetPos(s.x + 32, s.y)
			s.state.AddEntity(bullet)
		}
	}
	s.vx += x * DT * 5
	s.vy += y * DT * 5
	s.vx = math.Max(math.Min(s.vx, 1.0), -1.0)
	s.vy = math.Max(math.Min(s.vy, 1.0), -1.0)
	s.vx *= 0.95
	s.vy *= 0.95
	s.x += s.vx * DT * s.speed
	s.y += s.vy * DT * s.speed
	if s.x > 800 {	
		s.x = 800
	}
	if s.x < 0 {
		s.x = 0
	}
	if s.y < 0 {
		s.y = 0
	}
	if s.y > 600 {
		s.y = 600
	}
	s.op.GeoM.Reset()
	s.op.GeoM.Scale(s.scale, s.scale)
	s.op.GeoM.Translate(s.x, s.y)
}

func (s *Ship) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.img, s.op)
}