package armquantum_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quantum/armquantum"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/quantum/resource-manager/Microsoft.Quantum/preview/2022-01-10-preview/examples/quantumWorkspacesListResourceGroup.json
func ExampleWorkspacesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquantum.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacesClient().NewListByResourceGroupPager("quantumResourcegroup", nil)
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
		// page.WorkspaceListResult = armquantum.WorkspaceListResult{
		// 	Value: []*armquantum.Workspace{
		// 		{
		// 			Name: to.Ptr("quantumworkspace1"),
		// 			Type: to.Ptr("Microsoft.Quantum/Workspaces"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/quantumResourcegroup/providers/Microsoft.Quantum/Workspaces/quantumworkspace1"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"company": to.Ptr("Contoso"),
		// 				"department": to.Ptr("MightyMight"),
		// 			},
		// 			Identity: &armquantum.WorkspaceIdentity{
		// 				Type: to.Ptr(armquantum.ResourceIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 				TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 			},
		// 			Properties: &armquantum.WorkspaceResourceProperties{
		// 				EndpointURI: to.Ptr("https://quantumworkspace1.westus.quantum.azure.com"),
		// 				Providers: []*armquantum.Provider{
		// 					{
		// 						ApplicationName: to.Ptr("quantumworkspace1-h1"),
		// 						InstanceURI: to.Ptr("https://h1.endpoint.com"),
		// 						ProviderID: to.Ptr("Honeywell"),
		// 						ProviderSKU: to.Ptr("Basic"),
		// 						ProvisioningState: to.Ptr(armquantum.StatusSucceeded),
		// 						ResourceUsageID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 				}},
		// 				ProvisioningState: to.Ptr(armquantum.ProvisioningStatusSucceeded),
		// 				StorageAccount: to.Ptr("/subscriptions/1C4B2828-7D49-494F-933D-061373BE28C2/resourceGroups/quantumResourcegroup/providers/Microsoft.Storage/storageAccounts/testStorageAccount"),
		// 				Usable: to.Ptr(armquantum.UsableStatusYes),
		// 			},
		// 			SystemData: &armquantum.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armquantum.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armquantum.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
