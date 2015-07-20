package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"shape1/shapes"
)

func main() {
	log.SetFlags(0)
	const width, height = 400, 200
	img := shapes.FilledImage(width, height,
		color.RGBA{0xFF, 0xFF, 0xFF, 0xFF})
	x, y := width/4, height/2

	red := color.RGBA{0xFF, 0, 0, 0xFF}
	blue := color.RGBA{0, 0, 0xFF, 0xFF}
	// Purely for testing New() vs. New*()
	if len(os.Args) == 1 {
		fmt.Println("Using NewCircle() & NewRegularPolygon()")
		circle := shapes.NewCircle(blue, 90)
		circle.SetFill(red) // Uses the aggregated shape.SetFill method
		octagon := shapes.NewRegularPolygon(red, 75, 8)
		octagon.SetFill(blue) // Uses the aggregated circle.shape.SetFill
		polygon := shapes.NewRegularPolygon(image.Black, 65, 4)
		if err := shapes.DrawShapes(img, x, y, circle, octagon, polygon); err != nil {
			fmt.Println(err)
		}
		sanityCheck("circle", circle)
		sanityCheck("octagon", octagon)
		sanityCheck("polygon", polygon)
	} else {
		fmt.Println("Using New()")
		// The Shapers returned by New can only call
		// Shaper methods (Fill(), SetFill(), and Draw());
		// however, we can use type assertion if we need to access other
		// methods.
		if _, err := shapes.New("Misshapen", shapes.Option{blue, 5}); err == nil {
			fmt.Println("unexpectedly got a non-nil invalid shape!")
		}
		circle, _ := shapes.New("circle", shapes.Option{blue, 5})
		circle.SetFill(red)
		circle.(shapes.CircularShaper).SetRadius(90)
		octagon, _ := shapes.New("octagon", shapes.Option{red, 10})
		octagon.SetFill(blue)
		// This type assertion changes the original octagon because the new
		// octagon is in effect a reference to a shapes.RegularPolygonalShaper
		// object
		if octagon, ok := octagon.(shapes.RegularPolygonalShaper); ok {
			octagon.SetSides(75)
		}
		polygon, _ := shapes.New("square", shapes.Option{Radius: 65})
		if err := shapes.DrawShapes(img, x, y, circle, octagon, polygon); err != nil {
			fmt.Println(err)
		}
		sanityCheck("circle", circle)
		sanityCheck("octagon", octagon)
		sanityCheck("polygon", polygon)
	}
	polygon := shapes.NewRegularPolygon(color.RGBA{0, 0x7F, 0, 0xFF}, 65, 4)
	showShapeDetails(polygon)
	y = 30
	for i, radius := range []int{60, 55, 50, 45, 40} {
		polygon.SetRadius(radius)
		polygon.SetSides(i + 5)
		x += radius
		y += height / 8
		if err := shapes.DrawShapes(img, x, y, polygon); err != nil {
			fmt.Println(err)
		}
	}

	filename := "shapes.png"
	if err := shapes.SaveImage(img, filename); err != nil {
		log.Println(err)
	} else {
		fmt.Println("Saved", filename)
	}
	fmt.Println("OK")

	img = shapes.FilledImage(width, height, image.White)
	x, y = width/3, height/4
}

func sanityCheck(name string, shape shapes.Shaper) {
	fmt.Print("name=", name, " ")
	fmt.Print("fill=", shape.Fill(), " ")
	if shape, ok := shape.(shapes.CircularShaper); ok {
		fmt.Print("radius=", shape.Radius(), " ")
		if shape, ok := shape.(shapes.RegularPolygonalShaper); ok {
			fmt.Print("sides=", shape.Sides(), " ")
		}
	}
	fmt.Println()
}

func showShapeDetails(shape shapes.Shaper) {
	fmt.Print("fill=", shape.Fill(), " ")               // All shapes have a fill color
	if shape, ok := shape.(shapes.CircularShaper); ok { // shadow variable
		fmt.Print("radius=", shape.Radius(), " ")
		if shape, ok := shape.(shapes.RegularPolygonalShaper); ok { //shadow
			fmt.Print("sides=", shape.Sides(), " ")
		}
	}
	fmt.Println()
}
