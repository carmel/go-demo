package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"reflect"
	"time"
)

type Place struct {
	city    NullString `db:"city"`
	country NullString `db:"country"`
	telcode NullInt64  `db:"telcode"`
}
type CaResource struct {
	Remark     NullString `db:"remark"`      //
	Id         NullString `db:"id"`          //
	Name       NullString `db:"name"`        //
	Pid        NullString `db:"pid"`         //
	Sign       NullString `db:"sign"`        //
	Layer      NullString `db:"layer"`       //
	Path       NullString `db:"path"`        //
	Iconcls    NullString `db:"iconcls"`     //
	Creator    NullString `db:"creator"`     //
	Modifier   NullString `db:"modifier"`    //
	Version    NullString `db:"version"`     //
	CreateTime NullString `db:"create_time"` //
	Status     NullInt64  `db:"status"`      //
	ModifyTime NullString `db:"modify_time"` //
}

// CUSTOM NULL Handling structures

// NullInt64 is an alias for sql.NullInt64 data type
type NullInt64 sql.NullInt64

// Scan implements the Scanner interface for NullInt64
func (ni *NullInt64) Scan(value interface{}) error {
	var i sql.NullInt64
	if err := i.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ni = NullInt64{i.Int64, false}
	} else {
		*ni = NullInt64{i.Int64, true}
	}
	return nil
}

// NullBool is an alias for sql.NullBool data type
type NullBool sql.NullBool

// Scan implements the Scanner interface for NullBool
func (nb *NullBool) Scan(value interface{}) error {
	var b sql.NullBool
	if err := b.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nb = NullBool{b.Bool, false}
	} else {
		*nb = NullBool{b.Bool, true}
	}

	return nil
}

// NullFloat64 is an alias for sql.NullFloat64 data type
type NullFloat64 sql.NullFloat64

// Scan implements the Scanner interface for NullFloat64
func (nf *NullFloat64) Scan(value interface{}) error {
	var f sql.NullFloat64
	if err := f.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nf = NullFloat64{f.Float64, false}
	} else {
		*nf = NullFloat64{f.Float64, true}
	}

	return nil
}

// NullString is an alias for sql.NullString data type
type NullString sql.NullString

// Scan implements the Scanner interface for NullString
func (ns *NullString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ns = NullString{s.String, false}
	} else {
		*ns = NullString{s.String, true}
	}

	return nil
}

// NullTime is an alias for pq.NullTime data type
type NullTime pq.NullTime

// Scan implements the Scanner interface for NullTime
func (nt *NullTime) Scan(value interface{}) error {
	var t pq.NullTime
	if err := t.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nt = NullTime{t.Time, false}
	} else {
		*nt = NullTime{t.Time, true}
	}

	return nil
}

// MarshalJSON for NullInt64
func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

// UnmarshalJSON for NullInt64
func (ni *NullInt64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ni.Int64)
	ni.Valid = (err == nil)
	return err
}

// MarshalJSON for NullBool
func (nb *NullBool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nb.Bool)
}

// UnmarshalJSON for NullBool
func (nb *NullBool) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &nb.Bool)
	nb.Valid = (err == nil)
	return err
}

// MarshalJSON for NullFloat64
func (nf *NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nf.Float64)
}

// UnmarshalJSON for NullFloat64
func (nf *NullFloat64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &nf.Float64)
	nf.Valid = (err == nil)
	return err
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON for NullString
func (ns *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.String)
	ns.Valid = (err == nil)
	return err
}

// MarshalJSON for NullTime
func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	val := nt.Time.Format(time.RFC3339)
	return []byte(val), nil
}

// UnmarshalJSON for NullTime
func (nt *NullTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	// s = Stripchars(s, "\"")

	x, err := time.Parse(time.RFC3339, s)
	if err != nil {
		nt.Valid = false
		return err
	}

	nt.Time = x
	nt.Valid = true
	return nil
}

// MAIN program starts here
func main() {

	DB, err := sqlx.Connect("postgres", "postgres://postgres:tygspg2017@47.104.106.121:5432/sun_dev?sslmode=disable")
	defer DB.Close()
	if err != nil {
		fmt.Println(err)
	}
	//ts := []CaResource{}

	//	t := CaResource{}
	//	var ts []string
	//	rows, _ := DB.Queryx("SELECT remark,id,name,pid,sign,layer,path,iconcls,creator,modifier,version,create_time,status,modify_time FROM ca_resource")
	//	for rows.Next() {
	//		err = rows.StructScan(&t)
	//		v, _ := json.Marshal(&t)
	//		ts = append(ts, string(v))
	//	}
	var ts []interface{}
	rows, _ := DB.Queryx("SELECT remark,id,name,pid,sign,layer,path,iconcls,creator,modifier,version,create_time,status,modify_time FROM ca_resource")
	for rows.Next() {
		m := make(map[string]interface{})
		err = rows.MapScan(m)
		ts = append(ts, m)
	}
	sql := "SELECT remark,id,name,pid,sign,creator,modifier,version,create_time,status,modify_time FROM ca_resource limit $1"

	var m interface{}
	m = make(map[string]interface{})
	DB.QueryRowx(sql, 1).MapScan(m)
	fmt.Printf("m: ", m)

	//o := map[string]interface{}{"city": "hangzhou", "country": "zhongguo", "telcode": 198}
	//_, err = DB.NamedExec(`INSERT INTO place(city,telcode)VALUES(:city,:telcode)`, o)
	//fmt.Println(err)
	//p := Place{city: NullString{"杭州", true}}
	//fmt.Println(p)

}
