package nest

/* TODO
import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/blaskovicz/go-underarmour/mocks"
	"github.com/blaskovicz/go-underarmour/models"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	s := mocks.NewUnderArmourAPI()
	defer s.Close()
	os.Setenv("UNDERARMOUR_ROOT_URI", s.URL())
	os.Setenv("UNDERARMOUR_COOKIE_AUTH_TOKEN", "some_token.123")
	var err error
	var client *Client
	var user *models.User
	var route *models.Route
	t.Run("Init", func(t *testing.T) {
		client = New()
		require.NotNil(t, client, "client was nil")
	})
	t.Run("ReadUser", func(t *testing.T) {
		require.NotNil(t, client)
		user, err = client.ReadUser("self")
		require.NoError(t, err, "read user failed")
		require.NotNil(t, user, "user was nil")
		require.Equal(t, "Zach", user.FirstName)
		require.Equal(t, "Person", user.LastName)
		require.Equal(t, "Zach123", user.Username)
		require.Equal(t, "Zach Person", user.DisplayName)
		require.Equal(t, "P.", user.LastInitial)
		require.Equal(t, "M", user.Gender)
		require.Equal(t, "en-US", user.PreferredLanguage)
		require.Equal(t, "New York City", user.Location.Locality)
		require.Equal(t, "NY", user.Location.Region)
		require.Equal(t, "US", user.Location.Country)
		require.Equal(t, "running", user.Hobbies)
		require.Equal(t, "sup dog", user.Introduction)
		require.Equal(t, "America/New_York", user.TimeZone)
		require.Equal(t, "", user.GoalStatement)
		require.Equal(t, "", user.ProfileStatement)
		require.Equal(t, 2017, user.DateJoined.Year())
		require.Equal(t, time.Month(7), user.DateJoined.Month())
		require.Equal(t, 7, user.DateJoined.Day())
		require.Equal(t, 117774799, user.ID)
	})
	t.Run("ReadRoute", func(t *testing.T) {
		require.NotNil(t, client)

		// std format
		route, err = client.ReadRoute(1784229029)
		require.NoError(t, err, "read route failed")
		require.NotNil(t, route)
		require.Equal(t, "RUNNING RUNNERS - 9", route.Name)
		require.Equal(t, -57.9541283985, route.TotalDescent)
		require.Equal(t, 54.1198459896, route.TotalAscent)
		require.Equal(t, "1784229029", route.Links["self"][0]["id"])

		// gpx format
		gpxRoute, err := client.ReadRouteGPX(1784229029)
		require.NoError(t, err, "read gpx route failed")
		require.NotNil(t, gpxRoute)
		require.Equal(t, 1, len(gpxRoute.Tracks))
		// TODO leading and trailing whitespace for name... bad xmlns or bug?
		require.Equal(t, "RUNNING RUNNERS - 9", strings.TrimSpace(gpxRoute.Tracks[0].Name))
		require.Equal(t, 10, len(gpxRoute.Tracks[0].Segments[0].Waypoints))
		require.Equal(t, 41.10955, gpxRoute.Tracks[0].Segments[0].Waypoints[0].Lat)
		require.Equal(t, -73.418, gpxRoute.Tracks[0].Segments[0].Waypoints[0].Lon)
	})
}
*/
