package cacert

import (
	"encoding/pem"
	"fmt"
	"testing"
)

func TestCertImport(t *testing.T) {
	rest := Bundle
	var blocks []*pem.Block
	var block *pem.Block
	for len(rest) > 0 {
		block, rest = pem.Decode(rest)
		if block != nil {
			blocks = append(blocks, block)
		} else {
			break
		}
	}
	fmt.Println(len(blocks), "certificates")
}
