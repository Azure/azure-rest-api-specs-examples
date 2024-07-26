package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/redis/resource-manager/Microsoft.Cache/stable/2024-03-01/examples/RedisCacheAccessPolicyAssignmentCreateUpdate.json
func ExampleAccessPolicyAssignmentClient_BeginCreateUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccessPolicyAssignmentClient().BeginCreateUpdate(ctx, "rg1", "cache1", "accessPolicyAssignmentName1", armredis.CacheAccessPolicyAssignment{
		Properties: &armredis.CacheAccessPolicyAssignmentProperties{
			AccessPolicyName: to.Ptr("accessPolicy1"),
			ObjectID:         to.Ptr("6497c918-11ad-41e7-1b0f-7c518a87d0b0"),
			ObjectIDAlias:    to.Ptr("TestAADAppRedis"),
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
	// res.CacheAccessPolicyAssignment = armredis.CacheAccessPolicyAssignment{
	// 	Name: to.Ptr("accessPolicyAssignmentName1"),
	// 	Type: to.Ptr("Microsoft.Cache/Redis/accessPolicyAssignments"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/redis/cache1/accessPolicyAssignments/accessPolicyAssignmentName1"),
	// 	Properties: &armredis.CacheAccessPolicyAssignmentProperties{
	// 		AccessPolicyName: to.Ptr("accessPolicy1"),
	// 		ObjectID: to.Ptr("6497c918-11ad-41e7-1b0f-7c518a87d0b0"),
	// 		ObjectIDAlias: to.Ptr("TestAADAppRedis"),
	// 		ProvisioningState: to.Ptr(armredis.AccessPolicyAssignmentProvisioningStateSucceeded),
	// 	},
	// }
}
