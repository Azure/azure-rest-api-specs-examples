package armdeviceregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry/v2"
)

// Generated from example definition: 2026-03-01-preview/List_Credentials_ByResourceGroup.json
func ExampleCredentialsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCredentialsClient().NewListByResourceGroupPager("rgdeviceregistry", "mynamespace", nil)
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
		// page = armdeviceregistry.CredentialsClientListByResourceGroupResponse{
		// 	CredentialListResult: armdeviceregistry.CredentialListResult{
		// 		Value: []*armdeviceregistry.Credential{
		// 			{
		// 				Properties: &armdeviceregistry.CredentialProperties{
		// 					ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key7121": to.Ptr("mtdjqipusqaqhdvekrknyjeo"),
		// 				},
		// 				Location: to.Ptr("East US 2"),
		// 				ID: to.Ptr("/subscriptions/00000000-1366-430f-0000-cc873bcf2d27/resourceGroups/gbktestRG/providers/Microsoft.DeviceRegistry/namespaces/adr-ns-gbk-01/credentials/default"),
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.DeviceRegistry/namespaces/credentials"),
		// 				SystemData: &armdeviceregistry.SystemData{
		// 					CreatedBy: to.Ptr("00003442-0000-0000-0000-494059220000"),
		// 					CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-09-13T19:38:09.5283958Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("00003442-0000-0000-0000-494059220000"),
		// 					LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-09-13T19:38:16.6634263Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
