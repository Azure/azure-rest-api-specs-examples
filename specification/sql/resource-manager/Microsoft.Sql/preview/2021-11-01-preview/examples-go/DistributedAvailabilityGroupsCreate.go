package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/DistributedAvailabilityGroupsCreate.json
func ExampleDistributedAvailabilityGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDistributedAvailabilityGroupsClient().BeginCreateOrUpdate(ctx, "testrg", "testcl", "dag", armsql.DistributedAvailabilityGroup{
		Properties: &armsql.DistributedAvailabilityGroupProperties{
			PrimaryAvailabilityGroupName:   to.Ptr("BoxLocalAg1"),
			SecondaryAvailabilityGroupName: to.Ptr("testcl"),
			SourceEndpoint:                 to.Ptr("TCP://SERVER:7022"),
			TargetDatabase:                 to.Ptr("testdb"),
		},
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
	// res.DistributedAvailabilityGroup = armsql.DistributedAvailabilityGroup{
	// 	Name: to.Ptr("dag"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/distributedAvailabilityGroups"),
	// 	ID: to.Ptr("/subscriptions/f2669dff-5f08-45dd-b857-b2a60b72cdc9/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testcl/distributedAvailabilityGroups/dag"),
	// 	Properties: &armsql.DistributedAvailabilityGroupProperties{
	// 		PrimaryAvailabilityGroupName: to.Ptr("BoxLocalAg1"),
	// 		SecondaryAvailabilityGroupName: to.Ptr("testcl"),
	// 		SourceEndpoint: to.Ptr("TCP://SERVER:7022"),
	// 		TargetDatabase: to.Ptr("testdb"),
	// 	},
	// }
}
