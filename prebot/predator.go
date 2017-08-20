package prebot

import (
	"fmt"
	cv "github.com/lazywei/go-opencv/opencv"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/opencv"
)

type Predator struct {
	laser  *gpio.LedDriver
	endX   *gpio.LimitSwitchDriver
	endY   *gpio.LimitSwitchDriver
	motorX *gpio.StepperMotorDriver
	motorY *gpio.StepperMotorDriver

	camera *opencv.CameraDriver
	Window *opencv.WindowDriver
}

func NewPredator(cam *opencv.CameraDriver, laser *gpio.LedDriver, motorX, motorY *gpio.StepperMotorDriver) *Predator {
	predator := &Predator{
		camera: cam,
		laser:  laser,
		motorX: motorX,
		motorY: motorY,
	}
	return predator
}

func (p *Predator) Run() {
	var image *cv.IplImage
	p.camera.On(opencv.Frame, func(data interface{}) {
		image = data.(*cv.IplImage)

		if image != nil {
			println("DUPA JEST IMAGE")
			i := image.Clone()
			normal_i := i.ToImage()
			fmt.Println(normal_i.At(10, 10).RGBA())
			//faces := opencv.DetectFaces(cascade, i)

			//i = opencv.DrawRectangles(i, faces, 0, 255, 0, 5)

			if p.Window != nil {
				p.Window.ShowImage(i)
			}
		}
	})
	if p.Window != nil {
		p.Window.Start()
	}
	p.camera.Start()
}
