package armdatashare_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/ShareSubscriptions_CancelSynchronization.json
func ExampleShareSubscriptionsClient_BeginCancelSynchronization() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewShareSubscriptionsClient().BeginCancelSynchronization(ctx, "SampleResourceGroup", "Account1", "ShareSubscription1", armdatashare.ShareSubscriptionSynchronization{
		SynchronizationID: to.Ptr("7d0536a6-3fa5-43de-b152-3d07c4f6b2bb"),
	}, nil)
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
	// res.ShareSubscriptionSynchronization = armdatashare.ShareSubscriptionSynchronization{
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-06-30T02:37:48.497Z"); return t}()),
	// 	Status: to.Ptr("Cancelled"),
	// 	SynchronizationID: to.Ptr("343c4772-ad68-41aa-91b9-bab1c92f9c27"),
	// }
}
