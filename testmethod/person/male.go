package person

// Male male struct
//go:generate  mockgen -source=../person/male.go -destination=../mock/male_mock.go -package=mock
type Male interface {
	Get(id int64) (MaleInfo, error)
}

// MaleInfo male info
type MaleInfo struct {
	Male int32 `json:"male"`
}
