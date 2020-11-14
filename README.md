# Goption
this module provide option/result enum  to go

result example
```golang
func MockOpenFile(filename string) (*os.File, error) {
	if filename == "not exists" {
		return nil, errors.New("not exists")

	}
	return &os.File{}, nil
}
func main(){
    defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
    }()
    file1 := goption.ToResult(MockOpenFile("a file")).Unwrap().(*os.File)
	fmt.Printf("%v", file1)
	goption.ToResult(MockOpenFile("not exists")).Unwrap()
}
```

option example
```golang
// conver nil to None
func WillreturnNil() *os.File {
	return nil
}
func main(){
   file2 := goption.ToOption(WillreturnNil())
	if file2.Is_None() {
		fmt.Println("wget none")
	}
}
```

see example to know more usage
