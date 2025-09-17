package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense/v2"
)

// Generated from example definition: 2025-08-02/PasswordHashes_ListByFirmware_MaximumSet_Gen.json
func ExamplePasswordHashesClient_NewListByFirmwarePager_passwordHashesListByFirmwareMaximumSetGenGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPasswordHashesClient().NewListByFirmwarePager("rgiotfirmwaredefense", "exampleWorkspaceName", "00000000-0000-0000-0000-000000000000", nil)
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
		// page = armiotfirmwaredefense.PasswordHashesClientListByFirmwareResponse{
		// 	PasswordHashResourceListResult: armiotfirmwaredefense.PasswordHashResourceListResult{
		// 		Value: []*armiotfirmwaredefense.PasswordHashResource{
		// 			{
		// 				Properties: &armiotfirmwaredefense.PasswordHash{
		// 					PasswordHashID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					FilePath: to.Ptr("/path/to/file/containing/hash"),
		// 					Salt: to.Ptr("okk8R3sJ"),
		// 					Hash: to.Ptr("nOWjEHqTyMwLHT7puB6VM1"),
		// 					Context: to.Ptr("FoundUserName:$1$okk8R3sJ$nOWjEHqTyMwLHT7puB6VM1:14073:0:99999:7:::"),
		// 					Username: to.Ptr("FoundUserName"),
		// 					Algorithm: to.Ptr("MD5"),
		// 					ProvisioningState: to.Ptr(armiotfirmwaredefense.ProvisioningStateSucceeded),
		// 				},
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/ResourceGroupName/providers/Microsoft.IoTFirmwareDefense/workspaces/WorkspaceName/firmwares/00000000-0000-0000-0000-000000000000/passwordHashes/00000000-0000-0000-0000-000000000000"),
		// 				Name: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				Type: to.Ptr("Microsoft.IoTFirmwareDefense/workspaces/firmwares/passwordHashes"),
		// 				SystemData: &armiotfirmwaredefense.SystemData{
		// 					CreatedBy: to.Ptr("UserName"),
		// 					CreatedByType: to.Ptr(armiotfirmwaredefense.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-06-13T15:22:45.940Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("UserName"),
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
