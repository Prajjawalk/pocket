package raintree

import (
	"fmt"
	"strings"

	cryptoPocket "github.com/pokt-network/pocket/shared/crypto"
)

type target struct {
	address    cryptoPocket.Address
	serviceUrl string

	level                  uint32  // the level of the node in the RainTree tree (inverse of height in traditional computer science)
	percentage             float64 // the target percentage within the peer list used to select this as a target
	addrBookLengthAtHeight int     // the length of the addr book at the specified block height and tree level
	index                  int     // the index of this target peer within the addr book at the specific height and level
	isSelf                 bool
}

func (t target) DebugString(n *rainTreeNetwork) string {
	s := strings.Builder{}
	s.WriteString("[")
	peersManagerStateView := n.peersManager.getNetworkView()
	for i := 0; i < t.addrBookLengthAtHeight; i++ {
		serviceUrl := peersManagerStateView.addrBookMap[peersManagerStateView.addrList[i]].ServiceUrl
		switch {
		case i == t.index && t.isSelf:
			fmt.Fprintf(&s, " (**%s**) ", serviceUrl)
		case i == t.index:
			fmt.Fprintf(&s, " **%s** ", serviceUrl)
		case t.isSelf:
			fmt.Fprintf(&s, " (%s) ", serviceUrl)
		default:
			fmt.Fprintf(&s, " %s ", serviceUrl)

		}
	}
	s.WriteString("]")
	return s.String()
}
