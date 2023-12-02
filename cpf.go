package br

import "errors"

// IsCPF returns true if cpf is a valid CPF.
func IsCPF(cpf string) bool {
	_, err := CanonicalCPF(cpf)
	return err == nil
}

var errInvalidCPF = errors.New("invalid CPF")

// CanonicalCPF verifies if a CPF is valid and returns
// the canonical CPF form (without "." or "-") or error.
func CanonicalCPF(cpf string) (string, error) {
	cpf = removeNotDigits(cpf)
	if !isValidCPF(cpf) {
		return "", errInvalidCPF
	}
	return cpf, nil
}

func isValidCPF(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}
	for _, d := range cpf {
		if !isDigit(d) {
			return false
		}
	}
	if isAllEqual(cpf) {
		return false
	}
	if hashDigit(cpf, 9) != cpf[9] || hashDigit(cpf, 10) != cpf[10] {
		return false
	}
	return true
}

func removeNotDigits(cpf string) string {
	if len(cpf) == 14 && cpf[3] == '.' && cpf[7] == '.' && cpf[11] == '-' {
		return cpf[:3] + cpf[4:7] + cpf[8:11] + cpf[12:]
	}
	return cpf
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func isAllEqual(cpf string) bool {
	first := cpf[0]

	for i := range cpf {
		if cpf[i] != first {
			return false
		}
	}

	return true
}

func hashDigit(cpf string, n int) byte {
	var sum int
	mul := n + 1
	for i := 0; i < n; i++ {
		v := int(cpf[i] - '0')
		sum += v * mul
		mul--
	}
	rem := sum % 11
	if rem < 2 {
		return '0'
	}

	return byte(11-rem) + '0'
}
