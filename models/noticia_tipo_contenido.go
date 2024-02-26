package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/time_bogota"
)

type NoticiaTipoContenido struct {
	Id                int            `orm:"column(id);pk;auto"`
	Dato              string         `orm:"column(dato);type(json)"`
	Activo            bool           `orm:"column(activo)"`
	FechaCreacion     string         `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion string         `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
	IdNoticia         *Noticia       `orm:"column(id_noticia);rel(fk)"`
	IdTipoContenido   *TipoContenido `orm:"column(id_tipo_contenido);rel(fk)"`
}

func (t *NoticiaTipoContenido) TableName() string {
	return "noticia_tipo_contenido"
}

func init() {
	orm.RegisterModel(new(NoticiaTipoContenido))
}

// AddNoticiaTipoContenido insert a new NoticiaTipoContenido into database and returns
// last inserted Id on success.
func AddNoticiaTipoContenido(m *NoticiaTipoContenido) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetNoticiaTipoContenidoById retrieves NoticiaTipoContenido by Id. Returns error if
// Id doesn't exist
func GetNoticiaTipoContenidoById(id int) (v []*NoticiaTipoContenido, err error) {
	o := orm.NewOrm()
	if _, err := o.QueryTable("noticia_tipo_contenido").Filter("id_noticia", id).All(&v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetContenidosByIdNoticia retrieves NoticiaTipoContenido by IdNoticia. Returns error if
// IdNoticia doesn't exist
func GetContenidosByIdNoticia(id int) (v []*NoticiaTipoContenido, err error) {
	o := orm.NewOrm()
	if _, err := o.QueryTable("noticia_tipo_contenido").Filter("id_noticia", id).All(&v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllNoticiaTipoContenido retrieves all NoticiaTipoContenido matches certain condition. Returns empty list if
// no records exist
func GetAllNoticiaTipoContenido(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(NoticiaTipoContenido))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []NoticiaTipoContenido
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateNoticiaTipoContenido updates NoticiaTipoContenido by Id and returns error if
// the record to be updated doesn't exist
func UpdateNoticiaTipoContenidoById(m *NoticiaTipoContenido) (err error) {
	o := orm.NewOrm()
	v := NoticiaTipoContenido{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		m.FechaCreacion = time_bogota.TiempoCorreccionFormato(v.FechaCreacion)
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteNoticiaTipoContenido deletes NoticiaTipoContenido by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNoticiaTipoContenido(id int) (err error) {
	o := orm.NewOrm()
	v := NoticiaTipoContenido{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&NoticiaTipoContenido{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
