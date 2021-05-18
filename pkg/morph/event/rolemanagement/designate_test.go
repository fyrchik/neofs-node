package rolemanagement

import (
	"testing"

	"github.com/nspcc-dev/neo-go/pkg/core/native/noderoles"
	"github.com/nspcc-dev/neo-go/pkg/vm/stackitem"
	"github.com/stretchr/testify/require"
)

func TestParseRoleUpdate(t *testing.T) {
	t.Run("wrong number of arguments", func(t *testing.T) {
		_, err := ParseDesignate([]stackitem.Item{})
		require.Error(t, err)
	})
	t.Run("invalid item type", func(t *testing.T) {
		args := []stackitem.Item{stackitem.NewMap(), stackitem.Make(123)}
		_, err := ParseDesignate(args)
		require.Error(t, err)
	})
	t.Run("good", func(t *testing.T) {
		args := []stackitem.Item{stackitem.Make(int(noderoles.NeoFSAlphabet)), stackitem.Make(123)}
		e, err := ParseDesignate(args)
		require.NoError(t, err)
		require.Equal(t, noderoles.NeoFSAlphabet, e.(Designate).Role)
	})
}
