package conv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// KToC converts a Fahrenheit temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// KgToLb converts a Kilograms to Pounds
func KgToLb(k Kg) Lb { return Lb(k / KgInLb) }

// LbToKg converts a Pounds to Kilograms
func LbToKg(l Lb) Kg { return Kg(l * KgInLb) }

// MeterToFt converts a Meters to Foots
func MeterToFt(m Meter) Ft { return Ft(m / MetersInFt) }

// FtToMeter converts a Foots to Meters
func FtToMeter(f Ft) Meter { return Meter(f * MetersInFt) }
