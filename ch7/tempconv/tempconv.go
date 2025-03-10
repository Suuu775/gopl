package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const AbsoluteZeroC Celsius = -273.15

func KToC(f Kelvin) Celsius { return Celsius(f) + AbsoluteZeroC }

type celsiusFlag struct{ Celsius }

// String implements flag.Value.
func (f *celsiusFlag) String() string {
	return fmt.Sprintf("%f", f.Celsius)
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
