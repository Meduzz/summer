package framework

var (
	Instance *Summer
)

func init() {
	Instance = NewFramework()
}
