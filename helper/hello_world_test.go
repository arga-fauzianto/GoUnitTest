package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("fauzi")

	if result != "Hello fauzi" {
		//err

		panic("is not 'fauzi'")
	}

}

func TestHelloWorldArga(t *testing.T) {
	result := HelloWorld("Arga")

	if result != "Hello Arga" {
		//err

		panic("is not 'arga'")
	}

}
