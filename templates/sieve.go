const siz = 10000000
var primes []int
var siv [siz]int
func sieve() {
	primes = append(primes, 2)
	for i := 3; i*i < siz; i += 2 {
		if siv[i] == 0 {
			for j := i*i; j < siz; j += 2*i {
				siv[j] = 1
			}
		}
	}
	for i := 3; i < siz; i += 2 {
		if siv[i] == 0 {
			primes = append(primes, i)
		}
	}
}
