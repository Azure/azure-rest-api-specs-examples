package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/EHAliasList.json
func ExampleDisasterRecoveryConfigsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDisasterRecoveryConfigsClient().NewListPager("exampleResourceGroup", "sdk-Namespace-8859", nil)
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
		// page.ArmDisasterRecoveryListResult = armeventhub.ArmDisasterRecoveryListResult{
		// 	Value: []*armeventhub.ArmDisasterRecovery{
		// 		{
		// 			Name: to.Ptr("sdk-DisasterRecovery-3814"),
		// 			Type: to.Ptr("Microsoft.EventHub/Namespaces/DisasterRecoveryConfig"),
		// 			ID: to.Ptr("/subscriptions/exampleSubscriptionId/resourceGroups/exampleResourceGroup/providers/Microsoft.EventHub/namespaces/sdk-Namespace-8859/disasterRecoveryConfig/sdk-DisasterRecovery-3814"),
		// 			Properties: &armeventhub.ArmDisasterRecoveryProperties{
		// 				PartnerNamespace: to.Ptr("sdk-Namespace-37"),
		// 				ProvisioningState: to.Ptr(armeventhub.ProvisioningStateDRSucceeded),
		// 				Role: to.Ptr(armeventhub.RoleDisasterRecoveryPrimary),
		// 			},
		// 	}},
		// }
	}
}
