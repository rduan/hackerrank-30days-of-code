

func (p person) NewPerson(initialAge int) person {
	//Add some more code to run some checks on initialAge
	p.age = initialAge

	return p
}

func (p *person) amIOld() {
	//Do some computation in here and print out the correct statement to the console

	if p.age < 0 {
		p.age = 0
		fmt.Println("Age is not valid, setting age to 0.")
		fmt.Println("You are young.")
	} else if p.age < 13 {
		fmt.Println("You are young.")
	} else if p.age < 18 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are old.")
	}
}

func (p person) yearPasses() person {
	//Increment the age of the person in here
	p.age += 1
	return p
}
