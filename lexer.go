package main

type address struct{
	street, city, state string
}

func parseAddresses(str string) (map[int]address){
	addressList := make(map[int]address)
	foo := address{}
	a := ""
	b := 0
	c := 0
	for j := range str{
		if (string(str[j]) == "`"){
			if (b == 1){
				foo.street = a
				a = ""
			}else if (b == 2){
				foo.city = a
				a = ""
			}else if (b == 3){
				foo.state = a
				a = ""
			}
			b++
		}else if (string(str[j]) == "~"){
			foo.state = a
			a = ""
			addressList[c] = foo
			b = 0
			c++
		}else if (string(str[j]) == " "){
			a = a + "+"	
		}else{
			a = a + string(str[j])
		}
	}
	
	return addressList
}