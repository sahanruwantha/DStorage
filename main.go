package main

import (
	"log"
	"github.com/sahanruwantha/DStorage/p2p"
)

func main()  {
	tr := p2p.NewTCPTranport(":3000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}