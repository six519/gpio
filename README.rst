gpio
====

Write to Raspberry Pi GPIO Pins Go Library

Installation
============

`go get github.com/six519/gpio`

Sample Usage
============
::

    package main

    import (
        "github.com/six519/gpio"
        "time"
    )

    func main() {
        gpio.PinMode(21, gpio.OUTPUT)
        gpio.DigitalWrite(21, gpio.HIGH)
        time.Sleep(time.Second * 5)
        gpio.DigitalWrite(21, gpio.LOW)
        gpio.CleanUp(21)
    }   
