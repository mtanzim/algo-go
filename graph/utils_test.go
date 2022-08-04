package graph

import (
	"bufio"
	"os"
)

// unweighted, undirected for now
func GraphFromFile(filepath string) (*Graph[UnweightedEdgeType], error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	g, err := NewUndirectedFromScanner(scanner)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func WeightedGraphFromFile(filepath string) (*EdgeWeightedGraph, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	g, err := NewWeightedFromScanner(scanner)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func DigraphFromFile(filepath string) (*Graph[UnweightedEdgeType], error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	g, err := NewDigraphFromScanner(scanner)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func WeightedDigraphFromFile(filepath string) (*EdgeWeightedDigraph, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	g, err := NewWeightedDigraphFromScanner(scanner)
	if err != nil {
		return nil, err
	}
	return g, nil
}
