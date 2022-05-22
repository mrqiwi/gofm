package utils

type MemoryUnit string

const (
	Byte   MemoryUnit = "B"
	KiByte MemoryUnit = "KiB"
	MiByte MemoryUnit = "MiB"
	GiByte MemoryUnit = "GiB"
	TiByte MemoryUnit = "TiB"
	PiByte MemoryUnit = "PiB"
	EiByte MemoryUnit = "EiB"
)

func ConvertBytes(size int64) (float64, MemoryUnit) {
	const (
		kib = 1024
		mib = 1024 * 1024
		gib = 1024 * 1024 * 1024
		tib = 1024 * 1024 * 1024 * 1024
		pib = 1024 * 1024 * 1024 * 1024 * 1024
		eib = 1024 * 1024 * 1024 * 1024 * 1024 * 1024
	)

	calcResult := func(size, b int64) float64 {
		value := float64(size) / float64(b)
		return RoundFloat(value, 1)
	}

	if size >= eib {
		return calcResult(size, eib), EiByte
	} else if size >= pib {
		return calcResult(size, pib), PiByte
	} else if size >= tib {
		return calcResult(size, tib), TiByte
	} else if size >= gib {
		return calcResult(size, gib), GiByte
	} else if size >= mib {
		return calcResult(size, mib), MiByte
	} else if size >= kib {
		return calcResult(size, kib), KiByte
	}

	return float64(size), Byte
}
