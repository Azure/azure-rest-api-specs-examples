package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/SBAliasList.json
func ExampleDisasterRecoveryConfigsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDisasterRecoveryConfigsClient().NewListPager("ardsouzatestRG", "sdk-Namespace-8860", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ArmDisasterRecoveryListResult = armservicebus.ArmDisasterRecoveryListResult{
		// 	Value: []*armservicebus.ArmDisasterRecovery{
		// 		{
		// 			Name: to.Ptr("sdk-DisasterRecovery-3814"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces/DisasterRecoveryConfig"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ardsouzatestRG/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-8860/disasterRecoveryConfig/sdk-DisasterRecovery-3814"),
		// 			Properties: &armservicebus.ArmDisasterRecoveryProperties{
		// 				PartnerNamespace: to.Ptr("sdk-Namespace-37"),
		// 				ProvisioningState: to.Ptr(armservicebus.ProvisioningStateDRSucceeded),
		// 				Role: to.Ptr(armservicebus.RoleDisasterRecoveryPrimary),
		// 			},
		// 	}},
		// }
	}
}
