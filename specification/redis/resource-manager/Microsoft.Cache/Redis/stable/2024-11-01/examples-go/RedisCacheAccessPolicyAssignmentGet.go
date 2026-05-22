package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v4"
)

// Generated from example definition: 2024-11-01/RedisCacheAccessPolicyAssignmentGet.json
func ExampleAccessPolicyAssignmentClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccessPolicyAssignmentClient().Get(ctx, "rg1", "cache1", "accessPolicyAssignmentName1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armredis.AccessPolicyAssignmentClientGetResponse{
	// 	CacheAccessPolicyAssignment: &armredis.CacheAccessPolicyAssignment{
	// 		Name: to.Ptr("accessPolicyAssignmentName1"),
	// 		Type: to.Ptr("Microsoft.Cache/Redis/accessPolicyAssignments"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Cache/redis/cache1/accessPolicyAssignments/accessPolicyAssignmentName1"),
	// 		Properties: &armredis.CacheAccessPolicyAssignmentProperties{
	// 			AccessPolicyName: to.Ptr("accessPolicy1"),
	// 			ObjectID: to.Ptr("6497c918-11ad-41e7-1b0f-7c518a87d0b0"),
	// 			ObjectIDAlias: to.Ptr("TestAADAppRedis"),
	// 			ProvisioningState: to.Ptr(armredis.AccessPolicyAssignmentProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
