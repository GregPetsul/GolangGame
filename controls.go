package main

import (
	"math"
	"strings"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type playerMover struct {
	container *element
	speed     float64
	animator  *animator
}

func newPlayerMover(container *element, speed float64) *playerMover {
	return &playerMover{
		container: container,
		speed:     speed,
		animator:  container.getComponent(&animator{}).(*animator),
	}
}

func (mover *playerMover) onUpdate() error {
	keys := sdl.GetKeyboardState()
	cont := mover.container
	cont.velocity = vector{
		x: 0,
		y: 0,
	}

	//left and right movement
	if keys[sdl.SCANCODE_LEFT] == 1 || keys[sdl.SCANCODE_A] == 1 {
		if cont.position.x /*-(mover.sr.width/2.0)*/ > 0 {
			moveElements(mover, "left")
			cont.velocity.x = -mover.speed
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 || keys[sdl.SCANCODE_D] == 1 {
		if cont.position.x /*+(mover.sr.height/2.0)*/ < screenWidth {
			moveElements(mover, "right")
			cont.velocity.x = mover.speed
		}
	}
	//up and down movement
	if keys[sdl.SCANCODE_DOWN] == 1 || keys[sdl.SCANCODE_S] == 1 {
		moveElements(mover, "down")
		cont.velocity.y = -mover.speed
	} else if keys[sdl.SCANCODE_UP] == 1 || keys[sdl.SCANCODE_W] == 1 {
		moveElements(mover, "up")
		cont.velocity.y = mover.speed
	}

	mover.setAnimation()

	return nil
}

func moveElements(mover *playerMover, dir string) {
	for _, elem := range elements {
		if elem.active && elem.tag != "player" {
			if dir == "left" {
				elem.position.x += mover.speed * delta
			} else if dir == "right" {
				elem.position.x -= mover.speed * delta
			} else if dir == "up" {
				elem.position.y += mover.speed * delta
			} else if dir == "down" {
				elem.position.y -= mover.speed * delta
			}

		}
	}
}

func (mover *playerMover) setAnimation() {
	vel := mover.container.velocity
	an := mover.animator

	switch {
	case vel.x > 0:
		an.setSequence("right_walk")
	case vel.x < 0:
		an.setSequence("left_walk")
	case vel.y > 0:
		an.setSequence("back_walk")
	case vel.y < 0:
		an.setSequence("front_walk")
	case vel.y == 0 && vel.x == 0:
		dir := strings.Split(an.current, "_")[0]
		an.setSequence(dir + "_idle")
	}
}

func (mover *playerMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}
func (mover *playerMover) onCollision(other *element) error {
	return nil
}

type playerShooter struct {
	container *element
	cooldown  time.Duration
	lastShot  time.Time
}

func newPlayerShooter(container *element, cooldown time.Duration) *playerShooter {
	return &playerShooter{
		container: container,
		cooldown:  cooldown,
	}
}

func (mover *playerShooter) onUpdate() error {
	xmouse, ymouse, mouse := sdl.GetMouseState()

	pos := mover.container.position

	if mouse == 1 {
		if time.Since(mover.lastShot) >= mover.cooldown {
			mover.shoot(pos.x, pos.y, xmouse, ymouse)

			mover.lastShot = time.Now()
		}
	}

	return nil
}

func (mover *playerShooter) onDraw(renderer *sdl.Renderer) error {
	return nil
}
func (mover *playerShooter) onCollision(other *element) error {
	return nil
}

func (mover *playerShooter) shoot(x, y float64, xmouse, ymouse int32) {
	if bul, ok := bulletFromPool(); ok {
		bul.active = true
		bul.position.x = x
		bul.position.y = y
		bul.rotation = math.Atan2(float64(ymouse)-y, float64(xmouse)-x)
		bul.imageOffset = 45
		bul.tag = "bullet"
	}
}
