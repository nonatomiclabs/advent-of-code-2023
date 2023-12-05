package day_05

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type MappingRange struct {
	SourceRangeStart      int
	DestinationRangeStart int
	RangeLength           int
}

type Mapper struct {
	mappingRanges []MappingRange
}

func (m *Mapper) Map(input int) int {
	for _, mappingRange := range m.mappingRanges {
		if input >= mappingRange.SourceRangeStart && input < mappingRange.SourceRangeStart+mappingRange.RangeLength {
			offset := input - mappingRange.SourceRangeStart
			return mappingRange.DestinationRangeStart + offset
		}
	}
	return input
}

func NewMapper(mappingRanges []MappingRange) *Mapper {
	return &Mapper{
		mappingRanges: mappingRanges,
	}
}

type SeedRange struct {
	rangeStart  int
	rangeLength int
}

func MapSeedToLocation(seed int, mappers []Mapper) int {
	// fmt.Printf("Mapping seed value %v\n", seed)

	sourceValue := seed
	var destinationValue int
	for _, mapper := range mappers {
		// fmt.Printf("\tgoing through mapper %d: ", mapperIndex+1)
		destinationValue = mapper.Map(sourceValue)
		// fmt.Printf("%d mapped to %d\n", sourceValue, destinationValue)
		sourceValue = destinationValue
	}
	return destinationValue
}

func Solution(inputLines []string, part2 bool) int {
	seeds := []int{}
	mappers := []Mapper{}
	mappingRanges := []MappingRange{}
	var parseRange bool

	seedsRegex := regexp.MustCompile(`\d+`)
	rangeRegex := regexp.MustCompile(`\d+`)

	for line_index, line := range inputLines {
		if strings.HasPrefix(line, "seeds:") {
			seedsRaw := strings.Split(line, "seeds: ")[1]
			seedsMatch := seedsRegex.FindAllString(seedsRaw, -1)
			for _, seedString := range seedsMatch {
				seed, _ := strconv.Atoi(seedString)
				seeds = append(seeds, seed)
			}
		}

		if strings.HasSuffix(line, "map:") {
			fmt.Printf("Found start of map definition at line %d: %s\n", line_index+1, line)
			parseRange = true
			mappingRanges = []MappingRange{}
			continue
		}

		if (line == "" || line_index == len(inputLines)-1) && parseRange {
			fmt.Printf("Found end of map definition at line %d\n", line_index+1)
			parseRange = false
			mapper := NewMapper(mappingRanges)
			mappers = append(mappers, *mapper)
			continue
		}

		if parseRange {
			fmt.Printf("\tParsing range at line %d\n", line_index+1)
			rangeMatch := rangeRegex.FindAllString(line, -1)
			mappingRange := MappingRange{}
			for index, stringValue := range rangeMatch {
				value, _ := strconv.Atoi(stringValue)

				switch index {
				case 0:
					mappingRange.DestinationRangeStart = value
				case 1:
					mappingRange.SourceRangeStart = value
				case 2:
					mappingRange.RangeLength = value
				}
			}
			mappingRanges = append(mappingRanges, mappingRange)
		}

	}
	fmt.Printf("Found %d mappers\n", len(mappers))

	lowestLocationValue := MapSeedToLocation(seeds[0], mappers)
	if part2 {
		seedInput := seeds
		seedRanges := []SeedRange{}
		totalSeedsCount := 0
		for i := 0; i < len(seedInput)/2; i++ {
			rangeStart := seedInput[i*2]
			rangeLength := seedInput[i*2+1]
			seedRange := SeedRange{rangeStart: rangeStart, rangeLength: rangeLength}
			seedRanges = append(seedRanges, seedRange)
			totalSeedsCount += rangeLength
		}

		fmt.Printf("Will process %d seeds\n", totalSeedsCount)

		for _, seedRange := range seedRanges {
			for i := 0; i < seedRange.rangeLength; i++ {
				// fmt.Printf("Processing seed %d/%d\n", seedRangeIndex+i, totalSeedsCount)
				locationValue := MapSeedToLocation(seedRange.rangeStart+i, mappers)
				if locationValue < lowestLocationValue {
					lowestLocationValue = locationValue
				}
			}
		}
	} else {
		locationValues := []int{}
		for _, seed := range seeds[1:] {
			locationValue := MapSeedToLocation(seed, mappers)
			locationValues = append(locationValues, locationValue)
		}
		fmt.Printf("Location values: %v\n", locationValues)

		for _, locationValue := range locationValues {
			if locationValue < lowestLocationValue {
				lowestLocationValue = locationValue
			}
		}
	}

	fmt.Printf("Lowest location value: %d\n", lowestLocationValue)
	return lowestLocationValue
}
