package repository_test

import (
	"events-api/internal/events/adapter/repository"
	"events-api/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSqlite_SQLiteEventsModelToApp(t *testing.T) {
	testSQLiteEvents := test.MakeSQLiteEvents()
	testAppEvents := test.MakeAppEvents(true)

	result := repository.SQLiteEventsModelToApp(testSQLiteEvents)

	assert.Equal(t, testAppEvents, result)
}
