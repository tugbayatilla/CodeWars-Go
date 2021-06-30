# CodeWars-Go

Codewars with golang

## Mexican Wave

[Mexican Wave Link](https://www.codewars.com/kata/58f5c63f1e26ecda7e000029/go)

```go
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
```
