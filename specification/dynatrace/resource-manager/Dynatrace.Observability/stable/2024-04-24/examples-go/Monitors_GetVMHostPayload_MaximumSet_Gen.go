package armdynatrace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dynatrace/armdynatrace/v2"
)

// Generated from example definition: 2024-04-24/Monitors_GetVMHostPayload_MaximumSet_Gen.json
func ExampleMonitorsClient_GetVMHostPayload_monitorsGetVMHostPayloadMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdynatrace.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMonitorsClient().GetVMHostPayload(ctx, "myResourceGroup", "myMonitor", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdynatrace.MonitorsClientGetVMHostPayloadResponse{
	// 	VMExtensionPayload: &armdynatrace.VMExtensionPayload{
	// 		EnvironmentID: to.Ptr("abc123lsjlsfjlfjgd"),
	// 		IngestionKey: to.Ptr("abcd.efg"),
	// 	},
	// }
}
