package armservicefabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationPutOperation_recreate_example.json
func ExampleApplicationsClient_BeginCreateOrUpdate_putAnApplicationWithRecreateOption() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationsClient().BeginCreateOrUpdate(ctx, "resRg", "myCluster", "myApp", armservicefabric.ApplicationResource{
		Tags: map[string]*string{},
		Properties: &armservicefabric.ApplicationResourceProperties{
			Parameters: map[string]*string{
				"param1": to.Ptr("value1"),
			},
			TypeVersion: to.Ptr("1.0"),
			UpgradePolicy: &armservicefabric.ApplicationUpgradePolicy{
				RecreateApplication: to.Ptr(true),
			},
			TypeName: to.Ptr("myAppType"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
