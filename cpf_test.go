package br

import (
	"testing"
)

func TestCanonicalCPF(t *testing.T) {
	tests := []struct {
		name    string
		cpf     string
		wantCPF string
		wantErr bool
	}{
		{name: "valid", cpf: "10425095142", wantCPF: "10425095142", wantErr: false},
		{name: "valid ending with 0", cpf: "880.600.410-70", wantCPF: "88060041070", wantErr: false},
		{name: "invalid", cpf: "10425095143", wantCPF: "", wantErr: true},
		{name: "valid with .", cpf: "104.250.951-42", wantCPF: "10425095142", wantErr: false},
		{name: "valid with wrong .", cpf: "104250.951-42", wantCPF: "", wantErr: true},
		{name: "invalid all equal", cpf: "222.222.222-22", wantCPF: "", wantErr: true},
		{name: "invalid with letter", cpf: "10A.250.951-42", wantCPF: "", wantErr: true},
		{name: "invalid with emoji", cpf: "1😄095142", wantCPF: "", wantErr: true},
		{name: "invalid place of .", cpf: "1042509.4090", wantCPF: "", wantErr: true},
		{name: "invalid format", cpf: "12345--678909", wantCPF: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpf, err := CanonicalCPF(tt.cpf)
			if cpf != tt.wantCPF || (err != nil) != tt.wantErr {
				t.Errorf("got %q %v want %q %v", cpf, err, tt.wantCPF, tt.wantErr)
			}
		})
	}
}

func TestIsValidCPF(t *testing.T) {
	if !IsCPF("104.250.951-42") {
		t.Fatal("should be valid CPF")
	}
	if IsCPF("204.250.951-42") {
		t.Fatal("should be invalid CPF")
	}
}

func BenchmarkIsValidCPF(b *testing.B) {
	b.ReportAllocs()

	cpf := "104.250.951-42"
	for i := 0; i < b.N; i++ {
		IsCPF(cpf)
	}
}
