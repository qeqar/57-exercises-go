package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/crypto/bcrypt"
)

var users = map[string][]byte {
	"u1": []byte("$2y$05$YREkm3NWz2D2Et54Ouw9NuXWls./OX4tN7171ZaY1.Boq5NCPpDMC"), // test1
	"u2": []byte("$2y$05$vrhpPriSrCk2i1u7QoQHg.c.IDCptSrrQyygNf0SEWuvOEFLGgtXG"), // test2
	"u3": []byte("$2y$05$7/d5fjOygqZgux.vfxvU7u7FF0pRDlfnfZ4mrfvyR3nqnVpz6PbJq"), // test3
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("User? ")
	usernameRead, _ := reader.ReadString('\n')
	username := strings.TrimSuffix(usernameRead, "\n")

	fmt.Print("Password? ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println("")

	err := bcrypt.CompareHashAndPassword(users[username], bytePassword)

	if err == nil {
		fmt.Println("Welcome")
	} else {
		fmt.Printf("I don't know you => %s\n",err)
	}
}

