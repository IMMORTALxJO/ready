package auto

import (
	"hypercheck/probe/types"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"

	"gotest.tools/assert"
)

func getProbe() types.Probe {
	probe, _ := GenerateProbe()
	return probe
}

func getProbeResult(probe types.Probe) bool {
	result, _ := probe.Up(types.NewProbeInput("", "", "", ""))
	return result
}

func getProbeMsg(probe types.Probe) string {
	_, msg := probe.Up(types.NewProbeInput("", "", "", ""))
	return msg
}

func TestAutoProbeHTTP(t *testing.T) {
	os.Clearenv()
	os.Setenv("API_URL", "https://postman:password@postman-echo.com/basic-auth")
	assert.Assert(t, getProbeResult(getProbe()))
}

func TestAutoProbeRedis(t *testing.T) {
	os.Clearenv()
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "6379")
	assert.Assert(t, getProbeResult(getProbe()))
}

func TestAutoProbeDB(t *testing.T) {
	os.Clearenv()
	os.Setenv("DB_HOST", "postgres://user:password@localhost:5432/postgres?sslmode=disable")
	log.SetLevel(log.DebugLevel)
	assert.Assert(t, getProbeResult(getProbe()))
}
