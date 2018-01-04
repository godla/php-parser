package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type CastUnset struct {
	Cast
}

func NewCastUnset(expr node.Node) node.Node {
	return CastUnset{
		Cast{
			map[string]interface{}{},
			nil,
			expr,
		},
	}
}

func (n CastUnset) Attributes() map[string]interface{} {
	return n.attributes
}

func (n CastUnset) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n CastUnset) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n CastUnset) Position() *node.Position {
	return n.position
}

func (n CastUnset) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n CastUnset) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
