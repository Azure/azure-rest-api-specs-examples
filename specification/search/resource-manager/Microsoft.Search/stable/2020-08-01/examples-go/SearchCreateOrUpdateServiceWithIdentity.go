package armsearch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/SearchCreateOrUpdateServiceWithIdentity.json
func ExampleServicesClient_BeginCreateOrUpdate_searchCreateOrUpdateServiceWithIdentity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsearch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServicesClient().BeginCreateOrUpdate(ctx, "rg1", "mysearchservice", armsearch.Service{
		Location: to.Ptr("westus"),
		Tags: map[string]*string{
			"app-name": to.Ptr("My e-commerce app"),
		},
		Identity: &armsearch.Identity{
			Type: to.Ptr(armsearch.IdentityTypeSystemAssigned),
		},
		Properties: &armsearch.ServiceProperties{
			HostingMode:    to.Ptr(armsearch.HostingModeDefault),
			PartitionCount: to.Ptr[int32](1),
			ReplicaCount:   to.Ptr[int32](3),
		},
		SKU: &armsearch.SKU{
			Name: to.Ptr(armsearch.SKUNameStandard),
		},
	}, &armsearch.SearchManagementRequestOptions{ClientRequestID: nil}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Service = armsearch.Service{
	// 	Name: to.Ptr("mysearchservice"),
	// 	Type: to.Ptr("Microsoft.Search/searchServices"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Search/searchServices/mysearchservice"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"app-name": to.Ptr("My e-commerce app"),
	// 	},
	// 	Identity: &armsearch.Identity{
	// 		Type: to.Ptr(armsearch.IdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("9d1e1f18-2122-4988-a11c-878782e40a5c"),
	// 		TenantID: to.Ptr("f686d426-8d16-42db-81b7-ab578e110ccd"),
	// 	},
	// 	Properties: &armsearch.ServiceProperties{
	// 		HostingMode: to.Ptr(armsearch.HostingModeDefault),
	// 		NetworkRuleSet: &armsearch.NetworkRuleSet{
	// 			IPRules: []*armsearch.IPRule{
	// 			},
	// 		},
	// 		PartitionCount: to.Ptr[int32](1),
	// 		PrivateEndpointConnections: []*armsearch.PrivateEndpointConnection{
	// 		},
	// 		ProvisioningState: to.Ptr(armsearch.ProvisioningStateSucceeded),
	// 		PublicNetworkAccess: to.Ptr(armsearch.PublicNetworkAccessEnabled),
	// 		ReplicaCount: to.Ptr[int32](3),
	// 		Status: to.Ptr(armsearch.SearchServiceStatusProvisioning),
	// 		StatusDetails: to.Ptr(""),
	// 	},
	// 	SKU: &armsearch.SKU{
	// 		Name: to.Ptr(armsearch.SKUNameStandard),
	// 	},
	// }
}
