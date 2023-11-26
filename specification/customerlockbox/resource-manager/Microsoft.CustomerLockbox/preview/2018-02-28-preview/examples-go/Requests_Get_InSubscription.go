package armcustomerlockbox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerlockbox/armcustomerlockbox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/customerlockbox/resource-manager/Microsoft.CustomerLockbox/preview/2018-02-28-preview/examples/Requests_Get_InSubscription.json
func ExampleRequestsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerlockbox.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRequestsClient().Get(ctx, "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LockboxRequestResponse = armcustomerlockbox.LockboxRequestResponse{
	// 	Name: to.Ptr("LockboxRequest_xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
	// 	Type: to.Ptr("requests"),
	// 	ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/providers/Microsoft.CustomerLockbox/requests/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
	// 	Properties: &armcustomerlockbox.LockboxRequestResponseProperties{
	// 		CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-17T23:15:42.031Z"); return t}()),
	// 		Duration: to.Ptr("8"),
	// 		ExpirationDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-21T23:15:42.031Z"); return t}()),
	// 		Justification: to.Ptr("Microsoft Support Team is requesting access to your resource temporarily for troubleshooting."),
	// 		RequestID: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
	// 		ResourceIDs: to.Ptr("SUBSCRIPTIONS/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/RESOURCEGROUPS/DEMORG/PROVIDERS/MICROSOFT.COMPUTE/VIRTUALMACHINES/DEMO"),
	// 		ResourceType: to.Ptr("VMSS"),
	// 		Status: to.Ptr(armcustomerlockbox.StatusExpired),
	// 		SubscriptionID: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
	// 		SupportCaseURL: to.Ptr("https://portal.azure.com/#resource/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/providers/microsoft.support/supporttickets/12345"),
	// 		SupportRequest: to.Ptr("12345"),
	// 		Workitemsource: to.Ptr("Other"),
	// 	},
	// }
}
