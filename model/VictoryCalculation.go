package model
import (
	"math"
	"fmt"
	"strconv"
)


func CalculatePossibleVictory(playerName string,k int,l int) (bool,*ResultQueue){
	field := GetPlayerField(playerName)

	res1,q1 := getVerticalResult(field,k,l)
	if(res1){
		return true,q1
	}
	res2,q2 := getLineResult(field,k,l)
	if(res2){
		return true,q2
	}
	res3,q3 := get1DiagResult(field,k,l)
	if(res3){
		return true,q3
	}
	res4,q4 := get2DiagResult(field,k,l)
	if(res4){
		return true,q4
	}
	return false,q1
}






func get1DiagResult(fld *Field,k int,l int)(bool,*ResultQueue){
	x0 := int(math.Max(0,math.Max(float64(k-l),float64(k-4))))
	x1 := int(math.Min(float64(k+5),math.Min(FIELD_SIZE,float64(k+FIELD_SIZE-l))))
	y0 := int(math.Max(0,math.Max(float64(l-k),float64(l-4))))
	y1 := int(math.Min(float64(l+5),math.Min(FIELD_SIZE,float64(l+FIELD_SIZE-k))))


	q := &ResultQueue{0, [5]int{}, [5]int{}}
	//fmt.Println(fld)
	fmt.Println("init="+strconv.Itoa(x0)+" "+strconv.Itoa(x1)+" "+strconv.Itoa(k)+" "+strconv.Itoa(l))
	result := fld.Field[k][l]
	for i,j := x0,y0; i<x1 && j<y1; i,j=i+1,j+1 {

		res,q := getResultBase(q, fld, i, j, result)
		fmt.Println(strconv.Itoa(i)+" "+strconv.Itoa(j)+" "+strconv.Itoa(fld.Field[i][j])+strconv.Itoa(result))
		if (res) {
			return res, q
		}
	}
	return false, q
}


func getVerticalResult(fld *Field,k int,l int)(bool,*ResultQueue){

	x0:= int(math.Max(0,float64(k-4)))
	x1 := int(math.Min(float64(k+5),FIELD_SIZE))
	q := &ResultQueue{0,[5]int{},[5]int{}}
	//fmt.Println("init="+strconv.Itoa(x0)+" "+strconv.Itoa(x1)+" "+strconv.Itoa(k)+" "+strconv.Itoa(l))
	result := fld.Field[k][l]
	for i:=x0;i<x1;i++{
		res,q := getResultBase(q,fld,i,l,result)
		if(res){
			return res,q
		}
	}
	return false,q
}



func getLineResult(fld *Field,k int,l int)(bool,*ResultQueue) {
	y0 := int(math.Max(0, float64(l-4)))
	y1 := int(math.Min(float64(l+5), FIELD_SIZE))


	q := &ResultQueue{0, [5]int{}, [5]int{}}
	//fmt.Println("init="+strconv.Itoa(x0)+" "+strconv.Itoa(x1)+" "+strconv.Itoa(y0)+" "+strconv.Itoa(y1))
	result := fld.Field[k][l]
	for j := y0; j<y1; j++ {
		res,q := getResultBase(q, fld, k, j, result)
		if (res) {
			return res, q
		}
	}
	return false, q
}

func get2DiagResult(fld *Field,k int,l int)(bool,*ResultQueue){
	x0 := int(math.Max(0,math.Max(float64(k+l+1-FIELD_SIZE),float64(k-4))))
	x1 := int(math.Min(float64(k+5),math.Min(FIELD_SIZE,float64(l+k+1))))
	y1 := int(math.Max(-1,math.Max(float64(l+k-FIELD_SIZE),float64(l-5))))
	y0 := int(math.Min(float64(l+4),math.Min(FIELD_SIZE-1,float64(l+k))))
	fmt.Println("init="+strconv.Itoa(x1)+" "+strconv.Itoa(y1)+" "+strconv.Itoa(k)+" "+strconv.Itoa(l))
	q := &ResultQueue{0, [5]int{}, [5]int{}}
	//fmt.Println("init="+strconv.Itoa(x0)+" "+strconv.Itoa(x1)+" "+strconv.Itoa(y0)+" "+strconv.Itoa(y1))
	result := fld.Field[k][l]
	for i,j := x0,y0; i<x1 && j>y1; i,j=i+1,j-1 {
		fmt.Println(strconv.Itoa(i)+" "+strconv.Itoa(j)+" "+strconv.Itoa(fld.Field[i][j])+strconv.Itoa(result),q.ToString())
		res,q := getResultBase(q, fld, i, j, result)
		if (res) {
			return res, q
		}
	}
	return false, q

}


func getResultBase(q *ResultQueue,fld *Field,i int,j int,result int) (bool,*ResultQueue){
	field := (*fld).Field

	fmt.Println(field[i])
	//fmt.Println(strconv.Itoa(i)+" "+strconv.Itoa(j)+" "+strconv.Itoa(field[i][j])+strconv.Itoa(result))
	fmt.Println(field[i][j],result,q)
	if (field[i][j]==result){
		q.push(i,j)
	}else{
		q.free()
	}
	if (q.count==5){
		return true,q
	}

	return false,q

}