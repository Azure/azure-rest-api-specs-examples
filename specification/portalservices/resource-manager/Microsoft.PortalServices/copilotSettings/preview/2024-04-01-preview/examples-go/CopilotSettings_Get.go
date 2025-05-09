package armportalservicescopilot_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portalservicescopilot/armportalservicescopilot"
)

// Generated from example definition: 2024-04-01-preview/CopilotSettings_Get.json
func ExampleCopilotSettingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armportalservicescopilot.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCopilotSettingsClient().Get(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armportalservicescopilot.CopilotSettingsClientGetResponse{
	// 	CopilotSettingsResource: &armportalservicescopilot.CopilotSettingsResource{
	// 		Properties: &armportalservicescopilot.CopilotSettingsProperties{
	// 			AccessControlEnabled: to.Ptr(true),
	// 		},
	// 		ID: to.Ptr("/providers/Microsoft.Portal/copilotSettings/default"),
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.PortalServices/copilotSettings"),
	// 	},
	// }
}
