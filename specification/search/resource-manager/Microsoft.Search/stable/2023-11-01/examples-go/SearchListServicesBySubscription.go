package armsearch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7e29dd59eef13ef347d09e41a63f2585be77b3ca/specification/search/resource-manager/Microsoft.Search/stable/2023-11-01/examples/SearchListServicesBySubscription.json
func ExampleServicesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsearch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServicesClient().NewListBySubscriptionPager(&armsearch.SearchManagementRequestOptions{ClientRequestID: nil}, nil)
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
		// page.ServiceListResult = armsearch.ServiceListResult{
		// 	Value: []*armsearch.Service{
		// 		{
		// 			Name: to.Ptr("mysearchservice"),
		// 			Type: to.Ptr("Microsoft.Search/searchServices"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Search/searchServices/mysearchservice"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"app-name": to.Ptr("My e-commerce app"),
		// 			},
		// 			Properties: &armsearch.ServiceProperties{
		// 				HostingMode: to.Ptr(armsearch.HostingModeDefault),
		// 				NetworkRuleSet: &armsearch.NetworkRuleSet{
		// 					IPRules: []*armsearch.IPRule{
		// 					},
		// 				},
		// 				PartitionCount: to.Ptr[int32](1),
		// 				PrivateEndpointConnections: []*armsearch.PrivateEndpointConnection{
		// 				},
		// 				ProvisioningState: to.Ptr(armsearch.ProvisioningStateSucceeded),
		// 				PublicNetworkAccess: to.Ptr(armsearch.PublicNetworkAccessEnabled),
		// 				ReplicaCount: to.Ptr[int32](3),
		// 				Status: to.Ptr(armsearch.SearchServiceStatusRunning),
		// 				StatusDetails: to.Ptr(""),
		// 			},
		// 			SKU: &armsearch.SKU{
		// 				Name: to.Ptr(armsearch.SKUNameStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("mysearchservice2"),
		// 			Type: to.Ptr("Microsoft.Search/searchServices"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Search/searchServices/mysearchservice2"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"app-name": to.Ptr("My e-commerce app"),
		// 			},
		// 			Properties: &armsearch.ServiceProperties{
		// 				HostingMode: to.Ptr(armsearch.HostingModeDefault),
		// 				NetworkRuleSet: &armsearch.NetworkRuleSet{
		// 					IPRules: []*armsearch.IPRule{
		// 					},
		// 				},
		// 				PartitionCount: to.Ptr[int32](1),
		// 				PrivateEndpointConnections: []*armsearch.PrivateEndpointConnection{
		// 				},
		// 				ProvisioningState: to.Ptr(armsearch.ProvisioningStateSucceeded),
		// 				PublicNetworkAccess: to.Ptr(armsearch.PublicNetworkAccessEnabled),
		// 				ReplicaCount: to.Ptr[int32](1),
		// 				Status: to.Ptr(armsearch.SearchServiceStatusRunning),
		// 				StatusDetails: to.Ptr(""),
		// 			},
		// 			SKU: &armsearch.SKU{
		// 				Name: to.Ptr(armsearch.SKUNameBasic),
		// 			},
		// 	}},
		// }
	}
}
