package parsers

import (
	"fmt"
	"strconv"
	"strings"
)

// TempInfo - Simple CPU Temperature Information Strucutre
type TempInfo struct {
	PackageTemp float64
	CoreTemps   []float64
}

/**
 * Print all Temperature Information
 */
func (tInfo *TempInfo) Print() {
	fmt.Printf("Package Temp [%.2f]\n", tInfo.PackageTemp)
	for i := 0; i < len(tInfo.CoreTemps); i++ {
		fmt.Printf("Core%d [%.2f]\n", i, tInfo.CoreTemps[i])
	}
}

/**
 * Handle Error by Panicing
 */
func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

/**
 * Parses Output into Object
 */
func ParseOutput_intel(output []byte, tInfo *TempInfo) {
	strOut := string(output)

	// GET PACKAGE SECTION
	index := strings.Index(strOut, "Package")
	endIndex := index + (strings.Index(strOut[index:], "Core"))
	pkgSection := strOut[index:endIndex]

	// // PACKAGE TEMP
	index += strings.Index(pkgSection, "input")
	endIndex = index + (strings.Index(strOut[index:], "\n"))
	tempArr := strings.Split(strOut[index:endIndex], " ")[1]
	val, e := strconv.ParseFloat(tempArr, 64)
	handleError(e)
	tInfo.PackageTemp = val

	// // GET CORE TEMP SECTION
	foundAllCores := false
	var coreSection string
	tInfo.CoreTemps = make([]float64, 0, 10) // Allocate 10 Spots

	for !foundAllCores {
		// GET SECTION
		index = endIndex + strings.Index(strOut[endIndex:], "Core")
		endIndex = index + strings.Index(strOut[index+1:], "Core")

		// VERIFY FOUND
		if endIndex > index {
			coreSection = strOut[index:endIndex] // Still More
		} else {
			coreSection = strOut[index:] // Found Last Core
			foundAllCores = true
		}

		// PARSE TEMP
		i1 := strings.Index(coreSection, "input")
		i2 := i1 + strings.Index(coreSection[i1:], "\n")
		temp, err := strconv.ParseFloat(strings.Split(coreSection[i1:i2], " ")[1], 64)
		handleError(err)

		// STORE TEMP
		tInfo.CoreTemps = append(tInfo.CoreTemps, temp)
	}
}
