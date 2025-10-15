package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v4"
)

// Generated from example definition: 2025-07-01/ResourceGuardCRUD/ListUpdateProtectionPolicyRequests.json
func ExampleResourceGuardsClient_NewGetUpdateProtectionPolicyRequestsObjectsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("0b352192-dcac-4cc7-992e-a96190ccc68c", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewResourceGuardsClient().NewGetUpdateProtectionPolicyRequestsObjectsPager("SampleResourceGroup", "swaggerExample", nil)
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
		// page = armdataprotection.ResourceGuardsClientGetUpdateProtectionPolicyRequestsObjectsResponse{
		// 	DppBaseResourceList: armdataprotection.DppBaseResourceList{
		// 		Value: []*armdataprotection.DppBaseResource{
		// 			{
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.DataProtection/resourceGuards/updateProtectionPolicyRequests"),
		// 				ID: to.Ptr("/subscriptions/0b352192-dcac-4cc7-992e-a96190ccc68c/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/resourceGuards/swaggerExample/updateProtectionPolicyRequests/default"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
