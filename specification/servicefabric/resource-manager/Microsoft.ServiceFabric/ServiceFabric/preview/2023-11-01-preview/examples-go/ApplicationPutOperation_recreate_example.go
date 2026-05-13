package armservicefabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric/v3"
)

// Generated from example definition: 2023-11-01-preview/ApplicationPutOperation_recreate_example.json
func ExampleApplicationsClient_BeginCreateOrUpdate_putAnApplicationWithRecreateOption() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationsClient().BeginCreateOrUpdate(ctx, "resRg", "myCluster", "myApp", armservicefabric.ApplicationResource{
		Properties: &armservicefabric.ApplicationResourceProperties{
			Parameters: map[string]*string{
				"param1": to.Ptr("value1"),
			},
			TypeName:    to.Ptr("myAppType"),
			TypeVersion: to.Ptr("1.0"),
			UpgradePolicy: &armservicefabric.ApplicationUpgradePolicy{
				RecreateApplication: to.Ptr(true),
			},
		},
		Tags: map[string]*string{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
}
