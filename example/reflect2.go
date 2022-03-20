package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Message struct {
	ID       uint64 `db:"id"`
	Channel  string `db:"channel"`
	UserName string `db:"user_name"`
}

func main() {
	arr := "333" //[]string{"111", "222", "3333"}
	isType := Any(arr)
	fmt.Println(isType)

	message := Message{ID: 1, Channel: "ChannelValue", UserName: "UserNameValue"}
	message_value := reflect.ValueOf(message)
	fmt.Println(message_value)

	messageNum := message_value.NumField()

	message_fields := make([]struct {
		Name  string
		Tag   string
		Value interface{}
	}, message_value.NumField())

	for i := 0; i < messageNum; i++ {
		field := message_value.Field(i)
		fieldType := message_value.Type().Field(i)
		message_fields[i].Name = fieldType.Name
		message_fields[i].Value = message_value.Interface()
		message_fields[i].Tag = fieldType.Tag.Get("db")
		fmt.Println("字段：", field)
		fmt.Println("值：", message_fields)
		fmt.Println("Tag：", message_fields[i].Tag)
	}
}

// Any formats any value as a string.
//Any将任何值格式化为字符串。
func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

// formatAtom formats a value without inspecting its internal structure.
//formatAtom格式化值，而不检查其内部结构。
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}
