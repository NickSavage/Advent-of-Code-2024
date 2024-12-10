
func PartTwoMoveFilesContinguousSpace(input *Input) int {
	fileAlternate := true
	fileID := -1
	diskString := []string{}
	for _, char := range input.DiskMap {
		number, _ := strconv.Atoi(char)
		if fileAlternate {
			fileID += 1

			for range number {
				diskString = append(diskString, strconv.Itoa(fileID))
			}
			fileAlternate = false
		} else {
			for range number {
				diskString = append(diskString, ".")
			}
			fileAlternate = true
		}

	}

	log.Printf("%v", diskString)
	// identify last file
	index := len(diskString) - 1
	file := diskString[index]
	size := 1
	for {
		index -= 1

		if index == 0 {
			break
		}
		if diskString[index] == file {
			size += 1
			continue
		}
		if diskString[index] != file {
			if file != "." {
				foundSize := 0

				for i := range len(diskString) - 1 {
					if diskString[i] == "." {
						foundSize += 1
					} else {
						foundSize = 0
					}
					if i == index {
						break
					}
					if foundSize == size {
						moveIndex := i + 1 - size
						log.Printf("move file %v to moveIndex %v size %v", file, moveIndex, size)
						for n := range size {
							log.Printf("move %v to %v", diskString[index+n+1], moveIndex+n)
							diskString[moveIndex+n] = diskString[index+1+n]
							diskString[index+n+1] = "."
						}
						foundSize = 0
						break
					}
				}

			}

			//reset
			file = diskString[index]
			size = 1
		}

	}
	log.Printf("after %v", diskString)

	result := 0
	for i, char := range diskString {
		operand, _ := strconv.Atoi(char)
		result += i * operand
	}
	return result
}
