package guest

import (
	"regexp"

	"github.com/giantswarm/microerror"
)

const (
	// dnsNotReadyPattern is a regular expression representing DNS errors for
	// the guest API domain. Also see the following match example.
	//
	//     https://play.golang.org/p/ipBkwqlc4Td
	//
	dnsNotReadyPattern = "dial tcp: lookup .* on .*:53: no such host"

	// nodeEOFPattern is a regular expression representing EOF errors for the
	// guest API domain. Also see the following match example.
	//
	//     https://play.golang.org/p/L6f4ItJLufv
	//
	nodeEOFPattern = `Get https://api\..*/api/v1/nodes.* (unexpected )?EOF`

	// resourceEOFPattern is a regular expression representing EOF errors for the
	// guest API domain. Also see the following match example.
	//
	//     https://play.golang.org/p/2x2BXd5iHuP
	//
	resourceEOFPattern = `[Get|Post] https://api\..*/api/v1/namespaces/*/.* (unexpected )?EOF`

	// tlsHandshakeTimeoutPattern is a regular expression representing TLS
	// handshake timeout errors related to establishing connections to guest
	// clusters while the guest API is not fully up. Also see the following match
	// example.
	//
	//     https://play.golang.org/p/KuH2V0IXo-J
	//
	tlsHandshakeTimeoutPattern = `Get https://api\..*/api/v1/nodes.* net/http: TLS handshake timeout`

	// transientInvalidCertificatePattern is a regular expression representing the
	// kind of transient errors related to certificates returned while the guest
	// API is not fully up. Also see the following match example.
	//
	//     https://play.golang.org/p/iiYvBhPOg4f
	//
	transientInvalidCertificatePattern = `[Get|Post] https://api\..*: x509: certificate is valid for ingress.local, not api\..*`
)

var (
	dnsNotReadyRegexp                 = regexp.MustCompile(dnsNotReadyPattern)
	nodeEOFRegexp                     = regexp.MustCompile(nodeEOFPattern)
	resourceEOFRegexp                 = regexp.MustCompile(resourceEOFPattern)
	tlsHandshakeTimeoutRegexp         = regexp.MustCompile(tlsHandshakeTimeoutPattern)
	transientInvalidCertificateRegexp = regexp.MustCompile(transientInvalidCertificatePattern)
)

// APINotAvailableError is returned when the guest Kubernetes API is not
// available.
var APINotAvailableError = microerror.New("API not available")

// IsAPINotAvailable asserts APINotAvailableError.
func IsAPINotAvailable(err error) bool {
	if err == nil {
		return false
	}

	c := microerror.Cause(err)

	regexps := []*regexp.Regexp{
		dnsNotReadyRegexp,
		nodeEOFRegexp,
		resourceEOFRegexp,
		tlsHandshakeTimeoutRegexp,
		transientInvalidCertificateRegexp,
	}
	for _, re := range regexps {
		matched := re.MatchString(c.Error())

		if matched {
			return true
		}
	}

	if c == APINotAvailableError {
		return true
	}

	return false
}