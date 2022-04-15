package employee

import (
	"com.kq/dao"
	"com.kq/dto"
	"fmt"
)

func Add(employee dto.Employee) {
	fmt.Println(&employee)

	sqlStr := "INSERT INTO t_employee(username,name,age) VALUES(?,?,?)"
	ret, err := dao.Db.Exec(sqlStr, employee.Username, employee.Name, employee.Age)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // LastInsertId() 当插入新行时，返回新数据的id
	if err != nil {
		fmt.Printf("get LastInsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)

}

func Update(employee dto.Employee) {
	fmt.Println(&employee)

	sqlStr := "UPDATE t_employee SET username=?,name=?,age=? WHERE id=?"
	ret, err := dao.Db.Exec(sqlStr, employee.Username, employee.Name, employee.Age, employee.Id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	_, err = ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed,err:%v\n", err)
		return
	}

	fmt.Println(&employee)

}

func Get(id int) dto.Employee {

	sqlStr := "SELECT id,username,name,age FROM t_employee WHERE id=?"
	var e dto.Employee
	err := dao.Db.Get(&e, sqlStr, id)
	if err != nil {
		fmt.Printf("get failed,err:%v\n", err)
		return e
	}
	fmt.Printf("id:%d,name:%s,age:%d\n", e.Id, e.Name, e.Age)

	return e
}

func List(tenantid int) {

	sqlStr := "SELECT id,username,name,age FROM t_employee limit 10"
	var es []dto.Employee
	err := dao.Db.Select(&es, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed,err: %v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", es)

}

func Delete(id int) {

	sqlStr := "DELETE FROM t_employee WHERE id=?"
	ret, err := dao.Db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed,err:%v\n", err)
		return
	}
	fmt.Printf("delete success,affected rows is %v\n", n)

}
