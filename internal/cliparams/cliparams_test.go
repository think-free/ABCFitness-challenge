package cliparams_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/think-free/ABCFitness-challenge/internal/cliparams"
)

func TestXxx(t *testing.T) {
	os.Setenv("DBTYPE", "test_dbtype")
	os.Setenv("LOGLEVEL", "test_loglevel")

	cp := cliparams.New()
	require.Equal(t, "test_dbtype", cp.DatabaseType)
	require.Equal(t, "test_loglevel", cp.LogLevel)
}
