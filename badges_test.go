package gotaseries

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBadgesService_Badges(t *testing.T) {
	data, err := os.ReadFile("data/badges/badge.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "GET", fmt.Sprintf("/%s", "badges/badge?id=106"), string(data))
	defer ts.Close()

	badge, err := bc.Badges.Badge(context.Background(), BadgesBadgeParams{
		ID: 106,
	})
	assert.NoError(t, err)

	assert.Equal(t, 106, badge.ID)
	assert.Equal(t, "Marathonien", badge.Name)
	assert.Equal(t, "Vous avez regardé 6 épisodes en plein dimanche, c'est un grand chelem !", badge.Description)
	assert.Equal(t, "https://www.betaseries.com/images/badges/marathonien.png", badge.Image)
	assert.Equal(t, 256, badge.Width)
	assert.Equal(t, 256, badge.Height)
}
