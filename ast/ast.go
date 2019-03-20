package ast

type AST struct {
	FileData  []byte
	Positions *PositionStorage
	Nodes     *NodeStorage
	Edges     *EdgeStorage
	RootNode  NodeID
}
