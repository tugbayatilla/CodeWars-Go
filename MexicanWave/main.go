package kata

func wave(s string) []string {
	wave := []string{}

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}

		sb := []byte(s)
		sb[i] = sb[i] + 'A' - 'a'
		wave = append(wave, string(sb))
	}
	return wave
}
