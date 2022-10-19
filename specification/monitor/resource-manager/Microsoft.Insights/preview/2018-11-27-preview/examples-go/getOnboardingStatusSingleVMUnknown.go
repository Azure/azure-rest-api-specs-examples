package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/preview/2018-11-27-preview/examples/getOnboardingStatusSingleVMUnknown.json
func ExampleVMInsightsClient_GetOnboardingStatus_getStatusForAVmThatHasNotYetReportedData() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewVMInsightsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetOnboardingStatus(ctx, "subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourceGroups/vm-resource-group/providers/Microsoft.Compute/virtualMachines/ubuntu-vm", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
