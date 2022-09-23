package types

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		FortunesList: []Fortune{
			{
				Owner:   "",
				Price:   "100token",
				Fortune: "Trust, but still keep your eyes open.",
			},
			{
				Owner:   "",
				Price:   "100token",
				Fortune: "Unnecessary possessions are unnecessary burdens.",
			},
			{
				Owner:   "",
				Price:   "100token",
				Fortune: "A person is not wise simply because one talks a lot.",
			},
			{
				Owner:   "",
				Price:   "100token",
				Fortune: "Enjoy the small things you find on your path.",
			},
			{
				Owner:   "",
				Price:   "100token",
				Fortune: "You have a heart of gold.",
			},
			{
				Owner:   "",
				Price:   "100token",
				Fortune: "Some people never have anything except ideas, Go do it.",
			},
			{
				Owner:   "",
				Price:   "100token",
				Fortune: "Pure logic is the ruin of the spirit.",
			},
			{
				Owner:   "",
				Price:   "100token",
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
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
