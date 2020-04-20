package slack

import (
	"testing"

	"github.com/0mili/mili"
	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func miliConf(t *testing.T) *mili.Config {
	miliConf := new(mili.Config)
	require.NoError(t, mili.WithLogger(zaptest.NewLogger(t)).Apply(miliConf))
	return miliConf
}

func TestDefaultConfig(t *testing.T) {
	conf, err := newConf("my-secret-token", miliConf(t), nil)
	require.NoError(t, err)
	assert.NotNil(t, conf.Logger)
	assert.Equal(t, "full", conf.SendMsgParams.Parse)
	assert.Equal(t, 1, conf.SendMsgParams.LinkNames)
}

func TestWithLogger(t *testing.T) {
	logger := zaptest.NewLogger(t)
	conf, err := newConf("my-secret-token", miliConf(t), []Option{
		WithLogger(logger),
	})

	require.NoError(t, err)
	assert.Equal(t, logger, conf.Logger)
}

func TestWithDebug(t *testing.T) {
	conf, err := newConf("my-secret-token", miliConf(t), []Option{
		WithDebug(true),
	})

	require.NoError(t, err)
	assert.Equal(t, true, conf.Debug)

	conf, err = newConf("my-secret-token", miliConf(t), []Option{
		WithDebug(false),
	})

	require.NoError(t, err)
	assert.Equal(t, false, conf.Debug)
}

func TestWithMessageParams(t *testing.T) {
	conf, err := newConf("my-secret-token", miliConf(t), []Option{
		WithMessageParams(slack.PostMessageParameters{
			Parse:     "none",
			LinkNames: 0,
		}),
	})

	require.NoError(t, err)
	assert.NotNil(t, conf.Logger)
	assert.Equal(t, "none", conf.SendMsgParams.Parse)
	assert.Equal(t, 0, conf.SendMsgParams.LinkNames)
}

func TestWithLogUnknownMessageTypes(t *testing.T) {
	conf, err := newConf("my-secret-token", miliConf(t), nil)
	require.NoError(t, err)

	assert.Equal(t, false, conf.LogUnknownMessageTypes, "LogUnknownMessageTypes should be disabled by default")

	conf, err = newConf("my-secret-token", miliConf(t), []Option{
		WithLogUnknownMessageTypes(),
	})
	require.NoError(t, err)
	assert.Equal(t, true, conf.LogUnknownMessageTypes)
}

func TestWithListenPassive(t *testing.T) {
	conf, err := newConf("my-secret-token", miliConf(t), []Option{
		WithListenPassive(),
	})

	require.NoError(t, err)
	assert.Equal(t, true, conf.ListenPassive)
}
