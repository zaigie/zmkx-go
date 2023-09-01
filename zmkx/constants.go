package zmkx

const (
	ZmkxVID   uint16 = 0x1d50
	ZmkxPID   uint16 = 0x615e
	ZmkxUsage uint16 = 0xff14

	ReportCount int = 63
	PayloadSize int = ReportCount - 1

	EinkWidth  int = 128
	EinkHeight int = 296
)
