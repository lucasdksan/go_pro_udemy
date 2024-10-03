package examples

import (
	"bufio"
	"fmt"
	"net"
)

func Exec_1() {
	ls, err := net.Listen("tcp", ":5000")

	if err != nil {
		panic(err)
	}

	defer ls.Close()

	for {
		con, err := ls.Accept()

		if err != nil {
			panic(err)
		}

		go func(con net.Conn) {
			data, _ := bufio.NewReader(con).ReadString('\n')

			fmt.Println("Dado recebido: ", data)

			con.Write([]byte("Sua mensagem foi recebida com sucesso"))
			con.Close()
			fmt.Println("Conexão encerrada.")
		}(con)
	}
}
