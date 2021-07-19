package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main()  {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name: ")
	name, _ := reader.ReadString('\n')
	name = name[:len(name) - 1]
	fmt.Println("What is your email: ")
	email, _ := reader.ReadString('\n')
	email = email[:len(email) - 1]

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:7000", nil)
	req.Header.Set("name", name)
	req.Header.Set("email", email)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response status:", res.Status)

	scanner := bufio.NewScanner(res.Body)
	for i := 0; scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

}

