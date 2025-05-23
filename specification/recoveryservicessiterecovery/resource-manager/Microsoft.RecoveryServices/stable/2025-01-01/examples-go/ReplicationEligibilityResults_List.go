package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationEligibilityResults_List.json
func ExampleReplicationEligibilityResultsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReplicationEligibilityResultsClient().List(ctx, "testRg1", "testVm2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReplicationEligibilityResultsCollection = armrecoveryservicessiterecovery.ReplicationEligibilityResultsCollection{
	// 	Value: []*armrecoveryservicessiterecovery.ReplicationEligibilityResults{
	// 		{
	// 			Name: to.Ptr("default"),
	// 			Type: to.Ptr("Microsoft.RecoveryServices/replicationEligibilityResults"),
	// 			ID: to.Ptr("/subscriptions/d90d145a-4cdd-45a3-b2c4-971d69775278/resourceGroups/testRg1/providers/Microsoft.Compute/virtualMachines/testVm2/providers/Microsoft.RecoveryServices/replicationEligibilityResults/default"),
	// 			Properties: &armrecoveryservicessiterecovery.ReplicationEligibilityResultsProperties{
	// 				ClientRequestID: to.Ptr("a62c81df-e26e-47ea-ab4b-f1fcc1e5b135"),
	// 				Errors: []*armrecoveryservicessiterecovery.ReplicationEligibilityResultsErrorInfo{
	// 					{
	// 						Code: to.Ptr("AzureVmIsNotInDesiredProvisioningState"),
	// 						Message: to.Ptr("Azure virtual machine with Id (/subscriptions/d90d145a-4cdd-45a3-b2c4-971d69775278/resourceGroups/testRg1/providers/Microsoft.Compute/virtualMachines/testVm2) is with provisioning state 'failed'. To enable replication, VM's provisioning state should be 'succeeded'."),
	// 						PossibleCauses: to.Ptr("Virtual machine is not in desired state."),
	// 						RecommendedAction: to.Ptr("\n      Ensure that the VM's provisioning state is 'succeeded'.\n      Refer to https://aka.ms/a2a-vm-state-issues to troubleshoot VM provisioning state issues.\n    "),
	// 						Status: to.Ptr("Error"),
	// 					},
	// 					{
	// 						Code: to.Ptr("AzureVmIsNotInDesiredPowerState"),
	// 						Message: to.Ptr("Azure virtual machine with Id (/subscriptions/d90d145a-4cdd-45a3-b2c4-971d69775278/resourceGroups/testRg1/providers/Microsoft.Compute/virtualMachines/testVm2) is with power status 'deallocated'. To enable replication, VM's power status should be 'running'."),
	// 						PossibleCauses: to.Ptr("Virtual machine is not in desired state."),
	// 						RecommendedAction: to.Ptr("\n      Ensure that the VM's power status is 'running'.\n      You can check the power status in 'VM > Settings > Properties > Status' in Azure portal.\n    "),
	// 						Status: to.Ptr("Error"),
	// 				}},
	// 			},
	// 	}},
	// }
}
