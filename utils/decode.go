package utils

import (
	"bufio"
	"fmt"
)

func DecodeVariableByteInteger(reader *bufio.Reader) (int, error) {
	multiplier := 1
	value := 0
	var currentByte byte
	var err error

	for i := 0; i < 4; i++ { 
		currentByte, err = reader.ReadByte()
		if err != nil {
			return 0, fmt.Errorf("failed to read byte for variable integer: %w", err)
		}
		value += int(currentByte & 127) * multiplier
		if (currentByte & 128) == 0 {
			return value, nil
		}
		multiplier *= 128
	}
	return 0, fmt.Errorf("malformed variable byte integer (exceeded 4 bytes)")
}
