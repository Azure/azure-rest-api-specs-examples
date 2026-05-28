package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: 2025-05-01-preview/disasterRecoveryConfigs/SBAliasGet.json
func ExampleDisasterRecoveryConfigsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDisasterRecoveryConfigsClient().Get(ctx, "ardsouzatestRG", "sdk-Namespace-8860", "sdk-DisasterRecovery-3814", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armservicebus.DisasterRecoveryConfigsClientGetResponse{
	// 	ArmDisasterRecovery: armservicebus.ArmDisasterRecovery{
	// 		Name: to.Ptr("sdk-DisasterRecovery-3814"),
	// 		Type: to.Ptr("Microsoft.ServiceBus/Namespaces/DisasterRecoveryConfig"),
	// 		ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ardsouzatestRG/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-37/disasterRecoveryConfig/sdk-DisasterRecovery-3814"),
	// 		Properties: &armservicebus.ArmDisasterRecoveryProperties{
	// 			PartnerNamespace: to.Ptr("sdk-Namespace-8860"),
	// 			PendingReplicationOperationsCount: to.Ptr[int64](0),
	// 			ProvisioningState: to.Ptr(armservicebus.ProvisioningStateDRAccepted),
	// 			Role: to.Ptr(armservicebus.RoleDisasterRecoverySecondary),
	// 		},
	// 	},
	// }
}
