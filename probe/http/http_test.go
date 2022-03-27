package http

import (
	"probe/probe"
	"testing"

	"gotest.tools/assert"
)

func TestHttp(t *testing.T) {
	assert.Assert(t, probe.GetProbeResult(GenerateProbe("https://postman-echo.com/status/200"), "code", "", "==", "200"))
	assert.Assert(t, probe.GetProbeResult(GenerateProbe("https://postman-echo.com/basic-auth"), "code", "", "==", "401"))
	assert.Assert(t, probe.GetProbeResult(GenerateProbe("https://postman-echo.com/status/200"), "online", "", "", ""))
	assert.Assert(t, probe.GetProbeResult(GenerateProbe("https://postman-echo.com/status/200"), "content", "", "==", "{\"status\":200}"))
	assert.Assert(t, probe.GetProbeResult(GenerateProbe("https://postman-echo.com/status/200"), "content", "length", "==", "14"))
	assert.Assert(t, probe.GetProbeResult(GenerateProbe("https://postman-echo.com/response-headers?foo1=bar1&foo2=bar2"), "headers", "any", "==", "Foo1: bar1"))
	assert.Assert(t, probe.GetProbeResult(GenerateProbe("https://postman-echo.com/response-headers?foo1=bar1&foo2=bar2"), "headers", "count", ">", "1"))

	assert.Assert(t, !probe.GetProbeResult(GenerateProbe("https://offline-not-found.com"), "online", "", "", ""))
}
