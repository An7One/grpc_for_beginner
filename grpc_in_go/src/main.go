package main

import (
	"fmt"
	"io/ioutil"
	"log"

	firstpb "github.com/an7one/grpc_for_beginner/src/first"
	enumpb "github.com/an7one/grpc_for_beginner/src/second"
	complexpb "github.com/an7one/grpc_for_beginner/src/third"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	pm := NewPersonMessage()
	err := writeToFile("person.bin", pm)
	if err != nil {
		log.Fatalln("Failed with writing to the file", err.Error())
	}

	pm2 := &firstpb.PersonMessage{}
	err = readFromFile("person.bin", pm2)
	if err != nil {
		log.Fatalln("Failed with reading from the file", err.Error())
	}
	fmt.Println(pm2)

	pmStr := toJSON(pm)
	fmt.Println(pmStr)

	pm3 := &firstpb.PersonMessage{}
	err = fromJSON(pmStr, pm3)
	if err != nil {
		log.Fatalln("Failed with unmarshaling the message", err.Error())
	}
	fmt.Println(pmStr)

	em := NewEnumMessage()
	fmt.Println(enumpb.Gender_name[int32(em.Gender)])

	dm := NewDepartmentMessage()
	fmt.Println(dm)
}

func NewDepartmentMessage() *complexpb.DepartmentMessage {
	dm := &complexpb.DepartmentMessage{
		Id:   5678,
		Name: "R&D",
		Employees: []*complexpb.EmployeeMessage{
			{
				Id:   11,
				Name: "Dave",
			},
			{
				Id:   22,
				Name: "Mike",
			},
		},
		ParentDepartment: &complexpb.DepartmentMessage{
			Id:   1122,
			Name: "Headquarter",
		},
	}

	return dm
}

func NewEnumMessage() *enumpb.EnumMessage {
	em := enumpb.EnumMessage{
		Id:     345,
		Gender: enumpb.Gender_FEMALE,
	}
	return &em
}

func fromJSON(in string, pb proto.Message) error {
	mo := protojson.UnmarshalOptions{}
	err := mo.Unmarshal([]byte(in), pb)
	if err != nil {
		return err
	}
	return nil
}

func toJSON(pb proto.Message) string {
	mo := protojson.MarshalOptions{
		Indent:          "    ",
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}
	bytes, err := mo.Marshal(pb)
	if err != nil {
		log.Fatalln("Failed with converting to JSON", err.Error())
	}
	return string(bytes)
}

func readFromFile(fileName string, pb proto.Message) error {
	dataBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = proto.Unmarshal(dataBytes, pb)
	if err != nil {
		return err
	}

	log.Println("Successfully read from the file")
	return nil
}

func writeToFile(fileName string, pb proto.Message) error {
	dataBytes, err := proto.Marshal(pb)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(fileName, dataBytes, 0644); err != nil {
		return err
	}

	log.Println("Successfully write to the file")
	return nil
}

func NewPersonMessage() *firstpb.PersonMessage {
	pm := firstpb.PersonMessage{
		Id:          1234,
		IsAdult:     true,
		Name:        "Dave",
		LuckNumbers: []int32{1, 2, 3, 4, 5},
	}

	fmt.Println(pm)

	pm.Name = "Nick"

	fmt.Println(pm)

	fmt.Printf("The ID is: %d\n", pm.GetId())

	return &pm
}
