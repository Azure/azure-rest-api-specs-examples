package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a29126ca8200a6c981a4e908e41fe55730df4cad/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationEligibilityResults_Get.json
func ExampleReplicationEligibilityResultsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReplicationEligibilityResultsClient().Get(ctx, "testRg1", "testVm1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReplicationEligibilityResults = armrecoveryservicessiterecovery.ReplicationEligibilityResults{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/replicationEligibilityResults"),
	// 	ID: to.Ptr("/subscriptions/d90d145a-4cdd-45a3-b2c4-971d69775278/resourceGroups/testRg1/providers/Microsoft.Compute/virtualMachines/testVm1/providers/Microsoft.RecoveryServices/replicationEligibilityResults/default"),
	// 	Properties: &armrecoveryservicessiterecovery.ReplicationEligibilityResultsProperties{
	// 		ClientRequestID: to.Ptr("7d72ade7-b9f7-4d9b-8a19-15d9728c6190"),
	// 		Errors: []*armrecoveryservicessiterecovery.ReplicationEligibilityResultsErrorInfo{
	// 			{
	// 				Code: to.Ptr("A2AOperatingSystemNotSupported"),
	// 				Message: to.Ptr("The A2A operation could not be completed as the virtual machine is running OS 'ubuntu' with version '18.04' which is not supported for replication."),
	// 				PossibleCauses: to.Ptr("OS version not supported."),
	// 				RecommendedAction: to.Ptr("The virtual machine is running unsupported Operating system. Refer to the documentation for supported OS versions - https://aka.ms/a2a-os-support-matrix."),
	// 				Status: to.Ptr("Error"),
	// 		}},
	// 	},
	// }
}
