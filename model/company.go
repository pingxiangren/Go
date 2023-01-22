package model

import (
	"encoding/base64"

	"github.com/solenovex/it/common"
)

// Company ...
//
//	type Company struct {
//		ID       string
//		Name     string
//		NickName string
//	}
//
// Company ...
type Device struct {
	ID        string
	AssetNo   string
	DevType   string
	DevStatus string
	// Picture  []byte
	Picture []byte
}
type Device2 struct {
	ID        string
	AssetNo   string
	DevType   string
	DevStatus string
	// Picture  []byte
	SPicture string
}

// GetAllDevices ...
func GetAllCompanies() (companies []Device2, err error) {
	sql := "SELECT id, title, belongs, devstate, devimage FROM deviceorder WHERE level = 2;"
	rows, err := common.Db.Query(sql)
	if err != nil {
		return
	}
	for rows.Next() {
		c := Device{}
		err = rows.Scan(&c.ID, &c.AssetNo, &c.DevType, &c.DevStatus, &c.Picture)
		if err != nil {
			return
		}
		// fmt.Printf("%v,%v,%v \n", c.AssetNo, c.DevType, c.DevStaus)
		// 图片blob转成base64字符串
		SPic := base64.StdEncoding.EncodeToString(c.Picture)
		c1 := Device2{}
		c1.ID = c.ID
		c1.AssetNo = c.AssetNo
		c1.DevType = c.DevType
		c1.DevStatus = c.DevStatus
		c1.SPicture = SPic
		companies = append(companies, c1)
	}
	return
}

// GetSearchDevices ... 暂时没做防sql注入
func GetSearchDevices(assetNo string, devType string, devStatus string) (devices []Device2, err error) {
	// sql := "SELECT id, title, belongs, devstate, devimage FROM deviceorder WHERE level = 2;"
	sql := `SELECT id, title, belongs, devstate, devimage FROM deviceorder WHERE level = 2 
	AND (title LIKE ?) AND (belongs LIKE ?) AND (devstate LIKE ?);`
	// fmt.Printf("%v,%v,%v \n", assetNo, devType, devStatus)
	rows, err := common.Db.Query(sql, "%"+assetNo+"%", "%"+devType+"%", "%"+devStatus+"%")
	if err != nil {
		return
	}
	for rows.Next() {
		d := Device{}
		err = rows.Scan(&d.ID, &d.AssetNo, &d.DevType, &d.DevStatus, &d.Picture)
		if err != nil {
			return
		}
		// 图片blob转成base64字符串
		SPic := base64.StdEncoding.EncodeToString(d.Picture)
		d1 := Device2{}
		d1.ID = d.ID
		d1.AssetNo = d.AssetNo
		d1.DevType = d.DevType
		d1.DevStatus = d.DevStatus
		d1.SPicture = SPic
		// fmt.Printf("%v,%v,%v \n", d.AssetNo, d.DevType, d.DevStatus)
		devices = append(devices, d1)
	}
	return
}

// GetCompany ..
func GetCompany(id string) (company Device, err error) {
	sql := "SELECT id, title, belongs, devstate, devimage FROM deviceorder WHERE id=?"
	err = common.Db.QueryRow(sql, id).Scan(&company.ID, &company.AssetNo, &company.DevType, &company.DevStatus, &company.Picture)
	return
}

// Insert ...
func (company *Device) Insert() (err error) {
	sql := "INSERT INTO deviceorder (title, level, belongs, devstate, devimage) VALUES (?, ?, ?, ?, ?)"
	stmt, err := common.Db.Prepare(sql)
	if err != nil {
		return
	}
	_, err = stmt.Exec(company.AssetNo, 2, company.DevType, company.DevStatus, company.Picture)
	if err != nil {
		return
	}
	return
}

// Update ...
func (company *Device) Update() (err error) {
	sql := "UPDATE deviceorder set title=?, belongs=? WHERE id=?"
	_, err = common.Db.Exec(sql, company.AssetNo, company.DevType, company.ID)
	return
}

// DeleteCompany ...
func DeleteCompany(id string) (err error) {
	sql := "DELETE FROM deviceorder WHERE id=?"
	_, err = common.Db.Exec(sql, id)
	return
}
