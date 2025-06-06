package armlogz_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logz/armlogz"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SingleSignOnConfigurations_CreateOrUpdate.json
func ExampleSingleSignOnClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogz.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSingleSignOnClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myMonitor", "default", &armlogz.SingleSignOnClientBeginCreateOrUpdateOptions{Body: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SingleSignOnResource = armlogz.SingleSignOnResource{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Logz/monitors"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Logz/monitors/myMonitor/singleSignOnConfigurations/default"),
	// 	Properties: &armlogz.SingleSignOnProperties{
	// 		EnterpriseAppID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		ProvisioningState: to.Ptr(armlogz.ProvisioningStateSucceeded),
	// 		SingleSignOnState: to.Ptr(armlogz.SingleSignOnStatesEnable),
	// 		SingleSignOnURL: to.Ptr("https://www.logz.io/IAmSomeHash"),
	// 	},
	// }
}
