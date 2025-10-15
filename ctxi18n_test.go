package ctxi18n_test

import (
	"context"
	"testing"

	"github.com/josja/ctxi18n"
	"github.com/josja/ctxi18n/i18n"
	"github.com/josja/ctxi18n/internal/examples"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaults(t *testing.T) {
	assert.Equal(t, i18n.Code("en"), ctxi18n.DefaultLocale)
}

func TestLoad(t *testing.T) {
	err := ctxi18n.Load(examples.Content)
	assert.NoError(t, err)

	l := ctxi18n.Get("en")
	assert.NotNil(t, l)
	assert.Equal(t, "en", l.Code().String())
}

func TestLoadWithDefault(t *testing.T) {
	err := ctxi18n.LoadWithDefault(examples.Content, "en")
	assert.NoError(t, err)

	l := ctxi18n.Get("en")
	assert.NotNil(t, l)
	assert.Equal(t, "en", l.Code().String())
	assert.Equal(t, "Special Label", l.T("special_label"))
	l = ctxi18n.Get("es")
	assert.NotNil(t, l)
	assert.Equal(t, "es", l.Code().String())
	assert.Equal(t, "Special Label", l.T("special_label"))
}

func TestGet(t *testing.T) {
	err := ctxi18n.Load(examples.Content)
	assert.NoError(t, err)

	l := ctxi18n.Get("en")
	assert.NotNil(t, l)
	assert.Equal(t, "en", l.Code().String())

	l = ctxi18n.Get("bad")
	assert.Nil(t, l)
}

func TestMatch(t *testing.T) {
	err := ctxi18n.Load(examples.Content)
	require.NoError(t, err)

	l := ctxi18n.Match("en-US,en;q=0.9,es;q=0.8")
	assert.NotNil(t, l)
	assert.Equal(t, "en", l.Code().String())
}

func TestWithLocale(t *testing.T) {
	err := ctxi18n.Load(examples.Content)
	require.NoError(t, err)

	ctx := context.Background()
	ctx, err = ctxi18n.WithLocale(ctx, "en-US,en;q=0.9,es;q=0.8")
	require.NoError(t, err)

	l := ctxi18n.Locale(ctx)
	assert.NotNil(t, l)
	assert.Equal(t, "en", l.Code().String())

	// Use the default locale if not set
	ctx, err = ctxi18n.WithLocale(ctx, "inv")
	assert.NoError(t, err)
	l = ctxi18n.Locale(ctx)
	assert.NotNil(t, l)
	assert.Equal(t, "en", l.Code().String())

	ctxi18n.DefaultLocale = "bad"
	_, err = ctxi18n.WithLocale(ctx, "inv")
	assert.ErrorIs(t, err, ctxi18n.ErrMissingLocale)
	ctxi18n.DefaultLocale = "es"

}
