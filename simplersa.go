package main

import ("fmt")

func ifprime(a float64) bool {
	var b int64
	b = int64(a)
	if b == 2 {
		return true
	} else if b > 2 {
		var i int64
		for i = 2; i < b; i++ {
			if b % i == 0 {
				return false
			}
		}
	}
	return true
}

func decription(phi float64, e float64 ) float64 {
	x := 0.0
	for {
		var d float64
		d = (1.0 + (x * phi))/e
		g := float64(int(d))
		if g == d {
			return d
		}
		x = x + 1.0 
	}
	return e
}

func cipher (m int, e int, n int) int {
    ciph := 1
    x := 1
    for x <= e {
        ciph = ciph * m
        x = x + 1
    }
    return ciph % n
}


func decipher (c int, d int, n int) int {
    deciph := 1
    x := 1
    for x <= d {
        deciph = deciph * c
        x = x + 1
    }
    return deciph % n
}


func main() {
	// Getting p number
	var p float64

	fmt.Print("Enter the prime number 'p':")
	fmt.Scanln(&p)
	if ifprime(p) == false {
		fmt.Println("Please use the PRIME number!!!")
		return
	}

	// Getting q number
	var q float64

	fmt.Print("Enter the prime number 'q':")
	fmt.Scanln(&q)
	if ifprime(q) == false {
		fmt.Println("Please use the PRIME number!!!")
		return
	} else if q == p {
		fmt.Println("Please use the different from p number!!!")
		return
	}
	
	// Getting n number
	var n float64

	n = p*q
	fmt.Println("This is the 'n':", n)
	
	// Getting phi number
	var phi float64

	phi = (p-1)*(q-1)
	fmt.Println("This is the 'phi':", phi)

	// Getting e number
	var e float64

	fmt.Print("Enter the prime number 'e':")
	fmt.Scanln(&e)
	if e <= 1 || e >= phi {
		fmt.Println("Please use the 'e' which is 1 < e < phi!!!")
		return
	} else if ifprime(e) == false {
		fmt.Println("Please use the PRIME number!!!")
		return	
	} else if int(phi) % int (e) == 0 {
		fmt.Println("Please use the COPRIME number!! ")
		return
	}

	// Getting the d number
	var d float64

	d = decription(phi, e)
	fmt.Println("This is the 'd':", d)

	// Getting the message
	var m int

	fmt.Print("Enter the m num:")
	fmt.Scanln(&m)

	//Getting the cipher

	ciph := cipher(int(m), int(e), int(n))
	fmt.Println("This is the cipthertext:", ciph)

	//Getting the message back
	deciph := decipher(ciph, int(d), int(n))
	fmt.Println("This is the deciphered text:", deciph)
}

