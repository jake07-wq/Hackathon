package piscine

func Enigma(a ***int, b *int, c *******int, d ****int) {
	// Store current values in temporary variables
	valA := ***a
	valB := *b
	valC := *******c
	valD := ****d

	// Move values according to the rules:
	// a -> c, c -> d, d -> b, b -> a
	*******c = valA
	****d = valC
	*b = valD
	***a = valB
}
