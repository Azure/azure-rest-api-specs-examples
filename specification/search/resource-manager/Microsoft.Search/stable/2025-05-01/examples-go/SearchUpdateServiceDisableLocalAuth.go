package armsearch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fb3217991ff57b5760525aeba1a0670bfe0880fa/specification/search/resource-manager/Microsoft.Search/stable/2025-05-01/examples/SearchUpdateServiceDisableLocalAuth.json
func ExampleServicesClient_Update_searchUpdateServiceDisableLocalAuth() {
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
			DisableLocalAuth: to.Ptr(true),
			ReplicaCount:     to.Ptr[int32](2),
		},
		Tags: map[string]*string{
			"app-name": to.Ptr("My e-commerce app"),
			"new-tag":  to.Ptr("Adding a new tag"),
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
	// 	SystemData: &armsearch.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-01T00:00:00.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("My e-commerce app"),
	// 		CreatedByType: to.Ptr(armsearch.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-01T00:00:00.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("fakeuser@contoso.com"),
	// 		LastModifiedByType: to.Ptr(armsearch.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"app-name": to.Ptr("My e-commerce app"),
	// 		"new-tag": to.Ptr("Adding a new tag"),
	// 	},
	// 	Properties: &armsearch.ServiceProperties{
	// 		ComputeType: to.Ptr(armsearch.ComputeTypeDefault),
	// 		DataExfiltrationProtections: []*armsearch.SearchDataExfiltrationProtection{
	// 		},
	// 		DisableLocalAuth: to.Ptr(true),
	// 		EncryptionWithCmk: &armsearch.EncryptionWithCmk{
	// 			EncryptionComplianceStatus: to.Ptr(armsearch.SearchEncryptionComplianceStatusCompliant),
	// 			Enforcement: to.Ptr(armsearch.SearchEncryptionWithCmkUnspecified),
	// 		},
	// 		Endpoint: to.Ptr("https://mysearchservice.search.windows.net/"),
	// 		HostingMode: to.Ptr(armsearch.HostingModeDefault),
	// 		NetworkRuleSet: &armsearch.NetworkRuleSet{
	// 			Bypass: to.Ptr(armsearch.SearchBypassNone),
	// 			IPRules: []*armsearch.IPRule{
	// 			},
	// 		},
	// 		PartitionCount: to.Ptr[int32](1),
	// 		PrivateEndpointConnections: []*armsearch.PrivateEndpointConnection{
	// 		},
	// 		ProvisioningState: to.Ptr(armsearch.ProvisioningStateSucceeded),
	// 		PublicNetworkAccess: to.Ptr(armsearch.PublicNetworkAccessEnabled),
	// 		ReplicaCount: to.Ptr[int32](2),
	// 		SharedPrivateLinkResources: []*armsearch.SharedPrivateLinkResource{
	// 		},
	// 		Status: to.Ptr(armsearch.SearchServiceStatusProvisioning),
	// 		StatusDetails: to.Ptr(""),
	// 		UpgradeAvailable: to.Ptr(armsearch.UpgradeAvailableNotAvailable),
	// 	},
	// 	SKU: &armsearch.SKU{
	// 		Name: to.Ptr(armsearch.SKUNameStandard),
	// 	},
	// }
}
