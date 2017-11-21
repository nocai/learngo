package filter

import "testing"

func TestFilter(t *testing.T) {
	minSize, maxSize, suffixes, files := handleCommandLine()
	//sink(filterSize(minSize, maxSize, filterSuffixes(suffixes, source(files))))
	channel1 := source(files)
	channel2 := filterSuffixes(suffixes, channel1)
	channel3 := filterSize(minSize, maxSize, channel2)
	sink(channel3)
}

func handleCommandLine() (int, int, []string, []string) {
	return 0, 0, []string{}, []string{}
}
