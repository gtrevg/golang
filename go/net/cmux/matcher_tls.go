package cmux

import (
	"encoding/binary"
	"github.com/searKing/golang/go/crypto/tls"
)

// TLS matches HTTPS requests.
//
// By default, any TLS handshake packet is matched. An optional whitelist
// of versions can be passed in to restrict the matcher, for example:
//  TLS(tls.VersionTLS11, tls.VersionTLS12)
// reverse of crypto/tls/conn.go func (c *Conn) readRecordOrCCS(expectChangeCipherSpec bool) error {
// HandlerShake of TLS
// type byte	// recordTypeHandshake
// vers [2]byte
func TLS(versions ...int) MatcherFunc {
	const recordTypeHandshake = 22
	if len(versions) == 0 {
		versions = tls.Versions
	}
	var prefixes [][]byte
	for _, v := range versions {
		var ver = make([]byte, 2)
		binary.BigEndian.PutUint16(ver, uint16(v))
		var prefix []byte
		prefix = append(prefix, recordTypeHandshake)
		prefix = append(prefix, ver...)
		prefixes = append(prefixes, prefix)
	}
	return PrefixByteMatcher(prefixes...)
}
