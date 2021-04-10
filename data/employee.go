package data

type Employee struct{
	Id, Salary, Age int
	Name, Destigination string
}

func (emp *Employee) SetId(id int){
	emp.Id= id
}

func (emp *Employee) SetName(name string){
	emp.Name = name
}

func (emp *Employee) SetAge(age int){
	emp.Age = age
}
func (emp *Employee) SetDestigination(destigination string){
	emp.Destigination = destigination
}
func (emp *Employee) SetSalary(salary int){
	emp.Salary = salary
}


func (emp *Employee) GetId() int{
	return emp.Id
}
func (emp *Employee) GetName() string{
	return emp.Name
}
func (emp *Employee) GetAge() int{
	return emp.Age
}
func (emp *Employee) GetDestigination() string{
	return emp.Destigination
}
func (emp *Employee) GetSalary() int{
	return emp.Salary
}
