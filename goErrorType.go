package main

import "fmt"

type PasswordError struct {
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string {
	return "암호 길이가 짧습니다."
}

func RegisterAccount(name, password string) error {
	if len(password) <= 0 {
		return PasswordError{len(password), 0}
	}

	return nil
}

func printErrorType() {
	err := RegisterAccount("myID", "")
	if err != nil {
		if errInfo, ok := err.(PasswordError); ok {
			fmt.Printf("%v Len:%d RequirenLen: %d\n",
				errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("회원가입 했습니다.")
	}
}

func errorTypeEx() {
	printErrorType()
}
