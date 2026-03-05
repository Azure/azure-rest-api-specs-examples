package armdomainregistration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/domainregistration/armdomainregistration"
)

// Generated from example definition: 2024-11-01/ListDomainRecommendations.json
func ExampleDomainsClient_NewListRecommendationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdomainregistration.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDomainsClient().NewListRecommendationsPager(armdomainregistration.DomainRecommendationSearchParameters{
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
		// page = armdomainregistration.DomainsClientListRecommendationsResponse{
		// 	NameIdentifierCollection: armdomainregistration.NameIdentifierCollection{
		// 		Value: []*armdomainregistration.NameIdentifier{
		// 			{
		// 				Name: to.Ptr("domainnamesuggestion1"),
		// 			},
		// 			{
		// 				Name: to.Ptr("domainnamesuggestion2"),
		// 			},
		// 			{
		// 				Name: to.Ptr("domainnamesuggestion3"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
