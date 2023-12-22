package main

import (
	"strings"

	"github.com/thoas/go-funk"
)

const BROADCASTER = "broadcaster"
const BUTTON = "button"

const TYPE_BROADCASTER = 'b'
const TYPE_CONJUNCTION = '&'
const TYPE_FLIPFLOP = '%'

type Pulse struct {
	value   byte
	origin  string
	destiny string
}

type module struct {
	name    string
	tipe    byte
	inputs  map[string]byte
	outputs []string
	last    string
	status  bool
}

func (pulse Pulse) handlePulse(modules *map[string]module) []Pulse {
	var out byte = 0
	m := (*modules)[pulse.destiny]

	switch m.tipe {
	case TYPE_BROADCASTER:
		out = pulse.value
	case TYPE_FLIPFLOP:
		m.inputs[pulse.origin] = pulse.value
		m.last = pulse.origin

		if m.inputs[m.last] == 0 {
			m.status = !m.status
			out = Ternary[byte](1, 0, m.status)
		} else {
			return []Pulse{}
		}
	case TYPE_CONJUNCTION:
		m.inputs[pulse.origin] = pulse.value
		m.last = pulse.origin

		out = 0
		for _, i := range m.inputs {
			if i == 0 {
				out = 1
				break
			}
		}
	}

	(*modules)[pulse.destiny] = m

	res := funk.Reduce(m.outputs, func(acc []Pulse, cur string) []Pulse {
		return append(acc, Pulse{value: out, origin: m.name, destiny: cur})
	}, []Pulse{}).([]Pulse)

	return res
}

func solve_day_20_1(url string) int {
	lines := readLinesFromFile(url)
	modules := getModules(&lines)

	l, h := lowAndHighs(modules, 1000)

	return l * h
}

func solve_day_20_2(url string) int {
	lines := readLinesFromFile(url)
	modules := getModules(&lines)
	// assuming we have a conjunction before our end module (RX)
	conj := (*getInputs(modules, "rx"))[0]
	inputsForConj := getInputs(modules, conj)

	return neededButtonPresses(modules, inputsForConj, conj)
}

func getInputs(modules *map[string]module, inputsForConj string) *[]string {
	base := []string{}
	for k, m := range *modules {
		if funk.Contains(m.outputs[0], inputsForConj) {
			base = append(base, k)
		}
	}
	return &base
}

func neededButtonPresses(modules *map[string]module, inputs *[]string, conj string) int {
	cycles := make(map[string]int)
	pulses := make([]Pulse, 0)

	for i := 1; len(cycles) != len(*inputs); i++ {
		pulses = pressButton(modules)

		for len(pulses) > 0 {
			pulse := pulses[0]
			pulses = append(pulses[1:], pulse.handlePulse(modules)...)

			for _, k := range *inputs {
				_, ok := cycles[k]
				if !ok && (*modules)[conj].inputs[k] == 1 {
					cycles[k] = i
				}
			}
		}
	}

	return LCM(cycles[(*inputs)[0]], cycles[(*inputs)[1]], cycles[(*inputs)[2]], cycles[(*inputs)[3]])
}

func lowAndHighs(modules *map[string]module, times int) (int, int) {
	low, high := 0, 0
	pulses := []Pulse{}

	for i := 0; i < times; i++ {
		pulses = pressButton(modules)

		for len(pulses) > 0 {
			pulse := pulses[0]
			pulses = append(pulses[1:], pulse.handlePulse(modules)...)

			if pulse.value == 0 {
				low++
			} else {
				high++
			}

		}
	}

	return low, high
}

func getModules(lines *[]string) *map[string]module {
	modules := make(map[string]module, len(*lines))
	for _, l := range *lines {
		description := strings.Split(l, " -> ")
		name := Ternary[string](BROADCASTER, strings.TrimSpace(description[0])[1:], description[0][0] == TYPE_BROADCASTER)
		tipe := description[0][0]

		outputs := strings.Split(description[1], ",")
		for i, o := range outputs {
			outputs[i] = strings.TrimSpace(o)
		}

		modules[name] = module{name: name, tipe: tipe, inputs: make(map[string]byte, len(*lines)-1), outputs: outputs}
	}

	for k, m := range modules {
		for _, t := range m.outputs {
			if modules[t].tipe == TYPE_CONJUNCTION {
				modules[t].inputs[k] = 0
			}
		}
	}

	modules[BUTTON] = module{name: BUTTON, outputs: []string{BROADCASTER}}
	return &modules
}

func pressButton(modules *map[string]module) []Pulse {
	return []Pulse{{origin: BUTTON, destiny: BROADCASTER, value: 0}}
}
