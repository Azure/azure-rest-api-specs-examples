package armsearch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/search/resource-manager/Microsoft.Search/preview/2024-03-01-preview/examples/SearchUpdateServiceToAllowAccessFromPublicCustomIPs.json
func ExampleServicesClient_Update_searchUpdateServiceToAllowAccessFromPublicCustomIPs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsearch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().Update(ctx, "rg1", "mysearchservice", armsearch.ServiceUpdate{
		Properties: &armsearch.ServiceProperties{
			NetworkRuleSet: &armsearch.NetworkRuleSet{
				IPRules: []*armsearch.IPRule{
					{
						Value: to.Ptr("123.4.5.6"),
					},
					{
						Value: to.Ptr("123.4.6.0/18"),
					}},
			},
			PartitionCount:      to.Ptr[int32](1),
			PublicNetworkAccess: to.Ptr(armsearch.PublicNetworkAccessEnabled),
			ReplicaCount:        to.Ptr[int32](3),
		},
	}, &armsearch.SearchManagementRequestOptions{ClientRequestID: nil}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
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
	// 		"new-tag": to.Ptr("Adding a new tag"),
	// 	},
	// 	Properties: &armsearch.ServiceProperties{
	// 		AuthOptions: &armsearch.DataPlaneAuthOptions{
	// 			APIKeyOnly: map[string]any{
	// 			},
	// 		},
	// 		DisableLocalAuth: to.Ptr(false),
	// 		DisabledDataExfiltrationOptions: []*armsearch.SearchDisabledDataExfiltrationOption{
	// 		},
	// 		EncryptionWithCmk: &armsearch.EncryptionWithCmk{
	// 			EncryptionComplianceStatus: to.Ptr(armsearch.SearchEncryptionComplianceStatusCompliant),
	// 			Enforcement: to.Ptr(armsearch.SearchEncryptionWithCmkUnspecified),
	// 		},
	// 		HostingMode: to.Ptr(armsearch.HostingModeDefault),
	// 		NetworkRuleSet: &armsearch.NetworkRuleSet{
	// 			IPRules: []*armsearch.IPRule{
	// 				{
	// 					Value: to.Ptr("10.2.3.4"),
	// 			}},
	// 		},
	// 		PartitionCount: to.Ptr[int32](1),
	// 		PrivateEndpointConnections: []*armsearch.PrivateEndpointConnection{
	// 		},
	// 		ProvisioningState: to.Ptr(armsearch.ProvisioningStateSucceeded),
	// 		PublicNetworkAccess: to.Ptr(armsearch.PublicNetworkAccessEnabled),
	// 		ReplicaCount: to.Ptr[int32](3),
	// 		SharedPrivateLinkResources: []*armsearch.SharedPrivateLinkResource{
	// 		},
	// 		Status: to.Ptr(armsearch.SearchServiceStatusRunning),
	// 		StatusDetails: to.Ptr(""),
	// 	},
	// 	SKU: &armsearch.SKU{
	// 		Name: to.Ptr(armsearch.SKUNameStandard),
	// 	},
	// }
}
