package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense/v2"
)

// Generated from example definition: 2025-04-01-preview/Firmwares_ListByWorkspace_MinimumSet_Gen.json
func ExampleFirmwaresClient_NewListByWorkspacePager_firmwaresListByWorkspaceMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("685C0C6F-9867-4B1C-A534-AA3A05B54BCE", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFirmwaresClient().NewListByWorkspacePager("rgworkspaces-firmwares", "A7", nil)
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
		// page = armiotfirmwaredefense.FirmwaresClientListByWorkspaceResponse{
		// 	FirmwareListResult: armiotfirmwaredefense.FirmwareListResult{
		// 		Value: []*armiotfirmwaredefense.Firmware{
		// 		},
		// 	},
		// }
	}
}
