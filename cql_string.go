package nan

import "github.com/gocql/gocql"

// MarshalCQL - marshaller for cql
func (n NullString) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if !n.Valid {
		return nil, nil
	}

	return []byte(n.String), nil
}

// UnmarshalCQL - unmarshaller for cql
func (n *NullString) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	if data == nil {
		*n = NullString{}
		return nil
	}

	*n = NullString{Valid: true, String: string(data)}

	return nil
}
