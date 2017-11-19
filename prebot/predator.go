package prebot

import (
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/opencv"
	"gocv.io/x/gocv"
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
	//err := p.motorX.Min()
	//if err != nil {
	//	panic(err)
	//}
	//err = p.motorY.Min()
	//if err != nil {
	//	panic(err)
	//}

	//classifier := gocv.NewCascadeClassifier()
	//defer classifier.Close()
	//img := gocv.NewMat()
	//defer img.Close()
	//classifier.Load("haarcascade_frontalface_default.xml")
	//



	p.camera.On(opencv.Frame, func(data interface{}) {
		img := data.(gocv.Mat)

		// detect faces
		//rects := classifier.DetectMultiScale(img)
		//fmt.Printf("found %d faces\n", len(rects))
		//
		//// draw a rectangle around each face on the original image
		//for _, r := range rects {
		//	gocv.Rectangle(img, r, color.RGBA{0, 0, 255, 0}, 3)
		//}


		if p.Window != nil {
			p.Window.ShowImage(img)
			p.Window.WaitKey(1)
		}

	})
}
