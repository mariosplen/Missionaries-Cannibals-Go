package main

type movement struct {
	nMiss   int
	nCann   int
	towards pos
}

var moves = map[string]movement{
	"2M_LR": {
		nMiss:   2,
		nCann:   0,
		towards: right,
	},
	"2C_LR": {
		nMiss:   0,
		nCann:   2,
		towards: right,
	},
	"1M1C_LR": {
		nMiss:   1,
		nCann:   1,
		towards: right,
	},
	"1M_LR": {
		nMiss:   1,
		nCann:   0,
		towards: right,
	},
	"1C_LR": {
		nMiss:   0,
		nCann:   1,
		towards: right,
	},
	"2M_RL": {
		nMiss:   2,
		nCann:   0,
		towards: left,
	},
	"2C_RL": {
		nMiss:   0,
		nCann:   2,
		towards: left,
	},
	"1M1C_RL": {
		nMiss:   1,
		nCann:   1,
		towards: left,
	},
	"1M_RL": {
		nMiss:   1,
		nCann:   0,
		towards: left,
	},
	"1C_RL": {
		nMiss:   0,
		nCann:   1,
		towards: left,
	},
}
