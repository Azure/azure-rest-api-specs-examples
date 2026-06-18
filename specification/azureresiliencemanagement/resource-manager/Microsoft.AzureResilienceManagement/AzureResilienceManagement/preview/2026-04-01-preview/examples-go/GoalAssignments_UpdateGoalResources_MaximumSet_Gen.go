package armresiliencemanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resiliencemanagement/armresiliencemanagement"
)

// Generated from example definition: 2026-04-01-preview/GoalAssignments_UpdateGoalResources_MaximumSet_Gen.json
func ExampleGoalAssignmentsClient_BeginUpdateGoalResources() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresiliencemanagement.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGoalAssignmentsClient().BeginUpdateGoalResources(ctx, "sg1", "ga1", armresiliencemanagement.UpdateGoalResourceRequest{
		Resources: []*armresiliencemanagement.GoalResource{
			{
				Properties: &armresiliencemanagement.GoalResourceProperties{
					ResourceArmID:                     to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/MyResourceGroup/providers/Microsoft.Compute/virtualMachines/MyVirtualMachine"),
					HighAvailabilityGoalParticipation: to.Ptr(armresiliencemanagement.ExclusionStateExcluded),
					HighAvailabilityAttestationStatus: to.Ptr(armresiliencemanagement.AttestationStateManuallyAttested),
					DisasterRecoveryGoalParticipation: to.Ptr(armresiliencemanagement.ExclusionStateExcluded),
					DisasterRecoveryAttestationStatus: to.Ptr(armresiliencemanagement.AttestationStateManuallyAttested),
				},
			},
			{
				Properties: &armresiliencemanagement.GoalResourceProperties{
					ResourceArmID:                     to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/MyResourceGroup/providers/Microsoft.Compute/virtualMachines/MyVirtualMachine1"),
					HighAvailabilityGoalParticipation: to.Ptr(armresiliencemanagement.ExclusionStateExcluded),
					HighAvailabilityAttestationStatus: to.Ptr(armresiliencemanagement.AttestationStateManuallyAttested),
					DisasterRecoveryGoalParticipation: to.Ptr(armresiliencemanagement.ExclusionStateExcluded),
					DisasterRecoveryAttestationStatus: to.Ptr(armresiliencemanagement.AttestationStateManuallyAttested),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armresiliencemanagement.GoalAssignmentsClientUpdateGoalResourcesResponse{
	// }
}
