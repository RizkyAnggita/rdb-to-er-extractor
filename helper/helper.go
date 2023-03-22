package helper

import "rdb-to-er-extractor/model"

func GenerateProperSubsetPK(arr []model.PrimaryKey) (res [][]model.PrimaryKey) {
	size := len(arr)               //denotes length of input string
	numOfSubset := 1 << uint(size) //equal to (2^n)

	for i := 0; i < numOfSubset; i++ {
		subset := []model.PrimaryKey{}
		for j := 0; j < size; j++ {
			if i&(1<<uint(j)) != 0 {
				subset = append(subset, arr[j])
			}
		}
		if len(subset) != 0 && len(subset) != size {
			res = append(res, subset)
		}
	}
	return res
}

func IsExistInPrimaryKeys(str string, arr []model.PrimaryKey) bool {
	for _, a := range arr {
		if str == a.ColumnName {
			return true
		}
	}

	return false
}

func GetTableByTableName(name string, tables []model.Table) model.Table {
	for _, t := range tables {
		if t.Name == name {
			return t
		}
	}
	return model.Table{}
}

func IsExistInForeignKeys(str string, arr []model.ForeignKey) bool {
	for _, a := range arr {
		if str == a.ColumnName {
			return true
		}
	}

	return false
}

func IsSubset(setA, setB []string) bool {
	checkSet := map[string]bool{}

	for _, val := range setA {
		checkSet[val] = true
	}

	for _, val := range setB {
		if checkSet[val] {
			delete(checkSet, val)
		}
	}

	return len(checkSet) == 0
}

// SetDifference returns the elements in `setA` that aren't in `setB`.
func SetDifference(setA []string, setB []string) []string {
	mapDiff := map[string]bool{}
	diffStr := []string{}

	for _, s := range setB {
		mapDiff[s] = true
	}

	for _, s := range setA {
		if _, isExist := mapDiff[s]; !isExist {
			diffStr = append(diffStr, s)
		}
	}

	return diffStr
}
