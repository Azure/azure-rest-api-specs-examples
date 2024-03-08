package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/stable/2023-11-15/examples/ApiCollections/APICollections_ListBySubscription_example.json
func ExampleAPICollectionsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAPICollectionsClient().NewListBySubscriptionPager(nil)
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
		// page.APICollectionList = armsecurity.APICollectionList{
		// 	Value: []*armsecurity.APICollection{
		// 		{
		// 			Name: to.Ptr("echo-api"),
		// 			Type: to.Ptr("Microsoft.Security/apiCollections"),
		// 			ID: to.Ptr("/subscriptions/3fa85f64-5717-4562-b3fc-2c963f66afa6/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/providers/Microsoft.Security/apiCollections/echo-api"),
		// 			Properties: &armsecurity.APICollectionProperties{
		// 				BaseURL: to.Ptr("https://apimservice1.azure-api.net/echo"),
		// 				DiscoveredVia: to.Ptr("/subscriptions/3fa85f64-5717-4562-b3fc-2c963f66afa6/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1"),
		// 				DisplayName: to.Ptr("Echo API"),
		// 				NumberOfAPIEndpoints: to.Ptr[int64](6),
		// 				NumberOfAPIEndpointsWithSensitiveDataExposed: to.Ptr[int64](1),
		// 				NumberOfExternalAPIEndpoints: to.Ptr[int64](3),
		// 				NumberOfInactiveAPIEndpoints: to.Ptr[int64](3),
		// 				NumberOfUnauthenticatedAPIEndpoints: to.Ptr[int64](1),
		// 				ProvisioningState: to.Ptr(armsecurity.ProvisioningStateSucceeded),
		// 				SensitivityLabel: to.Ptr("Highly Confidential"),
		// 			},
		// 	}},
		// }
	}
}
