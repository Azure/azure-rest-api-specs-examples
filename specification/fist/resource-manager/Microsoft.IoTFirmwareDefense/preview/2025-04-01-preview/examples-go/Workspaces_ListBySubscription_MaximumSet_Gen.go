package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense/v2"
)

// Generated from example definition: 2025-04-01-preview/Workspaces_ListBySubscription_MaximumSet_Gen.json
func ExampleWorkspacesClient_NewListBySubscriptionPager_workspacesListBySubscriptionMaximumSetGenGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("5C707B5F-6130-4F71-819E-953A28942E88", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacesClient().NewListBySubscriptionPager(nil)
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
		// page = armiotfirmwaredefense.WorkspacesClientListBySubscriptionResponse{
		// 	WorkspaceListResult: armiotfirmwaredefense.WorkspaceListResult{
		// 		Value: []*armiotfirmwaredefense.Workspace{
		// 			{
		// 				Location: to.Ptr("emiscxuo"),
		// 				ID: to.Ptr("/subscriptions/07aed47b-60ad-4d6e-a07a-000000000000/resourceGroups/FirmwareAnalysisRG/providers/Microsoft.IoTFirmwareDefense/workspaces/default/firmwares/109a9886-50bf-85a8-9d75-000000000000/summaries/firmware"),
		// 				Properties: &armiotfirmwaredefense.WorkspaceProperties{
		// 					ProvisioningState: to.Ptr(armiotfirmwaredefense.ProvisioningStateSucceeded),
		// 				},
		// 				SKU: &armiotfirmwaredefense.SKU{
		// 					Name: to.Ptr("kigqyklxhxoml"),
		// 					Tier: to.Ptr(armiotfirmwaredefense.SKUTierFree),
		// 					Size: to.Ptr("jtw"),
		// 					Family: to.Ptr("krrulu"),
		// 					Capacity: to.Ptr[int32](2),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key9484": to.Ptr("emizfeemymb"),
		// 				},
		// 				Name: to.Ptr("yfczfjzlijdcvsbsncvy"),
		// 				Type: to.Ptr("qlemis"),
		// 				SystemData: &armiotfirmwaredefense.SystemData{
		// 					CreatedBy: to.Ptr("nqisshvdzqcxzbujvacin"),
		// 					CreatedByType: to.Ptr(armiotfirmwaredefense.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-06-13T15:22:45.940Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("of"),
		// 					LastModifiedByType: to.Ptr(armiotfirmwaredefense.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-06-13T15:22:45.940Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
