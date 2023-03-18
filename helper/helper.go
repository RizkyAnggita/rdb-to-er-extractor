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

func IsExistInForeignKeys(str string, arr []model.ForeignKey) bool {
	for _, a := range arr {
		if str == a.ColumnName {
			return true
		}
	}

	return false
}
