package main

import (
	"github.com/kpiotrowski/go-predator/prebot"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/api"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/opencv"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	master := gobot.NewMaster()
	api.NewAPI(master).Start()

	r := raspi.NewAdaptor()

	//Create end-stop drivers
	endX := gpio.NewLimitSwitchDriver(r, prebot.END_X_PIN)
	endX.SetName("EndX")
	endY := gpio.NewLimitSwitchDriver(r, prebot.END_Y_PIN)
	endY.SetName("EndY")
	//Create laser driver
	laser := gpio.NewLedDriver(r, prebot.LASER_PIN)
	//Create stepper motor drivers
	motorX := gpio.NewStepperMotorDriver(r, prebot.MOTOR_X_STEP, prebot.MOTOR_X_DIR, prebot.MOTOR_X_ENABLE)
	motorX.SetName("MOTOR X")
	motorY := gpio.NewStepperMotorDriver(r, prebot.MOTOR_Y_STEP, prebot.MOTOR_Y_DIR, prebot.MOTOR_Y_ENABLE)
	motorY.SetName("MOTOR Y")
	motorX.ConfigureEndDetection(endX, nil, 250)
	motorY.ConfigureEndDetection(endY, nil, 250)
	motorX.Microstepping, motorY.Microstepping = prebot.MICROSTEPPING, prebot.MICROSTEPPING
	//Create window and camera
	window := opencv.NewWindowDriver()
	camera := opencv.NewCameraDriver(0)
	//Create PREDATOR
	predator := prebot.NewPredator(camera, laser, motorX, motorY)
	predator.Window = window
	robot := gobot.NewRobot("predator",
		[]gobot.Connection{r},
		[]gobot.Device{motorX, motorY},
		predator.Run,
	)
	//Start devices
	master.AddRobot(robot)
	robot.Start()
}
