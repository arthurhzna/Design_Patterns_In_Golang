# Design Patterns in Golang

## Adapter vs Decorator

### Adapter

`Adapter` is used when the object we have does not yet match the contract expected by the client.

In short:
- the client expects a specific interface
- the legacy or external object does not satisfy that interface yet
- we create an adapter so the object can be used

Pattern example:

```go
type ShapeRenderer interface {
	Render() string
}

type LegacyCircle struct{}

func (l *LegacyCircle) Draw() string {
	return "circle"
}

type ShapeAdapter struct {
	legacy *LegacyCircle
	color  string
}

func (a *ShapeAdapter) Render() string {
	return a.legacy.Draw() + " with color " + a.color
}
```

In this example, the adapter also adds color, but its main role is still `adapter` because the real goal is to make `LegacyCircle` compatible with the `Render()` interface.

Easy way to remember:
- adapter = make it fit
- main focus = compatibility

### Decorator

`Decorator` is used when the original object already matches the interface, but we want to add behavior without changing the original struct.

In short:
- the interface already matches
- the original object is already usable
- we wrap that object to add features

Pattern example:

```go
type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f", c.Radius)
}

type ColoredShape struct {
	Shape Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has the color %s", c.Shape.Render(), c.Color)
}
```

Here, `Circle` already satisfies `Shape`. `ColoredShape` does not change the interface; it only adds behavior to `Render()` by adding color.

Easy way to remember:
- decorator = make it richer
- main focus = enhancement

### Main Differences

- `Adapter`: the object does not yet match the required interface, so a wrapper is created to make it compatible.
- `Decorator`: the object already matches the interface, and then it is wrapped to add behavior.
- `Adapter`: emphasizes contract compatibility.
- `Decorator`: emphasizes behavior enhancement.

### Very Short Summary

- if the problem is "this object cannot be used because its method/interface is different" -> `Adapter`
- if the problem is "this object already works, but I want to add features such as color, transparency, or logging" -> `Decorator`