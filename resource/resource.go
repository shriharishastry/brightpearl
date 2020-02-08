package resource

import "github.com/shriharishastry/brightpearl/connector"

const (
	ProductResourceURL string = "products"
	CategoryResourceURL string = "category"
	BrandResourceURL string = "brand"
	CollectionResourceURL string = "collection"
	SeasonResourceURL string = "season"
	ProductTypeResourceURL string = "product_type"
	CustomFieldResourceURL string = "custom_field"
	OptionResourceURL string = "option"
	OptionValueResourceURL string = "option_value"
)

type Resource struct {
	Product ProductResource
	Category CategoryResource
	Brand BrandtResource
	Season SeasonResource
	Option OptionResource
	OptionValue OptionValueResource
	ProductType ProductTypeResource
	Collection CollectionResource
	CustomField CustomFieldResource

}

func New(client connector.HttpClient) *Resource {
	return &Resource{
		Product: ProductResource{resourceUrl:ProductResourceURL, connection: client},
		Category:CategoryResource{resourceUrl:CategoryResourceURL, connection:client},
		Brand:BrandtResource{resourceUrl:BrandResourceURL, connection:client},
		Collection:CollectionResource{resourceUrl:CollectionResourceURL, connection:client},
		Season:SeasonResource{resourceUrl: SeasonResourceURL, connection:client},
		ProductType:ProductTypeResource{resourceUrl:ProductTypeResourceURL, connection:client},
		CustomField: CustomFieldResource{resourceUrl:CustomFieldResourceURL, connection:client},
		OptionValue:OptionValueResource{resourceUrl: OptionValueResourceURL, connection:client},
		Option:OptionResource{resourceUrl:OptionResourceURL, connection:client},
	}
}