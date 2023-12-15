package main

// https://exercism.org/tracks/go/exercises/protein-translation/edit

import "fmt"

var CodonProteinPairs = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func splitByCodon(rna string) []string {
	var sliceOfCodons []string
	for i := 0; i < len(rna); i += 3 {
		end := i + 3
		sliceOfCodons = append(sliceOfCodons, rna[i:end])
	}
	fmt.Printf("sliceOfCodons: %s", sliceOfCodons)
	return sliceOfCodons
}

func shouldStop(codon string) bool {
	if codon == "UAA" || codon == "UAG" || codon == "UGA" {
		return true
	}
	return false
}

func FromRNA(rna string) ([]string, error) {
	var sliceOfProteins []string
	theCodons := splitByCodon(rna)
	for _, codon := range theCodons {
		if shouldStop(codon) {
			break
		}
		value, exists := CodonProteinPairs[codon]
		if exists {
			sliceOfProteins = append(sliceOfProteins, value)
		}
	}
	return sliceOfProteins, nil
}

func FromCodon(codon string) (string, error) {
	result := CodonProteinPairs[codon]
	return result, nil
}
