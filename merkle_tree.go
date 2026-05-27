// Copyright 2017 Cameron Bergoon
// Licensed under the MIT License, see LICENCE file for details.

package merkletree

import (
	"hash"
)

// Content represents the data that is stored and verified by the tree. A type that
// implements this interface can be used as an item in the tree.
type Content interface {
	CalculateHash() ([]byte, error)
	Equals(other Content) (bool, error)
}

// MerkleTree is the container for the tree. It holds a pointer to the root of the tree,
// a list of pointers to the leaf nodes, and the merkle root.
type MerkleTree struct {
	Root         *Node
	merkleRoot   []byte
	Leafs        []*Node
	hashStrategy func() hash.Hash
	sort         bool
}

// Node represents a node, root, or leaf in the tree. It stores pointers to its immediate
// relationships, a hash, the content stored if it is a leaf, and other metadata.
type Node struct {
	Tree   *MerkleTree
	Parent *Node
	Left   *Node
	Right  *Node
	leaf   bool
	dup    bool
	Hash   []byte
	C      Content
	sort   bool
}

// sortAppend sort and append the nodes to be compatible with OpenZepplin libraries
// https://github.com/OpenZeppelin/openzeppelin-contracts-ethereum-package/blob/master/contracts/cryptography/MerkleProof.sol
func sortAppend(sort bool, a, b []byte) []byte { _ = "STUB: not implemented"; return nil }

// verifyNode walks down the tree until hitting a leaf, calculating the hash at each level
// and returning the resulting hash of Node n.
func (n *Node) verifyNode(sort bool) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// calculateNodeHash is a helper function that calculates the hash of the node.
func (n *Node) calculateNodeHash(sort bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewTree creates a new Merkle Tree using the content cs.
func NewTree(cs []Content) (*MerkleTree, error) { _ = "STUB: not implemented"; return nil, nil }

// NewTreeWithHashStrategy creates a new Merkle Tree using the content cs using the provided hash
// strategy. Note that the hash type used in the type that implements the Content interface must
// match the hash type provided to the tree.
func NewTreeWithHashStrategy(cs []Content, hashStrategy func() hash.Hash) (*MerkleTree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewTreeWithHashStrategySorted just like NewTreeWithHashStrategy
// but sorts the siblings before hashing, mostly to follow the OpenZepplin Merkle implementation
// https://github.com/OpenZeppelin/openzeppelin-contracts-ethereum-package/blob/master/contracts/cryptography/MerkleProof.sol
func NewTreeWithHashStrategySorted(cs []Content, hashStrategy func() hash.Hash, sort bool) (*MerkleTree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetMerklePath: Get Merkle path and indexes(left leaf or right leaf)
func (m *MerkleTree) GetMerklePath(content Content) ([][]byte, []int64, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// right leaf

// left leaf

// buildWithContent is a helper function that for a given set of Contents, generates a
// corresponding tree and returns the root node, a list of leaf nodes, and a possible error.
// Returns an error if cs contains no Contents.
func buildWithContent(cs []Content, t *MerkleTree) (*Node, []*Node, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// buildIntermediate is a helper function that for a given list of leaf nodes, constructs
// the intermediate and root levels of the tree. Returns the resulting root node of the tree.
func buildIntermediate(nl []*Node, t *MerkleTree) (*Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MerkleRoot returns the unverified Merkle Root (hash of the root node) of the tree.
func (m *MerkleTree) MerkleRoot() []byte { _ = "STUB: not implemented"; return nil }

// RebuildTree is a helper function that will rebuild the tree reusing only the content that
// it holds in the leaves.
func (m *MerkleTree) RebuildTree() error { _ = "STUB: not implemented"; return nil }

// RebuildTreeWith replaces the content of the tree and does a complete rebuild; while the root of
// the tree will be replaced the MerkleTree completely survives this operation. Returns an error if the
// list of content cs contains no entries.
func (m *MerkleTree) RebuildTreeWith(cs []Content) error { _ = "STUB: not implemented"; return nil }

// VerifyTree verify tree validates the hashes at each level of the tree and returns true if the
// resulting hash at the root of the tree matches the resulting root hash; returns false otherwise.
func (m *MerkleTree) VerifyTree() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// VerifyContent indicates whether a given content is in the tree and the hashes are valid for that content.
// Returns true if the expected Merkle Root is equivalent to the Merkle root calculated on the critical path
// for a given content. Returns true if valid and false otherwise.
func (m *MerkleTree) VerifyContent(content Content) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// String returns a string representation of the node.
func (n *Node) String() string { _ = "STUB: not implemented"; return "" }

// String returns a string representation of the tree. Only leaf nodes are included
// in the output.
func (m *MerkleTree) String() string { _ = "STUB: not implemented"; return "" }
