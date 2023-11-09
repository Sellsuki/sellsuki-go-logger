package log

import (
	"github.com/Sellsuki/sellsuki-go-logger/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestNewEvent(t *testing.T) {
	// Create a zap.Logger for testing purposes
	logger, _ := zap.NewDevelopment()

	// Create a sample config
	config := config.Config{
		// Initialize your configuration fields here.
	}

	// Define test data for the payload
	payload := EventPayload{
		Entity:      "{{your_bu}}.order",
		ReferenceID: "ODR_1234567890",
		Action:      EventActionCreate,
		Result:      EventResultSuccess,
		Data: map[string]any{
			"key1": "value1",
			"key2": "value2",
		},
	}

	// Define the test message
	msg := "Test event message"

	// Call the NewEvent function to create an Event object
	event := NewEvent(logger, config, msg, payload).(Event)

	// Use assert functions to make your assertions
	assert.Equal(t, logger, event.logger, "Logger should match")
	assert.Equal(t, config, event.config, "Config should match")
	assert.Equal(t, msg, event.Message, "Message should match")

	// You can add more assertions for other fields as needed, such as the payload.
	// Use withField and other methods to add fields to the Event object and check if they are correctly set.
}
