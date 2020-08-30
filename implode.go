// https://www.jianshu.com/p/becb41ec5ef5
func Implode(list interface{}, seq string) string {
    listValue := reflect.Indirect(reflect.ValueOf(list))
    if listValue.Kind() != reflect.Slice {
        return ""
    }
    count := listValue.Len()
    listStr := make([]string, 0, count)
    for i := 0; i < count; i++ {
        v := listValue.Index(i)
        if str, err := getValue(v); err == nil {
            listStr = append(listStr, str)
        }
    }
    return strings.Join(listStr, seq)
}


func getValue(value reflect.Value) (res string, err error) {
    switch value.Kind() {
    case reflect.Ptr:
        res, err = GetValue(value.Elem())
    default:
        res = fmt.Sprint(value.Interface())
    }
    return
}
