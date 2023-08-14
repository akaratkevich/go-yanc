package actions

import (
	"fmt"
	"strings"
)

func CidrMaskToBinary(mask string) string {
	maskInt := 0
	fmt.Sscanf(mask, "%d", &maskInt) // Convert the string to an integer
	binaryMask := strings.Repeat("1", maskInt) + strings.Repeat("0", 32-maskInt)
	return fmt.Sprintf("%s.%s.%s.%s", binaryMask[:8], binaryMask[8:16], binaryMask[16:24], binaryMask[24:32])
}
