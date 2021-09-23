package main

import "github.com/davecgh/go-spew/spew"


type Tx struct {
	Rank int
	Rset map[int]bool
	Wset map[int]bool
}


func (t * Tx) HasContention(set map[int]bool) bool {
	for k, _ := range set {
		_, ok := t.Wset[k]; if ok {
			return true
		}
	}
	return false
}

func main() {
	txs := []Tx{
		{0, map[int]bool{1:true}, make(map[int]bool)},
		{0, map[int]bool{2:true},make(map[int]bool)},
		{0, map[int]bool{3:true},make(map[int]bool)},
		{0, map[int]bool{4:true},make(map[int]bool)},
		{0, map[int]bool{5:true},make(map[int]bool)},
		{0, map[int]bool{6:true},make(map[int]bool)},
		{0, map[int]bool{7:true},make(map[int]bool)},
		{0, map[int]bool{8:true},map[int]bool{2:true}},
		{0, map[int]bool{9:true},map[int]bool{3:true}},
		{0, map[int]bool{10:true},map[int]bool{5:true}},
		{0, map[int]bool{11:true},map[int]bool{9:true, 4:true, 10:true}},
		{0, map[int]bool{12:true},map[int]bool{1:true}},
		{0, map[int]bool{13:true},map[int]bool{8:true}},
		{0, map[int]bool{14:true},map[int]bool{11:true}},
		{0, map[int]bool{15:true},map[int]bool{11:true}},
		{0, map[int]bool{16:true},map[int]bool{7:true}},
	}
	//Apply rang in order of occurrence in TX list
	for i, tx := range txs {
		for j := i; j < len(txs); j++ {
			if txs[j].HasContention(tx.Rset) { //If this transaction has contention on R or W set
				if txs[j].Rank <= tx.Rank { // If this transaction is not before given, move it further in rang list
					txs[j].Rank = tx.Rank + 1
				}
			}
		}
	}

	buckets := make(map[int][]Tx) //make buckets grouped by rang
	for _, tx := range txs{
		b := buckets[tx.Rank]
		buckets[tx.Rank] = append(b, tx)
	}

	spew.Dump(buckets)
}
