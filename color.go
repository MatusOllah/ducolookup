package main

import "github.com/fatih/color"

var (
	w  = color.New(color.FgWhite, color.Bold).SprintfFunc()
	y  = color.New(color.FgYellow, color.Bold).SprintfFunc()
	c  = color.New(color.FgCyan, color.Bold).SprintfFunc()
	m  = color.New(color.FgMagenta).SprintfFunc()
	mb = color.New(color.FgMagenta, color.Bold).SprintfFunc()
	r  = color.New(color.FgRed, color.Bold).SprintfFunc()
	g  = color.New(color.FgGreen, color.Bold).SprintfFunc()
)
