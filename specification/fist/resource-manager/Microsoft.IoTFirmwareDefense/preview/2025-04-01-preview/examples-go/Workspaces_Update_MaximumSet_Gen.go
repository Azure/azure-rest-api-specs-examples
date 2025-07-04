package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense/v2"
)

// Generated from example definition: 2025-04-01-preview/Workspaces_Update_MaximumSet_Gen.json
func ExampleWorkspacesClient_Update_workspacesUpdateMaximumSetGenGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("5C707B5F-6130-4F71-819E-953A28942E88", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacesClient().Update(ctx, "rgiotfirmwaredefense", "exampleWorkspaceName", armiotfirmwaredefense.WorkspaceUpdate{
		Tags: map[string]*string{},
		SKU: &armiotfirmwaredefense.SKU{
			Name:     to.Ptr("jmlbmmdyyxoreypd"),
			Tier:     to.Ptr(armiotfirmwaredefense.SKUTierFree),
			Size:     to.Ptr("rkoairmk"),
			Family:   to.Ptr("jcrsluqmbovznq"),
			Capacity: to.Ptr[int32](3),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armiotfirmwaredefense.WorkspacesClientUpdateResponse{
	// 	Workspace: &armiotfirmwaredefense.Workspace{
	// 		Properties: &armiotfirmwaredefense.WorkspaceProperties{
	// 			ProvisioningState: to.Ptr(armiotfirmwaredefense.ProvisioningStateSucceeded),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key4978": to.Ptr("ujdbllhyqeunsgpqtiq"),
	// 		},
	// 		Location: to.Ptr("emiscxuo"),
	// 		SKU: &armiotfirmwaredefense.SKU{
	// 			Name: to.Ptr("pb"),
	// 			Tier: to.Ptr(armiotfirmwaredefense.SKUTierFree),
	// 			Size: to.Ptr("unh"),
	// 			Family: to.Ptr("fwsu"),
	// 			Capacity: to.Ptr[int32](22),
	// 		},
	// 		ID: to.Ptr("/subscriptions/07aed47b-60ad-4d6e-a07a-000000000000/resourceGroups/FirmwareAnalysisRG/providers/Microsoft.IoTFirmwareDefense/workspaces/default/firmwares/109a9886-50bf-85a8-9d75-000000000000/summaries/firmware"),
	// 		Name: to.Ptr("kgcqmntw"),
	// 		Type: to.Ptr("gtbncrjzxrcvzzzgyjheabpjzzkc"),
	// 		SystemData: &armiotfirmwaredefense.SystemData{
	// 			CreatedBy: to.Ptr("nqisshvdzqcxzbujvacin"),
	// 			CreatedByType: to.Ptr(armiotfirmwaredefense.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-06-13T15:22:45.940Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("of"),
	// 			LastModifiedByType: to.Ptr(armiotfirmwaredefense.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-06-13T15:22:45.940Z"); return t}()),
	// 		},
	// 	},
	// }
}
