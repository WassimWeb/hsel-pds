/*
 * ewd123dekker.go
 *
 * A program to represent the Dekker mutex strategy, as described in EWD123.
 *
 * Copyright (c) 2019-2019 HS Emden/Leer
 * All Rights Reserved.
 *
 * version 1.00 - 23 Oct 2019 - GJV - initial version
 *
 * author: Gert Veltink, gert.veltink@hs-emden-leer.de (GJV)
 */

package ewd123dekker

import (
	"log"

	"./controller"
)

// global synchronization variables
var c1, c2 = 1, 1
var turn = 1

// Start starts the execution of Dekker's solution form EWD123.
func Start() {
	go process1()
	go process2()
}

// process1 simulates the behaviour of the first process
func process1() {

A1:
	c1 = 0
L1:
	if c2 == 0 {
		if turn == 1 {

			goto L1

		}
		c1 = 1
	B1:
		if turn == 2 {

			goto B1
		}
		log.Printf("Process 1 waiting\n")
		goto A1
	}

	controller.EnterCriticalSection(1)
	controller.InsideCriticalSection(1, 50)
	controller.LeaveCriticalSection(1)

	turn = 2
	c1 = 1

	controller.OutsideCriticalSection(1, 100)

	if controller.ProcessCrashed(0.00001) {
		log.Printf("Process 1 crashed\n")
		return
	}

	goto A1

}

// process2 simulates the behaviour of the second process
func process2() {

A2:
	c2 = 0
L2:
	if c1 == 0 {
		if turn == 2 {

			goto L2

		}
		c2 = 1
	B2:
		if turn == 1 {

			goto B2
		}
		log.Printf("Process 2 waiting\n")
		goto A2
	}

	controller.EnterCriticalSection(2)
	controller.InsideCriticalSection(2, 50)
	controller.LeaveCriticalSection(2)

	turn = 1
	c2 = 1

	controller.OutsideCriticalSection(2, 100)

	if controller.ProcessCrashed(0.00001) {
		log.Printf("Process 2 crashed\n")
		return
	}

	goto A2
}
