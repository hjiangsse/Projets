package main

import (
	pb "./tutorial"
	"fmt"
	proto "github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	p1 := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	p2 := pb.Person{
		Id:    1235,
		Name:  "hjiang",
		Email: "hjiang@sse.com.cn",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4320", Type: pb.Person_HOME},
		},
	}

	book := pb.AddressBook{
		People: []*pb.Person{
			&p1, &p2,
		},
	}

	//write the new address book to disk
	out, err := proto.Marshal(&book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

	if err := ioutil.WriteFile("./addressbook.txt", out, 0664); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	//read the address book back
	in, err := ioutil.ReadFile("./addressbook.txt")
	if err != nil {
		log.Fatalln("Failed to reading file:", err)
	}

	readed_book := &pb.AddressBook{}
	if err := proto.Unmarshal(in, readed_book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	fmt.Println(*readed_book)
}
