package model
import "strconv"

type ResultQueue struct{

	count int
	x [5]int
	y [5]int
}

func (q *ResultQueue) push(i int,j int){
	q.x[q.count] = i
	q.y[q.count] = j
	q.count++
}

func (q *ResultQueue) free(){
	q.x = [5]int{}
	q.y = [5]int{}
	q.count = 0
}

func (q *ResultQueue) ToString() string{
	result:=""
	for i:=0;i<5;i++{
		if(i>0){
			result+=";"
		}
		result+=strconv.Itoa(q.x[i])+","+strconv.Itoa(q.y[i])
	}
	return result
}