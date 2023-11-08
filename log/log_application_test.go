package log

import (
	"github.com/Sellsuki/sellsuki-go-logger/config"
	"github.com/Sellsuki/sellsuki-go-logger/level"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestNewDebug(t *testing.T) {
	// Create a zap.Logger for testing purposes
	logger, _ := zap.NewDevelopment()

	// Create a sample config
	c := config.Config{
		// Initialize your configuration fields here.
	}

	// Define the test message
	msg := "Test debug message"

	// Call the NewDebug function to create a Debug object
	debug := NewDebug(logger, c, msg).(Debug)

	// Use assert functions to make your assertions
	assert.Equal(t, logger, debug.logger, "Logger should match")
	assert.Equal(t, c, debug.config, "Config should match")
	assert.Equal(t, msg, debug.Message, "Message should match")
	assert.Equal(t, level.Debug, debug.Level, "Level should be Debug")

	// You can add more assertions for other fields as needed.
	// Use WithField and other methods to add fields to the Debug object and check if they are correctly set.
}

func TestNewInfo(t *testing.T) {
	// Create a zap.Logger for testing purposes
	logger, _ := zap.NewDevelopment()

	// Create a sample config
	c := config.Config{
		// Initialize your configuration fields here.
	}

	// Define the test message
	msg := "Test Info message"

	// Call the NewInfo function to create a Info object
	Info := NewInfo(logger, c, msg).(Info)

	// Use assert functions to make your assertions
	assert.Equal(t, logger, Info.logger, "Logger should match")
	assert.Equal(t, c, Info.config, "Config should match")
	assert.Equal(t, msg, Info.Message, "Message should match")
	assert.Equal(t, level.Info, Info.Level, "Level should be Info")

	// You can add more assertions for other fields as needed.
	// Use WithField and other methods to add fields to the Info object and check if they are correctly set.
}

func TestNewWarn(t *testing.T) {
	// Create a zap.Logger for testing purposes
	logger, _ := zap.NewDevelopment()

	// Create a sample config
	c := config.Config{
		// Initialize your configuration fields here.
	}

	// Define the test message
	msg := "Test Warn message"

	// Call the NewWarn function to create a Warn object
	Warn := NewWarn(logger, c, msg).(Warn)

	// Use assert functions to make your assertions
	assert.Equal(t, logger, Warn.logger, "Logger should match")
	assert.Equal(t, c, Warn.config, "Config should match")
	assert.Equal(t, msg, Warn.Message, "Message should match")
	assert.Equal(t, level.Warn, Warn.Level, "Level should be Warn")

	// You can add more assertions for other fields as needed.
	// Use WithField and other methods to add fields to the Warn object and check if they are correctly set.
}

func TestNewError(t *testing.T) {
	// Create a zap.Logger for testing purposes
	logger, _ := zap.NewDevelopment()

	// Create a sample config
	c := config.Config{
		// Initialize your configuration fields here.
	}

	// Define the test message
	msg := "Test Error message"

	// Call the NewError function to create a Error object
	Error := NewError(logger, c, msg).(Error)

	// Use assert functions to make your assertions
	assert.Equal(t, logger, Error.logger, "Logger should match")
	assert.Equal(t, c, Error.config, "Config should match")
	assert.Equal(t, msg, Error.Message, "Message should match")
	assert.Equal(t, level.Error, Error.Level, "Level should be Error")

	// You can add more assertions for other fields as needed.
	// Use WithField and other methods to add fields to the Error object and check if they are correctly set.
}

func TestNewPanic(t *testing.T) {
	// Create a zap.Logger for testing purposes
	logger, _ := zap.NewDevelopment()

	// Create a sample config
	c := config.Config{
		// Initialize your configuration fields here.
	}

	// Define the test message
	msg := "Test Panic message"

	// Call the NewPanic function to create a Panic object
	Panic := NewPanic(logger, c, msg).(Panic)

	// Use assert functions to make your assertions
	assert.Equal(t, logger, Panic.logger, "Logger should match")
	assert.Equal(t, c, Panic.config, "Config should match")
	assert.Equal(t, msg, Panic.Message, "Message should match")
	assert.Equal(t, level.Panic, Panic.Level, "Level should be Panic")

	// You can add more assertions for other fields as needed.
	// Use WithField and other methods to add fields to the Panic object and check if they are correctly set.
}

func TestNewFatal(t *testing.T) {
	// Create a zap.Logger for testing purposes
	logger, _ := zap.NewDevelopment()

	// Create a sample config
	c := config.Config{
		// Initialize your configuration fields here.
	}

	// Define the test message
	msg := "Test Fatal message"

	// Call the NewFatal function to create a Fatal object
	Fatal := NewFatal(logger, c, msg).(Fatal)

	// Use assert functions to make your assertions
	assert.Equal(t, logger, Fatal.logger, "Logger should match")
	assert.Equal(t, c, Fatal.config, "Config should match")
	assert.Equal(t, msg, Fatal.Message, "Message should match")
	assert.Equal(t, level.Fatal, Fatal.Level, "Level should be Fatal")

	// You can add more assertions for other fields as needed.
	// Use WithField and other methods to add fields to the Fatal object and check if they are correctly set.
}
