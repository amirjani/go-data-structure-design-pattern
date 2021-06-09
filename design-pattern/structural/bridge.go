// Bridge decouples the implementation from the abstraction

// The abstract base class can be subclassed to provide different implementations and allow implementation details to be modified easily

// * The bridge patterns allow the implementation details to change at runtime

// Runtime binding of the application, mapping of orthogonal class hierarchies, and platform independence implementation
// are the scenarios where the bridge pattern can be applied

package main

import "fmt"

type IDrawShape interface {
	drawShape(x [5]float32, y [5]float32)
}

type IContour interface {
	drawContour(x [5]float32, y [5]float32)
	resizeByFactor(factor int)
}

type DrawShape struct{}

type DrawContour struct {
	x      [5]float32
	y      [5]float32
	shape  DrawShape
	factor int
}

func (drawShape DrawShape) drawShape(x [5]float32, y [5]float32) {
	fmt.Println("drawing Shape")
}

func (contour DrawContour) drawContour(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Contour")
	contour.shape.drawShape(contour.x, contour.y)
}

func (contour DrawContour) resizeByFactor(factor int) {
	contour.factor = factor
}

func main() {
	var x = [5]float32{1, 2, 3, 4, 5}
	var y = [5]float32{1, 2, 3, 4, 5}

	var contour IContour = DrawContour{x, y, DrawShape{}, 2}
	contour.drawContour(x, y)
	contour.resizeByFactor(2)
}
