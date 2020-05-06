package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type Person struct {
	name string
	age int
}

func main()  {
	http.HandleFunc("/", search)

	http.ListenAndServe(":8080",nil)
}

func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}
	data.MaxResults = 10 // set default
	if err := Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}

	// ...rest of handler...
	fmt.Fprintf(resp, "Search: %+v\n", data)
}

func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	typeMatch := make(map[string]string)
	ty := reflect.TypeOf(ptr).Elem()
	for j :=0; j<ty.NumField();j++  {
		fi := ty.Field(j).Name
		tag := ty.Field(j).Tag
		tagName := tag.Get("http")
		typeMatch[fi] = tagName
	}

	for key,_:=range req.Form{
		_,ok:=typeMatch[key]
		if !ok{
			fmt.Println("参数不存在")
		}
		//t := reflect.ValueOf(ptr).Elem()


	}


	return nil

	for value,k := range typeMatch{
		fmt.Println(value,k)
	}



	// Build map of fields keyed by effective name.
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			fmt.Println(f,f.Kind())
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}

//通过发射来修改值
func modifyValue()  {
	x := 2
	d := reflect.ValueOf(&x).Elem()   // d refers to the variable x
	px := d.Addr().Interface().(*int) // px := &x
	*px = 3                           // x = 3
	fmt.Println(x)
}

func test()  {
	var num int64 = -100;
	r := reflect.ValueOf(num);
	fmt.Println("type",reflect.TypeOf(num))
	fmt.Println("value",reflect.ValueOf(num))
	fmt.Println(r.Kind())
	fmt.Println(reflect.PtrTo(reflect.TypeOf(num)))
	fmt.Println(num)
}


