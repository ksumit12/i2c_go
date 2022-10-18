package main

import (
	"log"
	"time"

	"periph.io/x/host/v3"

	"periph.io/x/conn/v3/i2c"

	"periph.io/x/conn/v3/i2c/i2creg"
)

func main() {

	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Host initialized\n")

	b, err := i2creg.Open("1")
	//time.Sleep(500 * time.Millisecond)
	if err != nil {
		log.Fatal(err)
	}
	defer b.Close()

	log.Printf("I2C opened\n")

	// Dev is a valid conn.Conn.
	d := &i2c.Dev{Addr: 04, Bus: b}

	log.Printf("Devie created\n")

	// Send a command 0x10 and expect a 5 bytes reply.
	for {
		write := []byte{0x16}
		read := make([]byte, 5)

		if err := d.Tx(write, read); err != nil {
			//log.Println(err)
		}

		log.Printf("%v\n", write)
		log.Printf("%v\n", read)
		time.Sleep(100 * time.Millisecond)
	}

}
