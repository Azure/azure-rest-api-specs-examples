package armdeviceprovisioningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
)

// Generated from example definition: 2025-02-01-preview/DPSGetValidSku.json
func ExampleIotDpsResourceClient_NewListValidSKUsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceprovisioningservices.NewClientFactory("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIotDpsResourceClient().NewListValidSKUsPager("myFirstProvisioningService", "myResourceGroup", nil)
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
		// page = armdeviceprovisioningservices.IotDpsResourceClientListValidSKUsResponse{
		// 	IotDpsSKUDefinitionListResult: armdeviceprovisioningservices.IotDpsSKUDefinitionListResult{
		// 		Value: []*armdeviceprovisioningservices.IotDpsSKUDefinition{
		// 			{
		// 				Name: to.Ptr(armdeviceprovisioningservices.IotDpsSKUS1),
		// 			},
		// 		},
		// 	},
		// }
	}
}
