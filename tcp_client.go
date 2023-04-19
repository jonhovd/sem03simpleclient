package main

import (
	"log"
	"net"
	"os"

	"github.com/jonhovd/is105sem03/mycrypt"
)

func main() {
	conn, err := net.Dial("tcp", "172.17.0.3:40000")
	if err != nil {
		log.Fatal(err)
	}

	// Krypter meldingen fra kommandolinjen
	kryptertMelding := mycrypt.Krypter([]rune(os.Args[1]), mycrypt.ALF_SEM03, 4)
	log.Println("Kryptert melding: ", string(kryptertMelding))

	// Send kryptert melding til serveren
	_, err = conn.Write([]byte(string(kryptertMelding)))
	if err != nil {
		log.Fatal(err)
	}

	// Motta kryptert svar fra serveren
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	kryptertSvar := string(buf[:n])

	// Dekrypter meldingen fra serveren
	dekryptertSvar := mycrypt.Krypter([]rune(kryptertSvar), mycrypt.ALF_SEM03, len(mycrypt.ALF_SEM03)-4)
	log.Println("Dekryptert svar: ", string(dekryptertSvar))

	// Skriv ut svar fra serveren
	log.Printf("Svar fra proxy: %s", string(dekryptertSvar))
}
