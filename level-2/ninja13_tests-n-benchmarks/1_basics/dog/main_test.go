package dog

/*
1) go test
2) godoc -http=0.0.0.0:8080
3) go test -bench .
4) go test -cover

5) go test -coverprofile c.out && go tool cover -html=c.out
*/
import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	n := Years(10)
	if n != 70 {
		t.Error("expected", 70, "got", n)
	}
}

func TestYearsTwo(t *testing.T) {
	n := YearsTwo(10)
	if n != 70 {
		t.Error("expected", 70, "got", n)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	// Output:
	// 70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
