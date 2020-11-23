package practice

func CanCompleteCircuit(gas []int, cost []int) int {
	l := len(gas)
	for idx, _ :=range gas {
		if gas[idx] < cost[idx] {
			continue
		}
		reGass := 0
		idy := idx
		for{
			reGass += gas[idy]
			reGass -= cost[idy]
			idy = (idy + 1) % l
			if idy == idx {
				break
			}
			if reGass < 0 {
				break
			}
		}
		if reGass < 0 {
			continue
		} else {
			return idx
		}
	}
	return -1
}
