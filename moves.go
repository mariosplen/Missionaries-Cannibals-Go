package main

type move struct {
	nMiss     int  // number of missionaries moving
	nCann     int  // number of canibals moving
	leftwards bool // type of movement
}

// All the posible moves that can be applied to a state
var moves = map[string]move{
	"2M_LR": {
		nMiss:     2,
		nCann:     0,
		leftwards: false,
	},
	"2C_LR": {
		nMiss:     0,
		nCann:     2,
		leftwards: false,
	},
	"1M1C_LR": {
		nMiss:     1,
		nCann:     1,
		leftwards: false,
	},
	"1M_LR": {
		nMiss:     1,
		nCann:     0,
		leftwards: false,
	},
	"1C_LR": {
		nMiss:     0,
		nCann:     1,
		leftwards: false,
	},
	"2M_RL": {
		nMiss:     2,
		nCann:     0,
		leftwards: true,
	},
	"2C_RL": {
		nMiss:     0,
		nCann:     2,
		leftwards: true,
	},
	"1M1C_RL": {
		nMiss:     1,
		nCann:     1,
		leftwards: true,
	},
	"1M_RL": {
		nMiss:     1,
		nCann:     0,
		leftwards: true,
	},
	"1C_RL": {
		nMiss:     0,
		nCann:     1,
		leftwards: true,
	},
}
