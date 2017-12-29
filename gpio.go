package gpio

import (
	"fmt"
	"os"
	"strconv"
)

const OUTPUT string = "out"
const LOW string = "0"
const HIGH string = "1"

var export_path string = "/sys/class/gpio/export"
var gpio_path string = "/sys/class/gpio/gpio"
var unexport_path string = "/sys/class/gpio/unexport"

func writeInt(path string, val int) {
	f, err := os.Create(path)

	if err != nil {
		fmt.Println("Can't open file " + path)
	} else {
		defer f.Close()

		f.WriteString(strconv.Itoa(val))
	}	
}

func writeStr(path string, val string) {
	f, err := os.Create(path)

	if err != nil {
		fmt.Println("Can't open file " + path)
	} else {
		defer f.Close()

		f.WriteString(val)
	}	
}

func PinMode(pin int, mode string) {
	writeInt(export_path, pin)
	direction := gpio_path + strconv.Itoa(pin) + "/direction"
	writeStr(direction, mode)
}

func DigitalWrite(pin int, val string) {
	value := gpio_path + strconv.Itoa(pin) + "/value"
	writeStr(value, val)
}

func CleanUp(pin int) {
	writeInt(unexport_path, pin)
}