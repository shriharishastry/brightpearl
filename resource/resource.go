package resource

type Resource struct {
	Product ProductResource
	Category CategoryResource
	Brand BrandtResource
	Season SeasonResource
	Option OptionResource
	OptionValue OptionValue
	ProductType ProductTypeResource
	Collection CollectionResource

}

func New() *Resource {
	return &Resource{}
}