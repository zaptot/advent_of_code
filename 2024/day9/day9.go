package main

import (
	"container/list"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const ENTITY_TYPE_FILE = 0
const ENTITY_TYPE_SPACE = 1

type Entity struct {
	id         int
	size       int
	entityType int
}

func main() {
	var part int
	flag.IntVar(&part, "part", 2, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1()
		fmt.Println("Output:", ans)
	} else {
		ans := part2()
		fmt.Println("Output:", ans)
	}
}

func part1() int {
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day9/input1")
	check(err)

	fileRows := strings.Split(string(fileData), "\n")
	filesystem := processFileSystem(strings.Join(fileRows, ""))
	res := 0
	l := 0
	r := len(filesystem) - 1
	idx := 0

	for l <= r {
		if filesystem[l].size < 1 {
			l += 1
			continue
		}
		if filesystem[r].entityType == ENTITY_TYPE_SPACE {
			r -= 1
			continue
		}

		if filesystem[r].size < 1 {
			r -= 1
			continue
		}

		if filesystem[l].entityType == ENTITY_TYPE_FILE {
			res += idx * filesystem[l].id
			filesystem[l].size -= 1
		} else {
			res += idx * filesystem[r].id
			filesystem[r].size -= 1
			filesystem[l].size -= 1
		}

		idx += 1
	}

	return res
}

// wrong: 6379071213627 6379047350239
// right: 6376648986651
func part2() int {
	fileData, err := os.ReadFile("/home/udz/advent_of_code/2024/day9/input1")
	check(err)

	fileRows := strings.Split(string(fileData), "\n")
	filesystem := processFileSystem(strings.Join(fileRows, ""))
	fileIdToElemtnt := make(map[int]*list.Element, len(filesystem)/2)
	filesList := list.New()
	res := 0
	idx := 0

	for i := 0; i < len(filesystem); i++ {
		fileEntity := filesystem[i]
		element := filesList.PushBack(&Entity{size: fileEntity.size, id: fileEntity.id, entityType: fileEntity.entityType})

		if fileEntity.entityType == ENTITY_TYPE_FILE {
			fileIdToElemtnt[fileEntity.id] = element
		}
	}

	for i := len(filesystem) - 1; i > 0; i-- {
		if filesystem[i].entityType == ENTITY_TYPE_SPACE {
			continue
		}

		file := filesystem[i]
		space, found := findLeftmostSpaceForFile(filesList, file)

		if found {
			fileEntity := fileIdToElemtnt[file.id]
			moveFileIntoSpace(filesList, fileEntity, space)
		}
	}

	for entityFromList := filesList.Front(); entityFromList != nil; entityFromList = entityFromList.Next() {
		entity := entityFromList.Value.(*Entity)

		for i := 0; i < entity.size; i++ {
			if entity.entityType == ENTITY_TYPE_FILE {
				res += idx * entity.id
			}

			idx += 1
		}
	}

	return res
}

func processFileSystem(fsString string) []Entity {
	fs := make([]Entity, len(fsString))

	for i := 0; i < len(fsString); i++ {
		size, err := strconv.Atoi(string(fsString[i]))
		check(err)

		id := i / 2
		entityType := ENTITY_TYPE_FILE
		if i%2 == 1 {
			entityType = ENTITY_TYPE_SPACE
		}
		fs[i] = Entity{size: size, id: id, entityType: entityType}
	}

	return fs
}

func findLeftmostSpaceForFile(list *list.List, file Entity) (*list.Element, bool) {
	for entityFromList := list.Front(); entityFromList != nil; entityFromList = entityFromList.Next() {
		entity := entityFromList.Value.(*Entity)
		if entity.id == file.id {
			break
		}

		if entity.entityType == ENTITY_TYPE_SPACE && entity.size >= file.size {
			return entityFromList, true
		}
	}

	return nil, false
}

func moveFileIntoSpace(list *list.List, file *list.Element, space *list.Element) {
	spaceEntity := space.Value.(*Entity)
	fileEntity := file.Value.(*Entity)
	fileParent := file.Prev()

	if fileEntity.size > spaceEntity.size {
		panic("cannot move file to space")
	}

	if fileEntity.entityType != ENTITY_TYPE_FILE {
		panic("Cannot move space")
	}

	list.InsertBefore(&Entity{size: fileEntity.size, id: fileEntity.id, entityType: ENTITY_TYPE_FILE}, space)
	list.Remove(file)
	list.InsertAfter(&Entity{size: fileEntity.size, id: 0, entityType: ENTITY_TYPE_SPACE}, fileParent)

	spaceEntity.size -= fileEntity.size
	if spaceEntity.size < 1 {
		list.Remove(space)
	}
}
