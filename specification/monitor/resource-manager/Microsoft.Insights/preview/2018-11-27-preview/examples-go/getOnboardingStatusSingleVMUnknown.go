package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/preview/2018-11-27-preview/examples/getOnboardingStatusSingleVMUnknown.json
func ExampleVMInsightsClient_GetOnboardingStatus_getStatusForAVmThatHasNotYetReportedData() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVMInsightsClient().GetOnboardingStatus(ctx, "subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourceGroups/vm-resource-group/providers/Microsoft.Compute/virtualMachines/ubuntu-vm", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VMInsightsOnboardingStatus = armmonitor.VMInsightsOnboardingStatus{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Insights/vmInsightsOnboardingStatuses"),
	// 	ID: to.Ptr("/subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourceGroups/vm-resource-group/providers/Microsoft.Compute/virtualMachines/ubuntu-vm/providers/Microsoft.Insights/vmInsightsOnboardingStatuses/default"),
	// 	Properties: &armmonitor.VMInsightsOnboardingStatusProperties{
	// 		Data: []*armmonitor.DataContainer{
	// 		},
	// 		DataStatus: to.Ptr(armmonitor.DataStatusNotPresent),
	// 		OnboardingStatus: to.Ptr(armmonitor.OnboardingStatusUnknown),
	// 		ResourceID: to.Ptr("/subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourceGroups/vm-resource-group/providers/Microsoft.Compute/virtualMachines/ubuntu-vm"),
	// 	},
	// }
}
