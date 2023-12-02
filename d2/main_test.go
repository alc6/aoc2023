package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseSession(t *testing.T) {
	session := " 4 red"

	gotSession, gotErr := parseSession(session)
	assert.NoError(t, gotErr)
	assert.Equal(t, ColorsSet{"red": 4}, gotSession)
}

func Test_parseSessions(t *testing.T) {
	sessions := " 4 red; 5 blue"

	gotSessions, gotErr := parseSessions(sessions)
	assert.NoError(t, gotErr)
	assert.Equal(t, []ColorsSet{
		{"red": 4},
		{"blue": 5},
	}, gotSessions)
}
