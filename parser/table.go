package parser

import "strings"

// ParseTable parses a string in format
//
// The first line is considered the Header
//
// Each line is separated by newlines
func ParseTable(table string) []map[string]string {
	final := make([]map[string]string, 0)

	// If this string doesn't contain any newline characters, return an empty []map[string]string
	if strings.Index(table, "\n") == -1 {
		return final
	}

	table = strings.ReplaceAll(table, "\t", "    ")

	words := make([]string, 0)
	wordIndexes := make([]int, 0)
	currentWord := ""
	spaceCount := 0
	for index, char := range strings.Split(table, "\n")[0] {
		// If the current character is a space & currentWord isn't already
		// empty, add a space to this spaceCount
		if char == ' ' && currentWord != "" {
			spaceCount++
		} else if char != ' ' {
			// If the word is empty and the current character isn't a space,
			// we've found the start of a new word
			if currentWord == "" {
				wordIndexes = append(wordIndexes, index)
			}

			// Add this character to our current word
			currentWord += string(char)
		}

		// If there has been two consecutive spaces, OR we're at the last character before newline
		if spaceCount == 2 || index == strings.Index(table, "\n")-1 {
			words = append(words, currentWord)
			currentWord = ""
			spaceCount = 0
		}
	}

	// Split the rows by newline, start after the first newline
	rows := strings.Split(table[strings.Index(table, "\n")+1:len(table)], "\n")

	for _, row := range rows {
		// Ensure the length is >= 1
		if len(row) <= 0 {
			continue
		}

		// Create a map to store the values
		values := make(map[string]string, 0)

		// Iterator through each word
		for i := 0; i < len(words); i++ {
			// Set the end to the last index of the row
			end := len(row)

			// If this is isn't the last word, set the index to the index of the next word
			if i+1 < len(wordIndexes) && wordIndexes[i+1] < end {
				end = wordIndexes[i+1]
			}

			// If the end header is ommited because it is null, just replace it with ""
			// eg:
			// H1      H2      H3
			// v1      test
			// v2      test2   notOmmited
			if wordIndexes[i] > end {
				values[words[i]] = ""
			} else {
				// Store the values[currentWord] = row[startIndex : endIndex]
				values[words[i]] = strings.TrimSpace(row[wordIndexes[i]:end])
			}

		}

		// Append it to the final vector
		final = append(final, values)
	}

	return final
}
