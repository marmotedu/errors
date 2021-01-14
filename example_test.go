package errors

import (
	"fmt"
)

func ExampleNew() {
	err := New("whoops")
	fmt.Println(err)

	// Output: whoops
}

func ExampleNew_printf() {
	err := New("whoops")
	fmt.Printf("%+v", err)

	// Example output:
	// whoops
	// github.com/marmotedu/errors_test.ExampleNew_printf
	//         /home/dfc/src/github.com/marmotedu/errors/example_test.go:17
	// testing.runExample
	//         /home/dfc/go/src/testing/example.go:114
	// testing.RunExamples
	//         /home/dfc/go/src/testing/example.go:38
	// testing.(*M).Run
	//         /home/dfc/go/src/testing/testing.go:744
	// main.main
	//         /github.com/marmotedu/errors/_test/_testmain.go:106
	// runtime.main
	//         /home/dfc/go/src/runtime/proc.go:183
	// runtime.goexit
	//         /home/dfc/go/src/runtime/asm_amd64.s:2059
}

func ExampleWithMessage() {
	cause := New("whoops")
	err := WithMessage(cause, "oh noes")
	fmt.Println(err)

	// Output: oh noes
}

func ExampleWithStack() {
	cause := New("whoops")
	err := WithStack(cause)
	fmt.Println(err)

	// Output: whoops
}

func ExampleWithStack_printf() {
	cause := New("whoops")
	err := WithStack(cause)
	fmt.Printf("%+v", err)

	// Example Output:
	// whoops
	// github.com/marmotedu/errors_test.ExampleWithStack_printf
	//         /home/fabstu/go/src/github.com/marmotedu/errors/example_test.go:55
	// testing.runExample
	//         /usr/lib/go/src/testing/example.go:114
	// testing.RunExamples
	//         /usr/lib/go/src/testing/example.go:38
	// testing.(*M).Run
	//         /usr/lib/go/src/testing/testing.go:744
	// main.main
	//         github.com/marmotedu/errors/_test/_testmain.go:106
	// runtime.main
	//         /usr/lib/go/src/runtime/proc.go:183
	// runtime.goexit
	//         /usr/lib/go/src/runtime/asm_amd64.s:2086
	// github.com/marmotedu/errors_test.ExampleWithStack_printf
	//         /home/fabstu/go/src/github.com/marmotedu/errors/example_test.go:56
	// testing.runExample
	//         /usr/lib/go/src/testing/example.go:114
	// testing.RunExamples
	//         /usr/lib/go/src/testing/example.go:38
	// testing.(*M).Run
	//         /usr/lib/go/src/testing/testing.go:744
	// main.main
	//         github.com/marmotedu/errors/_test/_testmain.go:106
	// runtime.main
	//         /usr/lib/go/src/runtime/proc.go:183
	// runtime.goexit
	//         /usr/lib/go/src/runtime/asm_amd64.s:2086
}

func ExampleWrap() {
	cause := New("whoops")
	err := Wrap(cause, "oh noes")
	fmt.Println(err)

	// Output: oh noes
}

func fn() error {
	e1 := New("error")
	e2 := Wrap(e1, "inner")
	e3 := Wrap(e2, "middle")
	return Wrap(e3, "outer")
}

func ExampleCause() {
	err := fn()
	fmt.Println(err)
	fmt.Println(Cause(err))

	// Output: outer
	// error
}

func ExampleWrap_extended() {
	err := fn()
	fmt.Printf("%+v\n", err)

	// Example output:
	// error
	// github.com/marmotedu/errors_test.fn
	//         /home/dfc/src/github.com/marmotedu/errors/example_test.go:47
	// github.com/marmotedu/errors_test.ExampleCause_printf
	//         /home/dfc/src/github.com/marmotedu/errors/example_test.go:63
	// testing.runExample
	//         /home/dfc/go/src/testing/example.go:114
	// testing.RunExamples
	//         /home/dfc/go/src/testing/example.go:38
	// testing.(*M).Run
	//         /home/dfc/go/src/testing/testing.go:744
	// main.main
	//         /github.com/marmotedu/errors/_test/_testmain.go:104
	// runtime.main
	//         /home/dfc/go/src/runtime/proc.go:183
	// runtime.goexit
	//         /home/dfc/go/src/runtime/asm_amd64.s:2059
	// github.com/marmotedu/errors_test.fn
	// 	  /home/dfc/src/github.com/marmotedu/errors/example_test.go:48: inner
	// github.com/marmotedu/errors_test.fn
	//        /home/dfc/src/github.com/marmotedu/errors/example_test.go:49: middle
	// github.com/marmotedu/errors_test.fn
	//      /home/dfc/src/github.com/marmotedu/errors/example_test.go:50: outer
}

func ExampleWrapf() {
	cause := New("whoops")
	err := Wrapf(cause, "oh noes #%d", 2)
	fmt.Println(err)

	// Output: oh noes #2
}

func ExampleErrorf_extended() {
	err := Errorf("whoops: %s", "foo")
	fmt.Printf("%+v", err)

	// Example output:
	// whoops: foo
	// github.com/marmotedu/errors_test.ExampleErrorf
	//         /home/dfc/src/github.com/marmotedu/errors/example_test.go:101
	// testing.runExample
	//         /home/dfc/go/src/testing/example.go:114
	// testing.RunExamples
	//         /home/dfc/go/src/testing/example.go:38
	// testing.(*M).Run
	//         /home/dfc/go/src/testing/testing.go:744
	// main.main
	//         /github.com/marmotedu/errors/_test/_testmain.go:102
	// runtime.main
	//         /home/dfc/go/src/runtime/proc.go:183
	// runtime.goexit
	//         /home/dfc/go/src/runtime/asm_amd64.s:2059
}

func Example_stackTrace() {
	type stackTracer interface {
		StackTrace() StackTrace
	}

	err, ok := Cause(fn()).(stackTracer)
	if !ok {
		panic("oops, err does not implement stackTracer")
	}

	st := err.StackTrace()
	fmt.Printf("%+v", st[0:2]) // top two frames

	// Example output:
	// github.com/marmotedu/errors_test.fn
	//	/home/dfc/src/github.com/marmotedu/errors/example_test.go:47
	// github.com/marmotedu/errors_test.Example_stackTrace
	//	/home/dfc/src/github.com/marmotedu/errors/example_test.go:127
}

func ExampleCause_printf() {
	err := Wrap(func() error {
		return func() error {
			return New("hello world")
		}()
	}(), "failed")

	fmt.Printf("%v", err)

	// Output: failed
}

func ExampleWithCode() {
	var err error

	err = WithCode(ConfigurationNotValid, "this is an error message")
	fmt.Println(err)

	err = Wrap(err, "this is a wrap error message with error code not change")
	fmt.Println(err)

	err = WrapC(err, ErrInvalidJSON, "this is a wrap error message with new error code")
	fmt.Println(codes[err.(*withCode).code].String())
	fmt.Println(err)
	//fmt.Printf("%+v\n", err)
	//fmt.Printf("%#+v\n", err)

	// Output:
	// ConfigurationNotValid error
	// ConfigurationNotValid error
	// Data is not valid JSON
	// Data is not valid JSON
}

func ExamplewithCode_code() {
	err := loadConfig()
	if nil != err {
		err = WrapC(err, ErrLoadConfigFailed, "failed to load configuration")
	}

	fmt.Println(err.(*withCode).code)
	// Output: 1003
}

func ExampledefaultCoder_HTTPStatus() {
	err := loadConfig()
	if nil != err {
		err = WrapC(err, ErrLoadConfigFailed, "failed to load configuration")
	}

	fmt.Println(codes[err.(*withCode).code].HTTPStatus())
	// Output: 500
}

func ExampleCoder_HTTPStatus() {
	err := loadConfig()
	if nil != err {
		err = WrapC(err, ErrLoadConfigFailed, "failed to load configuration")
	}

	coder := ParseCoder(err)
	fmt.Println(coder.HTTPStatus())
	// Output: 500
}

func ExampleString() {
	err := loadConfig()
	if nil != err {
		err = WrapC(err, ErrLoadConfigFailed, "failed to load configuration")
	}

	fmt.Println(codes[err.(*withCode).code].String())
	// Output: Load configuration file failed
}
