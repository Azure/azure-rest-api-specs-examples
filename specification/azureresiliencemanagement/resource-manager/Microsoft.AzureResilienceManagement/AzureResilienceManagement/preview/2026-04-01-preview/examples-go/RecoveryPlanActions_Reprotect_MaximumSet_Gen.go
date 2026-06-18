package armresiliencemanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resiliencemanagement/armresiliencemanagement"
)

// Generated from example definition: 2026-04-01-preview/RecoveryPlanActions_Reprotect_MaximumSet_Gen.json
func ExampleRecoveryPlanActionsClient_BeginReprotect() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresiliencemanagement.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRecoveryPlanActionsClient().BeginReprotect(ctx, "sampleServiceGroupName", "<operationID>", "samplePlanName", &armresiliencemanagement.RecoveryPlanActionsClientBeginReprotectOptions{
		Body: &armresiliencemanagement.ReprotectRequest{
			ReprotectRequestProperties: &armresiliencemanagement.ReprotectRequestProperties{
				SelectedResourceIDs: []*string{
					to.Ptr("/providers/Microsoft.Management/serviceGroups/sampleServiceGroupName/providers/Microsoft.AzureResilienceManagement/recoveryPlans/samplePlanName/recoveryResources/12345678-9012-3456-7890-123456789012"),
				},
			},
		}})
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
	// res = armresiliencemanagement.RecoveryPlanActionsClientReprotectResponse{
	// 	RecoveryPlanActionBaseResponse: armresiliencemanagement.RecoveryPlanActionBaseResponse{
	// 		JobID: to.Ptr("11111111-1111-1111-1111-000000000012"),
	// 	},
	// }
}
