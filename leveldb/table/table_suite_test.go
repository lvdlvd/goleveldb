package table

import (
	"testing"

	"github.com/lvdlvd/goleveldb/leveldb/testutil"
)

func TestTable(t *testing.T) {
	testutil.RunSuite(t, "Table Suite")
}
