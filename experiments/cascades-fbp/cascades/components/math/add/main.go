package main

import (
	"encoding/binary"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cascades-fbp/cascades/components/utils"
	"github.com/cascades-fbp/cascades/runtime"
	zmq "github.com/pebbe/zmq4"
)

var (
	// Flags
	inputaEndpoint = flag.String("port.ina", "", "Component's input port endpoint")
	inputbEndpoint = flag.String("port.inb", "", "Component's input port endpoint")
	sumEndpoint    = flag.String("port.sum", "", "Component's output port endpoint")
	jsonFlag       = flag.Bool("json", false, "Print component documentation in JSON")
	debug          = flag.Bool("debug", false, "Enable debug mode")

	// Internal
	inaPort, inbPort, sumPort *zmq.Socket
	inaCh, inbCh, sumCh       chan bool
	exitCh                    chan os.Signal
	err                       error
)

type operands struct {
	A float32 `json:"A"`
	B float32 `json:"B"`
}

func readjson(path string) (operands, error) {
	var op operands
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("failed in reading json file")
		log.Fatal(err)
		return op, err
	}
	err = json.Unmarshal(raw, &op)
	if err != nil {
		fmt.Println("failed in unmarshalling paper")
		log.Fatal(err)
		return op, err
	}
	fmt.Println("operands = ", op)
	return op, nil
}

// func addition(bytes []byte) []byte {
// 	bits := binary.LittleEndian.Uint32(bytes[0:4])
// 	a := math.Float32frombits(bits)
// 	bits = binary.LittleEndian.Uint32(bytes[4:8])
// 	b := math.Float32frombits(bits)

// 	bits = math.Float32bits(a + b)
// 	sumbytes := make([]byte, 4)
// 	binary.LittleEndian.PutUint32(sumbytes, bits)
// 	return sumbytes
// }

func addition(op operands) float32 {
	return (op.A + op.B)
}

func main() {

	flag.Parse()

	if *jsonFlag {
		doc, _ := registryEntry.JSON()
		fmt.Println(string(doc))
		os.Exit(0)
	}

	log.SetFlags(0)
	if *debug {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(ioutil.Discard)
	}

	validateArgs()
	// Communication channels
	inaCh = make(chan bool)
	inbCh = make(chan bool)
	sumCh = make(chan bool)
	exitCh = make(chan os.Signal, 1)

	// Start the communication & processing logic
	fmt.Println("entering mainLoop")
	go mainLoop()
	fmt.Println("main loop done")
	// Wait for the end...
	signal.Notify(exitCh, os.Interrupt, syscall.SIGTERM)
	<-exitCh

	fmt.Println("Done")
}

func mainLoop() {
	openPorts()
	fmt.Println("ports opened")
	defer closePorts()
	waitCh := make(chan bool)
	go func() {
		total := 0
		fmt.Println("stop 5")
		for {
			select {
			case v := <-inaCh:
				fmt.Println("case v := <-inaCh:", v)
				if !v {
					fmt.Println("ina port is closed. Interrupting execution", v)
					exitCh <- syscall.SIGTERM
					break
				} else {
					total++
				}
			case v := <-inbCh:
				fmt.Println("case v := <-inbCh:", v)
				if !v {
					fmt.Println("inb port is closed. Interrupting execution", v)
					exitCh <- syscall.SIGTERM
					break
				} else {
					total++
				}
			case v := <-sumCh:
				fmt.Println("case v := <-sumCh:", v)
				if !v {
					fmt.Println("OUT port is closed. Interrupting execution", v)
					exitCh <- syscall.SIGTERM
					break
				} else {
					total++
				}
			}
			fmt.Println("total = ", total)
			if total >= 3 && waitCh != nil {
				waitCh <- true
			}
		}
	}()
	fmt.Println("Waiting for port connections to establish... ")

	var (
		ip [][]byte
	)

	select {
	case <-waitCh:
		fmt.Println("Ports connected")
		waitCh = nil
	case <-time.Tick(30 * time.Second):
		fmt.Println("Timeout: port connections were not established within provided interval")
		exitCh <- syscall.SIGTERM
		return
	}
	fmt.Println("Starting Add...")

	var OP operands
	for {
		ip, err = inaPort.RecvMessageBytes(0)
		if err != nil {
			continue
		}
		if !runtime.IsValidIP(ip) {
			fmt.Println("Invalid IP:", ip)
			continue
		}
		operand, _ := readjson(string(ip[1]))
		OP.A = operand.A

		ip, err = inbPort.RecvMessageBytes(0)
		if err != nil {
			continue
		}
		if !runtime.IsValidIP(ip) {
			fmt.Println("Invalid IP:", ip)
			continue
		}
		operand, _ = readjson(string(ip[1]))
		OP.B = operand.B

		sum := addition(OP)
		fmt.Println("Sum = ", sum)
		var buf []byte
		binary.BigEndian.PutUint32(buf[:], math.Float32bits(sum))
		sumPort.SendMessage(runtime.NewPacket(buf))
	}

}

// validateArgs checks all required flags
func validateArgs() {
	fmt.Println("inside validateArgs")
	if *inputaEndpoint == "" {
		fmt.Println("stop2")
		flag.Usage()
		os.Exit(1)
	}
	if *inputbEndpoint == "" {
		fmt.Println("stop3")
		flag.Usage()
		os.Exit(1)
	}
	if *sumEndpoint == "" {
		fmt.Println("stop4")
		flag.Usage()
		os.Exit(1)
	}
}

// openPorts create ZMQ sockets and start socket monitoring loops
func openPorts() {

	inaPort, err = utils.CreateInputPort("add.ina", *inputaEndpoint, inaCh)
	fmt.Println("inaPort", inaPort, err)
	utils.AssertError(err)

	inbPort, err = utils.CreateInputPort("add.inb", *inputbEndpoint, inbCh)
	fmt.Println("inbPort", inbPort, err)
	utils.AssertError(err)

	sumPort, err = utils.CreateOutputPort("add.sum", *sumEndpoint, sumCh)
	fmt.Println("sumPort", sumPort, err)
	utils.AssertError(err)

}

// closePorts closes all active ports and terminates ZMQ context
func closePorts() {
	fmt.Println("Closing ports...")
	inaPort.Close()
	fmt.Println("Port a closed...")
	inbPort.Close()
	fmt.Println("Port b closed...")
	sumPort.Close()
	fmt.Println("Port sum closed...")
	zmq.Term()
	fmt.Println("Terminated zmq default context")
}
