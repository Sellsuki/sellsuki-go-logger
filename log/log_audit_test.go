package log

import (
	"github.com/Sellsuki/sellsuki-go-logger/config"
	"github.com/Sellsuki/sellsuki-go-logger/level"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestNewAudit(t *testing.T) {
	// Create a zap.Logger for testing purposes
	logger, _ := zap.NewDevelopment()

	// Create a sample config
	config := config.Config{
		// Initialize your configuration fields here.
	}

	// Define test data
	msg := "Test audit message"
	payload := AuditPayload{
		// Initialize your payload data as needed.
	}

	// Call the NewAudit function to create a Log object
	log := NewAudit(logger, config, msg, payload).(Audit)

	// Assert that the Log object is correctly initialized with the provided values
	assert.Equal(t, msg, log.Message)
	assert.Equal(t, level.Info, log.Level)
	assert.Equal(t, false, log.Alert)
	assert.Equal(t, config, log.config)
	assert.Equal(t, logger, log.logger)
	assert.Equal(t, TypeAudit, log.Type)

	// You can add more assertions for other fields as needed, such as the payload.
	// Use withField and other methods to add fields to the Log object and check if they are correctly set.
}
