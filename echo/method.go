package main

//go:generate stringer -type=Method method.go
type Method int

const (
	GET Method = iota
	POST
)
