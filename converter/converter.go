package converter

type Converter interface {
	Convert(amount float64, from, to string) (float64, error)
}
