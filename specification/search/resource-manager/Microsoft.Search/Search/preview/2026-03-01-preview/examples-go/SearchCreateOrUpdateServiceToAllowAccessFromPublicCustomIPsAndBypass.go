package armsearch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch/v2"
)

// Generated from example definition: 2026-03-01-preview/SearchCreateOrUpdateServiceToAllowAccessFromPublicCustomIPsAndBypass.json
func ExampleServicesClient_BeginCreateOrUpdate_searchCreateOrUpdateServiceToAllowAccessFromPublicCustomIPsAndBypass() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsearch.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServicesClient().BeginCreateOrUpdate(ctx, "rg1", "mysearchservice", armsearch.Service{
		Location: to.Ptr("westus"),
		Tags: map[string]*string{
			"app-name": to.Ptr("My e-commerce app"),
		},
		SKU: &armsearch.SKU{
			Name: to.Ptr(armsearch.SKUNameStandard),
		},
		Properties: &armsearch.ServiceProperties{
			ReplicaCount:   to.Ptr[int32](1),
			PartitionCount: to.Ptr[int32](1),
			NetworkRuleSet: &armsearch.NetworkRuleSet{
				IPRules: []*armsearch.IPRule{
					{
						Value: to.Ptr("123.4.5.6"),
					},
					{
						Value: to.Ptr("123.4.6.0/18"),
					},
				},
				Bypass: to.Ptr(armsearch.SearchBypassAzurePortal),
			},
			HostingMode: to.Ptr(armsearch.HostingModeDefault),
			ComputeType: to.Ptr(armsearch.ComputeTypeDefault),
		},
	}, nil)
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
	// res = armsearch.ServicesClientCreateOrUpdateResponse{
	// 	Service: &armsearch.Service{
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Search/searchServices/mysearchservice"),
	// 		Name: to.Ptr("mysearchservice"),
	// 		Location: to.Ptr("westus"),
	// 		Type: to.Ptr("Microsoft.Search/searchServices"),
	// 		Tags: map[string]*string{
	// 			"app-name": to.Ptr("My e-commerce app"),
	// 		},
	// 		SKU: &armsearch.SKU{
	// 			Name: to.Ptr(armsearch.SKUNameStandard),
	// 		},
	// 		Properties: &armsearch.ServiceProperties{
	// 			ReplicaCount: to.Ptr[int32](1),
	// 			PartitionCount: to.Ptr[int32](1),
	// 			Endpoint: to.Ptr("https://mysearchservice.search.windows.net/"),
	// 			Status: to.Ptr(armsearch.SearchServiceStatusProvisioning),
	// 			StatusDetails: to.Ptr(""),
	// 			HostingMode: to.Ptr(armsearch.HostingModeDefault),
	// 			ComputeType: to.Ptr(armsearch.ComputeTypeDefault),
	// 			ProvisioningState: to.Ptr(armsearch.ProvisioningStateProvisioning),
	// 			PublicNetworkAccess: to.Ptr(armsearch.PublicNetworkAccessEnabled),
	// 			NetworkRuleSet: &armsearch.NetworkRuleSet{
	// 				IPRules: []*armsearch.IPRule{
	// 					{
	// 						Value: to.Ptr("123.4.5.6"),
	// 					},
	// 					{
	// 						Value: to.Ptr("123.4.6.0/18"),
	// 					},
	// 				},
	// 				Bypass: to.Ptr(armsearch.SearchBypassAzurePortal),
	// 			},
	// 			PrivateEndpointConnections: []*armsearch.PrivateEndpointConnection{
	// 			},
	// 			SharedPrivateLinkResources: []*armsearch.SharedPrivateLinkResource{
	// 			},
	// 			EncryptionWithCmk: &armsearch.EncryptionWithCmk{
	// 				Enforcement: to.Ptr(armsearch.SearchEncryptionWithCmkUnspecified),
	// 				EncryptionComplianceStatus: to.Ptr(armsearch.SearchEncryptionComplianceStatusCompliant),
	// 			},
	// 			DisableLocalAuth: to.Ptr(false),
	// 			AuthOptions: &armsearch.DataPlaneAuthOptions{
	// 				AADOrAPIKey: &armsearch.DataPlaneAADOrAPIKeyAuthOption{
	// 					AADAuthFailureMode: to.Ptr(armsearch.AADAuthFailureModeHttp401WithBearerChallenge),
	// 				},
	// 			},
	// 			DataExfiltrationProtections: []*armsearch.SearchDataExfiltrationProtection{
	// 			},
	// 			UpgradeAvailable: to.Ptr(armsearch.UpgradeAvailableNotAvailable),
	// 		},
	// 		SystemData: &armsearch.SystemData{
	// 			CreatedBy: to.Ptr("My e-commerce app"),
	// 			CreatedByType: to.Ptr(armsearch.CreatedByTypeApplication),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-01T00:00:00Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("fakeuser@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armsearch.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-01T00:00:00Z"); return t}()),
	// 		},
	// 	},
	// }
}
