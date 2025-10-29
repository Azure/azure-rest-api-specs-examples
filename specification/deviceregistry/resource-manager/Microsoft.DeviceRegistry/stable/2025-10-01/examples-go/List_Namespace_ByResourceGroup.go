package armdeviceregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry/v2"
)

// Generated from example definition: 2025-10-01/List_Namespace_ByResourceGroup.json
func ExampleNamespacesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNamespacesClient().NewListByResourceGroupPager("myResourceGroup", nil)
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
		// page = armdeviceregistry.NamespacesClientListByResourceGroupResponse{
		// 	NamespaceListResult: armdeviceregistry.NamespaceListResult{
		// 		Value: []*armdeviceregistry.Namespace{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-1366-430f-0000-cc873bcf2d27/resourceGroups/gbktestRG/providers/Microsoft.DeviceRegistry/namespaces/adr-ns-gbk-01"),
		// 				Name: to.Ptr("adr-ns-gbk-01"),
		// 				Type: to.Ptr("Microsoft.DeviceRegistry/namespaces"),
		// 				Location: to.Ptr("North Europe"),
		// 				SystemData: &armdeviceregistry.SystemData{
		// 					CreatedBy: to.Ptr("00003442-0000-0000-0000-494059220000"),
		// 					CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-09-13T19:38:09.5283958Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("00003442-0000-0000-0000-494059220000"),
		// 					LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-09-13T19:38:16.6634263Z"); return t}()),
		// 				},
		// 				Properties: &armdeviceregistry.NamespaceProperties{
		// 					ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-1366-430f-0000-cc873bcf2d27/resourceGroups/gbktestRG/providers/Microsoft.DeviceRegistry/namespaces/adr-ns-gbk-02"),
		// 				Name: to.Ptr("adr-ns-gbk-02"),
		// 				Type: to.Ptr("Microsoft.DeviceRegistry/namespaces"),
		// 				Location: to.Ptr("North Europe"),
		// 				SystemData: &armdeviceregistry.SystemData{
		// 					CreatedBy: to.Ptr("00003442-0000-0000-0000-494059220000"),
		// 					CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-09-13T20:22:49.5920101Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("00003442-0000-0000-0000-494059220000"),
		// 					LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-09-13T20:22:49.5920101Z"); return t}()),
		// 				},
		// 				Properties: &armdeviceregistry.NamespaceProperties{
		// 					ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
		// 					UUID: to.Ptr("cbfe124a-6971-4c90-a7a9-99be82def1ab"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
