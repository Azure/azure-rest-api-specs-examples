package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/domainregistration/resource-manager/Microsoft.DomainRegistration/DomainRegistration/stable/2024-11-01/examples/ListDomainRecommendations.json
func ExampleDomainsClient_NewListRecommendationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDomainsClient().NewListRecommendationsPager(armappservice.DomainRecommendationSearchParameters{
		Keywords:                 to.Ptr("example1"),
		MaxDomainRecommendations: to.Ptr[int32](10),
	}, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.NameIdentifierCollection = armappservice.NameIdentifierCollection{
		// 	Value: []*armappservice.NameIdentifier{
		// 		{
		// 			Name: to.Ptr("domainnamesuggestion1"),
		// 		},
		// 		{
		// 			Name: to.Ptr("domainnamesuggestion2"),
		// 		},
		// 		{
		// 			Name: to.Ptr("domainnamesuggestion3"),
		// 	}},
		// }
	}
}
