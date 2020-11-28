package functionalConstants

import (
	"github.com/gofrs/uuid"
)

var (
	namespaceDNS  = uuid.Must(uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8"))
	namespaceURL  = uuid.Must(uuid.FromString("6ba7b811-9dad-11d1-80b4-00c04fd430c8"))
	namespaceOID  = uuid.Must(uuid.FromString("6ba7b812-9dad-11d1-80b4-00c04fd430c8"))
	namespaceX500 = uuid.Must(uuid.FromString("6ba7b814-9dad-11d1-80b4-00c04fd430c8"))
)

func GetnamespaceDNS() uuid.UUID {
	return namespaceDNS
}
func GetnamespaceURL() uuid.UUID {
	return namespaceURL
}
func GetnamespaceOID() uuid.UUID {
	return namespaceOID
}
func GetnamespaceX500() uuid.UUID {
	return namespaceX500
}
