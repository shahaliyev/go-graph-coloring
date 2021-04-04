package main

// Global declarations
var graph = make(map[int][]int)
var domains = make(map[int][]int)

var colors = make(map[int]int)
var used = make(map[int]int)

var colorCount int
var domainCount int

const INF = 0x3f3f3f3
