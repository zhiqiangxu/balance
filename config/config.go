package config

// Account ...
type Account struct {
	Eth     []string
	BSC     []string
	Heco    []string
	Polygon []string
	OK      []string
	Ont     []string
	Zil     []string
	Neo     []string
}

// Threshold ...
type Threshold struct {
	Eth     uint
	BSC     uint
	Heco    uint
	Polygon uint
	Ont     uint
	Zil     uint
	Neo     uint
	OK      uint
}

// Node ...
type Node struct {
	Eth     []string
	BSC     []string
	Heco    []string
	OK      []string
	Ont     []string
	Zil     []string
	Neo     []string
	Polygon []string
}

// Config ...
type Config struct {
	Account   Account
	Threshold Threshold
	Node      Node
}
