package communication

import (
	"fmt"

	pb "github.com/ajaypp123/sample_grpc/proto/service"
)

var empMap map[string]*pb.Employee

func initEmpDB() {
	empMap = make(map[string]*pb.Employee)
	empMap["1"] = &pb.Employee{
		EmployeeId: "1",
		Name:       "Emp1",
		Salary:     100,
		Company:    "MyCompany",
	}
	empMap["2"] = &pb.Employee{
		EmployeeId: "2",
		Name:       "Emp2",
		Salary:     100,
		Company:    "MyCompany",
	}
	empMap["3"] = &pb.Employee{
		EmployeeId: "3",
		Name:       "Emp3",
		Salary:     100,
		Company:    "MyCompany",
	}
	empMap["4"] = &pb.Employee{
		EmployeeId: "4",
		Name:       "Emp4",
		Salary:     100,
		Company:    "MyCompany",
	}
	empMap["5"] = &pb.Employee{
		EmployeeId: "5",
		Name:       "Emp5",
		Salary:     100,
		Company:    "MyCompany",
	}

	fmt.Println(empMap)
}
