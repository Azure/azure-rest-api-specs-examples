package armservicefabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ServicePutOperation_example_min.json
func ExampleServicesClient_BeginCreateOrUpdate_putAServiceWithMinimumParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServicesClient().BeginCreateOrUpdate(ctx, "resRg", "myCluster", "myApp", "myService", armservicefabric.ServiceResource{
		Tags: map[string]*string{},
		Properties: &armservicefabric.StatelessServiceProperties{
			PartitionDescription: &armservicefabric.SingletonPartitionSchemeDescription{
				PartitionScheme: to.Ptr(armservicefabric.PartitionSchemeSingleton),
			},
			ServiceKind:     to.Ptr(armservicefabric.ServiceKindStateless),
			ServiceTypeName: to.Ptr("myServiceType"),
			InstanceCount:   to.Ptr[int32](1),
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
