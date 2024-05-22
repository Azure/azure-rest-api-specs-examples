package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2023-08-01-preview/examples/DistributedAvailabilityGroupsCreateMin.json
func ExampleDistributedAvailabilityGroupsClient_BeginCreateOrUpdate_createADistributedAvailabilityGroupWithMinimalProperties() {
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
			Databases: []*armsql.DistributedAvailabilityGroupDatabase{
				{
					DatabaseName: to.Ptr("testdb"),
				}},
			InstanceAvailabilityGroupName: to.Ptr("testcl"),
			PartnerAvailabilityGroupName:  to.Ptr("BoxLocalAg1"),
			PartnerEndpoint:               to.Ptr("TCP://SERVER:7022"),
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
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testcl/distributedAvailabilityGroups/dag"),
	// 	Properties: &armsql.DistributedAvailabilityGroupProperties{
	// 		Databases: []*armsql.DistributedAvailabilityGroupDatabase{
	// 			{
	// 				DatabaseName: to.Ptr("testdb"),
	// 		}},
	// 		DistributedAvailabilityGroupName: to.Ptr("dag"),
	// 		FailoverMode: to.Ptr(armsql.FailoverModeTypeNone),
	// 		InstanceAvailabilityGroupName: to.Ptr("testcl"),
	// 		InstanceLinkRole: to.Ptr(armsql.LinkRoleSecondary),
	// 		PartnerAvailabilityGroupName: to.Ptr("BoxLocalAg1"),
	// 		PartnerEndpoint: to.Ptr("TCP://SERVER:7022"),
	// 		SeedingMode: to.Ptr(armsql.SeedingModeTypeAutomatic),
	// 	},
	// }
}
