package main

import (
	"fmt"
	"github.com/ByteArena/box2d"
	"github.com/faiface/pixel/pixelgl"
	"phy/src/physic"
	"sort"
)

func main() {

	pixelgl.Run(Render)

}

func Render() {
	gravity := box2d.MakeB2Vec2(0.0, -10.0)
	world := box2d.MakeB2World(gravity)
	sys := &physic.CollisionSystem{}
	sys.New(&world)
	characters := make(map[string]*box2d.B2Body)

	// Ground body
	{
		bd := box2d.MakeB2BodyDef()
		ground := world.CreateBody(&bd)

		shape := box2d.MakeB2EdgeShape()
		shape.Set(box2d.MakeB2Vec2(-20.0, 0.0), box2d.MakeB2Vec2(20.0, 0.0))
		ground.CreateFixture(&shape, 0.0)
		characters["ground"] = ground
		sys.Add(&physic.Box2dComponent{ground})
	}

	// Circle character
	{
		bd := box2d.MakeB2BodyDef()
		bd.Position.Set(3.0, 25.0)
		bd.Type = box2d.B2BodyType.B2_dynamicBody
		bd.FixedRotation = true
		bd.AllowSleep = false

		body := world.CreateBody(&bd)

		shape := box2d.MakeB2CircleShape()
		shape.M_radius = 0.5

		fd := box2d.MakeB2FixtureDef()
		fd.Shape = &shape
		fd.Density = 20.0
		body.CreateFixtureFromDef(&fd)
		characters["circle"] = body
		sys.Add(&physic.Box2dComponent{body})
	}

	timeStep := 1.0 / 60.0
	velocityIterations := 8
	positionIterations := 3

	output := ""

	characterNames := make([]string, 0)
	for k, _ := range characters {
		characterNames = append(characterNames, k)
	}
	sort.Strings(characterNames)

	for {
		world.Step(timeStep, velocityIterations, positionIterations)
		sys.Update(float32(timeStep))
		// Now print the position and angle of the body.
		for _, name := range characterNames {
			character := characters[name]
			position := character.GetPosition()
			angle := character.GetAngle()
			msg := fmt.Sprintf("(%s): %4.3f %4.3f %4.3f\n", name, position.X, position.Y, angle)
			output += msg
			//fmt.Println(msg)
		}
	}

}
