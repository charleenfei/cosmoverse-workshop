package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		FortunesList: []Fortune{
			{
				Owner:   "",
				Price:   "",
				Fortune: "Trust, but sill keep your eyes open.",
			},
			{
				Owner:   "",
				Price:   "",
				Fortune: "Unnecessary possessions are unnecessary burdens.",
			},
			{
				Owner:   "",
				Price:   "",
				Fortune: "A person is not wise simply because one talks a lot.",
			},
			{
				Owner:   "",
				Price:   "",
				Fortune: "Enjoy the small things you find on your path.",
			},
			{
				Owner:   "",
				Price:   "",
				Fortune: "You have a heart of gold.",
			},
			{
				Owner:   "",
				Price:   "",
				Fortune: "Some people never have anything except ideas, Go do it.",
			},
			{
				Owner:   "",
				Price:   "",
				Fortune: "Pure logic is the ruin of the spirit.",
			},
			{
				Owner:   "",
				Price:   "",
				Fortune: "When the moment comes, take the one from the right.",
			},
		},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in fortunes
	fortunesIndexMap := make(map[string]struct{})

	for _, elem := range gs.FortunesList {
		index := string(FortuneKey(elem.Owner))
		if _, ok := fortunesIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for fortunes")
		}
		fortunesIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
