package memory_test

import (
	"github.com/think-free/ABCFitness-challenge/internal/database"
	"github.com/think-free/ABCFitness-challenge/internal/database/memory"
)

// Ensuring that the Memory type implements the Database interface
var _ database.Database = (*memory.Memory)(nil)
