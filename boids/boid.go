package main

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	position Vector2D
	velocity Vector2D
	id       int
}

func (b *Boid) calcAcceleration() Vector2D {
	upper, lower := b.position.AddV(viewRad), b.position.AddV(-viewRad)
	avgPosition, avgVelocity, separation := Vector2D{0, 0}, Vector2D{0, 0}, Vector2D{0, 0}
	count := 0.0
	rwlock.RLock()
	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, width); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, height); j++ {
			if otherBoidId := boidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.id {
				if dist := boids[otherBoidId].position.Distance(b.position); dist < viewRad {
					count++
					avgVelocity = avgVelocity.Add(boids[otherBoidId].velocity)
					avgPosition = avgPosition.Add(boids[otherBoidId].position)
					separation = separation.Add(b.position.Sub(boids[otherBoidId].position).DivV(dist))
				}
			}
		}
	}
	rwlock.RUnlock()

	accel := Vector2D{b.borderBounce(b.position.x, width), b.borderBounce(b.position.y, height)}
	if count > 0 {
		avgVelocity, avgPosition = avgVelocity.DivV(count), avgPosition.DivV(count)
		accelAlignment := avgVelocity.Sub(b.velocity).MulV(adjRate)
		accelCohesion := avgPosition.Sub(b.position).MulV(adjRate)
		accelSeparation := separation.MulV(adjRate)
		accel = accel.Add(accelAlignment).Add(accelCohesion).Add(accelSeparation)
	}
	return accel
}

func (b *Boid) borderBounce(pos, maxBorderPos float64) float64 {
	if pos < viewRad {
		return 1 / pos
	} else if pos > maxBorderPos-viewRad {
		return 1 / (pos - maxBorderPos)
	}
	return 0
}

func (b *Boid) moveOne() {
	acceleration := b.calcAcceleration()
	rwlock.Lock()
	b.velocity = b.velocity.Add(acceleration).limit(-1, 1)
	boidMap[int(b.position.x)][int(b.position.y)] = -1
	b.position = b.position.Add(b.velocity)
	boidMap[int(b.position.x)][int(b.position.y)] = b.id
	rwlock.Unlock()
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func createBoid(bid int) {
	b := Boid{
		position: Vector2D{rand.Float64() * width, rand.Float64() * height},
		velocity: Vector2D{(rand.Float64() * 2) - 1.0, (rand.Float64() * 2) - 1.0},
		id:       bid,
	}
	boids[bid] = &b
	boidMap[int(b.position.x)][int(b.position.y)] = b.id
	go b.start()
}
