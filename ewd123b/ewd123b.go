/*
 * ewd123b.go
 *
 * A program to represent the second mutex strategy, as described in EWD123.
 *
 * Copyright (c) 2019-2019 HS Emden/Leer
 * All Rights Reserved.
 *
 * version 1.00 - 21 Oct 2019 - GJV - initial version
 *
 * author: Gert Veltink, gert.veltink@hs-emden-leer.de (GJV)
 */

package ewd123b

import (
	"log"

	"./controller"
)

// global synchronization variables
var c1, c2 = 1, 1

// Start starts the execution of EWD123b.
func Start() {
	go process1()
	go process2()
}

// process1 simulates the behaviour of the first process
func process1() {

L1:
	if c2 == 0 {
		log.Printf("Process 1 waiting\n")
		goto L1
	}

	c1 = 0

	controller.EnterCriticalSection(1)
	controller.InsideCriticalSection(1, 1)
	controller.LeaveCriticalSection(1)

	c1 = 1

	controller.OutsideCriticalSection(1, 1)

	if controller.ProcessCrashed(0.000001) {
		log.Printf("Process 1 crashed\n")
		return
	}

	goto L1

}

// process2 simulates the behaviour of the second process
func process2() {

L2:
	if c1 == 0 {
		log.Printf("Process 2 waiting\n")
		goto L2
	}

	c2 = 0

	controller.EnterCriticalSection(2)
	controller.InsideCriticalSection(2, 1)
	controller.LeaveCriticalSection(2)

	c2 = 1

	controller.OutsideCriticalSection(2, 1)

	if controller.ProcessCrashed(0.000001) {
		log.Printf("Process 2 crashed\n")
		return
	}

	goto L2

}
