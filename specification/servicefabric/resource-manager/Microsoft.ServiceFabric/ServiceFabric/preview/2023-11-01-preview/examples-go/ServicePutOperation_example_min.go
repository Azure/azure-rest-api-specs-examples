package armservicefabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric/v3"
)

// Generated from example definition: 2023-11-01-preview/ServicePutOperation_example_min.json
func ExampleServicesClient_BeginCreateOrUpdate_putAServiceWithMinimumParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServicesClient().BeginCreateOrUpdate(ctx, "resRg", "myCluster", "myApp", "myService", armservicefabric.ServiceResource{
		Properties: &armservicefabric.StatelessServiceProperties{
			InstanceCount: to.Ptr[int32](1),
			PartitionDescription: &armservicefabric.SingletonPartitionSchemeDescription{
				PartitionScheme: to.Ptr(armservicefabric.PartitionSchemeSingleton),
			},
			ServiceKind:     to.Ptr(armservicefabric.ServiceKindStateless),
			ServiceTypeName: to.Ptr("myServiceType"),
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
