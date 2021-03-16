package example

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
	"unsafe"
)

/*
go test -v fib_test.go -test.run TestAssignment
*/

// assignment
func TestAssignment(t *testing.T) {
	var a int
	a = 1
	var b = 2
	c := 3
	t.Log(a, b, c)
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	b, a = a, b
	t.Log(a, b)
}

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Saturday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday, Saturday)
	t.Log(Readable, Writable, Executable)
}

// type alias
type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64 = 2
	var c MyInt

	//c = b
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

type user struct {
	name string
}

// format
func TestFormat(t *testing.T) {
	u := user{"damao"}
	//Printf 格式化输出
	fmt.Printf("%+v\n", u)       //格式化输出结构
	fmt.Printf("%#v\n", u)       //输出值的 Go 语言表示方法
	fmt.Printf("%T\n", u)        //输出值的类型的 Go 语言表示
	fmt.Printf("%t\n", true)     //输出值的 true 或 false
	fmt.Printf("%b\n", 1024)     //二进制表示
	fmt.Printf("%c\n", 11111111) //数值对应的 Unicode 编码字符
	fmt.Printf("%d\n", 10)       //十进制表示
	fmt.Printf("%o\n", 8)        //八进制表示
	fmt.Printf("%q\n", 22)       //转化为十六进制并附上单引号
	fmt.Printf("%x\n", 1223)     //十六进制表示，用a-f表示
	fmt.Printf("%X\n", 1223)     //十六进制表示，用A-F表示
	fmt.Printf("%U\n", 1233)     //Unicode表示
	fmt.Printf("%b\n", 12.34)    //无小数部分，两位指数的科学计数法6946802425218990p-49
	fmt.Printf("%e\n", 12.345)   //科学计数法，e表示
	fmt.Printf("%E\n", 12.34455) //科学计数法，E表示
	fmt.Printf("%f\n", 12.3456)  //有小数部分，无指数部分
	fmt.Printf("%g\n", 12.3456)  //根据实际情况采用%e或%f输出
	fmt.Printf("%G\n", 12.3456)  //根据实际情况采用%E或%f输出
	fmt.Printf("%s\n", "wqdew")  //直接输出字符串或者[]byte
	fmt.Printf("%q\n", "dedede") //双引号括起来的字符串
	fmt.Printf("%x\n", "abczxc") //每个字节用两字节十六进制表示，a-f表示
	fmt.Printf("%X\n", "asdzxc") //每个字节用两字节十六进制表示，A-F表示
}

func TestString(t *testing.T) {
	var a string
	t.Log("*" + a + "*")
}

// loop
func TestFibList(t *testing.T) {
	a := 1
	b := 1
	fmt.Print(a)
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}

func TestWhileLoop(t *testing.T) {
	n := 0
	// while n <5
	for n < 5 {
		t.Log(n)
		n++
	}
}

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	aa := [...]int{1, 2, 3, 4}
	//b := [...]int{1,2,3,4,5}
	//c := [...]int{1,3,2}
	//t.Log(a==b)
	//t.Log(a==c)
	t.Log(a == aa)

}

func TestArrayInt(t *testing.T) {
	var arr [3]int
	arr2 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5}
	arr[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr2, arr3)
	t.Log(arr2[1:])

	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	for idx, e := range arr3 {
		t.Log(idx, e)
	}
}

// condition
func TestIfMultiSec(t *testing.T) {
	b := 1
	if a := b == 1; a {
		t.Log("1==1")
	}
}

// switch
func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("4")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("?")

		}
	}
}

// array
func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	arr_sec := arr[1:]
	t.Log(arr)
	t.Log(arr_sec)

	arr_sec[0] = 100
	t.Log(arr)
	t.Log(arr_sec)
}

// slice
func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	//t.Log(s2[4])

	s3 := []int{}
	for i := 0; i < 10; i++ {
		s3 = append(s3, i)
		t.Log(len(s3), cap(s3))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "unknow"

	t.Log(year)
	t.Log(summer)
	t.Log(Q2)
}

// map
func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 5: 9}
	t.Log(m1[2])
	t.Log(m1)
	t.Logf("len m1=%d", len(m1))
	//m2 := map[int]
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	if v, ok := m1[3]; ok {
		t.Log(ok)
		t.Logf("Key 3's b value is %d", v)
	} else {
		t.Log(ok)
		t.Log("Key 3 is not existing.")
	}
}

func TestTraveMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 5: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return 1 * op }
	m[2] = func(op int) int { return 2 * op }
	m[3] = func(op int) int { return 3 * op }
	m[4] = func(op int) int { return 4 * op }
	t.Log(m[1](2), m[2](2))
	t.Log(m)
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 1
	if mySet[n] {
		t.Log("existing")
	} else {
		t.Log("not existing")
	}
	t.Log(len(mySet))
	delete(mySet, n)
}

func TestString2(t *testing.T) {
	s := "中"
	t.Log(s)
	t.Log(len(s)) //byte数量

	c := []rune(s)
	t.Log(len(c))

	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)

	s2 := "你好，世界"
	for _, c := range s2 {
		t.Logf("%[1]c %[1]d %[1]x", c)
	}

	s3 := "A,B,C"
	parts := strings.Split(s3, ",")
	t.Log(s3)
	t.Log(parts)
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))

	s4 := strconv.Itoa(10)
	t.Log("str" + s4)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}

func returnMultiValues() (int, int) {
	return rand.Intn(11), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
}

// closure
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)

		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn2(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFunc)
	t.Log(tsSF(10))
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resources.")
}

// panic
// defer
func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	panic("err")
}

type Employee struct {
	Id   string
	Name string
	Age  int
}

// struct
func TestCreatEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee) //返回指针
	e2.Id = "2"
	e2.Age = 22
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	e1.Age = 300
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)
	t.Log(e.String())
	fmt.Printf("origin Address is %x\n", unsafe.Pointer(&e.Name))
	e3 := &Employee{"0", "Bob", 20}
	t.Logf(e3.String())

	t.Log(e.String2())
	t.Logf(e3.String2())
}

func (e Employee) String() string { //复制了一份数据
	fmt.Printf("instance Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func (e *Employee) String2() string { //地址不变
	fmt.Printf("pointer Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g GoProgrammer) WriteHelloWorld() string {
	return "hello world"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}

type IntConv func(op int) int

func timeSpent2(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func TestFn3(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent2(slowFunc)
	t.Log(tsSF(10))
}

// object oriented
type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("hello")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	p *Pet
}

func (d *Dog) Speak() {
	d.p.Speak()
}

func (d *Dog) SpeakTo(host string) {
	//d.Speak()
	//fmt.Println(" ", host)
	d.p.SpeakTo(host)
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("damao")
}

type Dog2 struct {
	Pet
}

func (d *Dog2) Speak() {
	fmt.Println("hi")
}

func TestDog2(t *testing.T) {
	dog := new(Dog2)
	dog.SpeakTo("damao")
}

// polymorphism
type Code string
type Coder interface {
	WriteHelloWorld() Code
}

type GoCoder struct {
}

func (g *GoCoder) WriteHelloWorld() Code {
	return "hello go coder"
}

type PythonCoder struct {
}

func (p *PythonCoder) WriteHelloWorld() Code {
	return "hello python coder"
}

func writeFirstProgram(c Coder) {
	fmt.Printf("%T %v\n", c, c.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goer := new(GoCoder) // goer := &GoCoder{}
	pyer := new(PythonCoder)
	writeFirstProgram(goer)
	writeFirstProgram(pyer)
}

func DoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
		return
	}
	if s, ok := p.(string); ok {
		fmt.Println("string", s)
		return
	}
	fmt.Println("Unknow Type")
}

func DoSomething2(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknow")

	}
}

func TestEmptyInterFaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
	DoSomething2(10)
	DoSomething2("10")
}

func Getfibonacci(n int) ([]int, error) {
	if n < 0 {
		return nil, LessERROR
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

var LessERROR = errors.New("<2")

func TestGetfibonacci(t *testing.T) {
	if v, err := Getfibonacci(-10); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestPanicVxExit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("wrong!"))
}

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}

// lock
func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterWithLock(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter = %d", counter)
}

// csp
func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Microsecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	//retCh := make(chan string)
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsynService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Microsecond * 100000):
		t.Error("time out")
	}
}

//close channel
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}

// cancel
func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Microsecond * 5)
			}
			fmt.Println(i, "canceled")
		}(i, cancelChan)
	}
	//cancel_1(cancelChan)
	cancel_2(cancelChan)
	time.Sleep(time.Second * 1)
}

/*
context

根context: 通过context.Background() 创建
子context: context.WithCancel(parentContext) 创建
  ctx, cancel := context.WithCancel(context.Background())
当前context被取消时，基于他的子context都会被取消
接收取消通知 <- ctx.Done()
*/

func isCancelledWithContext(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancelWithContext(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelledWithContext(ctx) {
					break
				}
				time.Sleep(time.Microsecond * 5)
			}
			fmt.Println(i, "canceled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}

// singleton
type Singleton struct {
}

var singleInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			//fmt.Println(obj)
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}

// any task return, main func return
func runTask(id int) string {
	time.Sleep(10 * time.Microsecond)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}

// all task return, main func return
func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	finalRet := ""
	for j := 0; j < numOfRunner; j++ {
		finalRet += <-ch + "\n"
	}
	return finalRet
}

func TestAllResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}

// example: pool
type ResuableObj struct {
}

type ObjPool struct {
	bufChan chan *ResuableObj //用于缓冲可重用对象
}

func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ResuableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ResuableObj{}
	}
	return &objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ResuableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout): //超时控制
		return nil, errors.New("tiem out")
	}
}

func (p *ObjPool) RelaseObj(obj *ResuableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			if err := pool.RelaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
	fmt.Println("Done")
}
