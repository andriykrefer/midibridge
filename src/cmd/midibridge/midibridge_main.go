package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/andriykrefer/midibridge/midi_parser"

	midi "gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/drivers"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv" // autoregisters driver
	"go.bug.st/serial"
)

const CH = 0

func main() {
	defer midi.CloseDriver()
	time.Sleep(200 * time.Millisecond)

	// serialPort := "COM8"
	serialPort := askSerialPort()
	// noteOut, _ := midi.FindOutPort("loop")
	noteOut := askMidiPort()

	sendNote, err := midi.SendTo(noteOut)
	if err != nil {
		panic(err)
	}

	// SERIAL
	mode := &serial.Mode{
		BaudRate: 31250,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open(serialPort, mode)
	if err != nil {
		panic(err)
	}
	b := midi_parser.BufferParser{}
	buf := make([]byte, 200)
	for {
		n, _ := port.Read(buf)
		msgs := b.Process(buf[0:n])
		for k := range msgs {
			// fmt.Println(msgs[k])
			sendNote(msgs[k])
		}
	}
}

func askSerialPort() string {
	list, err := serial.GetPortsList()
	if err != nil {
		panic(err)
	}
	fmt.Println("Available serial ports: ")
	for k, port := range list {
		fmt.Printf("[%d] %s", k, port)
		fmt.Println("")
		fmt.Println("")
	}
	fmt.Print("Enter option: ")
	reader := bufio.NewReader(os.Stdin)
	strIn, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	opt, err := strconv.Atoi(strings.TrimSpace(strIn))
	if err != nil {
		panic(err)
	}
	return list[opt]
}

func askMidiPort() drivers.Out {
	list := midi.GetOutPorts()
	fmt.Println("Available midi OUT ports: ")
	fmt.Println(list.String())
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter option number: ")
	strIn, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	opt, err := strconv.Atoi(strings.TrimSpace(strIn))
	if err != nil {
		panic(err)
	}
	return list[opt]
}
