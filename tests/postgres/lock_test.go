package postgres

import (
	"context"
	"github.com/goranbjelanovic/jet/v2/internal/testutils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	. "github.com/goranbjelanovic/jet/v2/postgres"
	. "github.com/goranbjelanovic/jet/v2/tests/.gentestdata/jetdb/dvds/table"
)

func TestLockTable(t *testing.T) {
	expectedSQL := `
LOCK TABLE dvds.address IN`

	var testData = []TableLockMode{
		LOCK_ACCESS_SHARE,
		LOCK_ROW_SHARE,
		LOCK_ROW_EXCLUSIVE,
		LOCK_SHARE_UPDATE_EXCLUSIVE,
		LOCK_SHARE,
		LOCK_SHARE_ROW_EXCLUSIVE,
		LOCK_EXCLUSIVE,
		LOCK_ACCESS_EXCLUSIVE,
	}

	for _, lockMode := range testData {
		query := Address.LOCK().IN(lockMode)

		testutils.AssertDebugStatementSql(t, query, expectedSQL+" "+string(lockMode)+" MODE;\n")

		tx, _ := db.Begin()

		_, err := query.Exec(tx)

		require.NoError(t, err)

		err = tx.Rollback()

		require.NoError(t, err)
		requireLogged(t, query)
	}

	for _, lockMode := range testData {
		query := Address.LOCK().IN(lockMode).NOWAIT()

		testutils.AssertDebugStatementSql(t, query, expectedSQL+" "+string(lockMode)+" MODE NOWAIT;\n")

		tx, _ := db.Begin()

		_, err := query.Exec(tx)

		require.NoError(t, err)

		err = tx.Rollback()

		require.NoError(t, err)
		requireLogged(t, query)
	}
}

func TestLockExecContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)
	defer cancel()

	time.Sleep(10 * time.Millisecond)

	tx, _ := db.Begin()
	defer tx.Rollback()

	_, err := Address.LOCK().IN(LOCK_ACCESS_SHARE).ExecContext(ctx, tx)

	require.Error(t, err, "context deadline exceeded")
}
