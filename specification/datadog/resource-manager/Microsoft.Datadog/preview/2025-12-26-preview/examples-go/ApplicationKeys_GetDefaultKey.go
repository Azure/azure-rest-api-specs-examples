package armdatadog_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog/v2"
)

// Generated from example definition: 2025-12-26-preview/ApplicationKeys_GetDefaultKey.json
func ExampleMonitorsClient_GetDefaultApplicationKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMonitorsClient().GetDefaultApplicationKey(ctx, "myResourceGroup", "myMonitor", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatadog.MonitorsClientGetDefaultApplicationKeyResponse{
	// 	ApplicationKey: &armdatadog.ApplicationKey{
	// 		Name: to.Ptr("MyApplicationKey"),
	// 		CreatedBy: to.Ptr("john@example.com"),
	// 		Key: to.Ptr("aaaaaaaaaaaaaaaa1111111111111111"),
	// 	},
	// }
}
