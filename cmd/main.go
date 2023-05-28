package main

import "heji-server/utils"

var log = utils.Log

func main() {
	log.Println("Start heji server")
	pinter()

}
func pinter() {
	log.Println("this is pinter func")
}
