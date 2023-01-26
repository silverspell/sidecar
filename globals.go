package main

import "sync"

var (
	HOST string = "0.0.0.0"
	PORT string = ":9001"
	TYPE string = "tcp"
	mut  sync.RWMutex
	once sync.Once
)
