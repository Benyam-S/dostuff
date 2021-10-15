package translate

// IService is an interface that defines all the functions that the translate service provides
type IService interface {
	TranslateStuff(location string) (string, error)
}
