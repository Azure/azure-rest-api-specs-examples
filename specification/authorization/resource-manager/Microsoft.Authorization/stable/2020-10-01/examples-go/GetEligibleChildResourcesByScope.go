package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/GetEligibleChildResourcesByScope.json
func ExampleEligibleChildResourcesClient_NewGetPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEligibleChildResourcesClient().NewGetPager("providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f", &armauthorization.EligibleChildResourcesClientGetOptions{Filter: to.Ptr("resourceType+eq+'resourcegroup'")})
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
		// page.EligibleChildResourcesListResult = armauthorization.EligibleChildResourcesListResult{
		// 	Value: []*armauthorization.EligibleChildResource{
		// 		{
		// 			Name: to.Ptr("RG-1"),
		// 			Type: to.Ptr("resourcegroup"),
		// 			ID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/resourceGroups/RG-1"),
		// 		},
		// 		{
		// 			Name: to.Ptr("RG-2"),
		// 			Type: to.Ptr("resourcegroup"),
		// 			ID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/resourceGroups/RG-2"),
		// 	}},
		// }
	}
}
