package practice

import (
	"math/rand"
	"time"
)

type RandomizedCollection struct {
	//值->下标
	index map[int]map[int]bool
	data []int
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	this.data = append(this.data, val)
	if _, ok := this.index[val]; !ok {
		this.index[val] = map[int]bool{len(this.data) - 1:true}
		return true
	} else {
		this.index[val][len(this.data)-1] = true
		return false
	}
}


/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	if valMap, ok := this.index[val]; ok && len(valMap) != 0 {
		lastIdx := len(this.data) - 1
		lastVal := this.data[lastIdx]
		if val == lastVal {
			delete(this.index[lastVal], lastIdx)
			this.data = this.data[:len(this.data) - 1]
			return true
		}
		//获取val的一个索引
		var idx int
		for key := range valMap {
			idx = key
			break
		}
		//移动数据
		this.data[idx] = this.data[len(this.data) - 1]
		//处理索引
		delete(this.index[val], len(this.data) - 1)
		this.index[lastVal][idx] = true
		delete(this.index[lastVal], lastIdx)
		//删除最后一个数
		this.data = this.data[:len(this.data) - 1]
		return true
	}
	return false
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return this.data[rand.Intn(len(this.data))]
}
