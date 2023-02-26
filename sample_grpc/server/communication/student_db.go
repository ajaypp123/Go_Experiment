package communication

import (
	"fmt"

	pb "github.com/ajaypp123/sample_grpc/proto/service"
)

var stdMap map[string]*pb.Student

func initStdDB() {
	stdMap = make(map[string]*pb.Student)
	stdMap["1"] = &pb.Student{
		StudentId: "1",
		Name:      "Std1",
		Salary:    100,
		Company:   "MyCompany",
	}
	stdMap["2"] = &pb.Student{
		StudentId: "2",
		Name:      "Std2",
		Salary:    100,
		Company:   "MyCompany",
	}
	stdMap["3"] = &pb.Student{
		StudentId: "3",
		Name:      "Std3",
		Salary:    100,
		Company:   "MyCompany",
	}
	stdMap["4"] = &pb.Student{
		StudentId: "4",
		Name:      "Std4",
		Salary:    100,
		Company:   "MyCompany",
	}
	stdMap["5"] = &pb.Student{
		StudentId: "5",
		Name:      "Std5",
		Salary:    100,
		Company:   "MyCompany",
	}

	fmt.Println(stdMap)
}
