package armconnectedcache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedcache/armconnectedcache"
)

// Generated from example definition: 2023-05-01-preview/EnterpriseCustomerOperations_Update_MaximumSet_Gen.json
func ExampleEnterpriseCustomerOperationsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedcache.NewClientFactory("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEnterpriseCustomerOperationsClient().Update(ctx, "rgConnectedCache", "MCCTPTest2", armconnectedcache.PatchResource{
		Tags: map[string]*string{
			"key1653": to.Ptr("nzjczrhclhkndesgy"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armconnectedcache.EnterpriseCustomerOperationsClientUpdateResponse{
	// 	EnterprisePreviewResource: &armconnectedcache.EnterprisePreviewResource{
	// 		Properties: &armconnectedcache.CacheNodeOldResponse{
	// 			ProvisioningState: to.Ptr(armconnectedcache.ProvisioningStateSucceeded),
	// 			StatusCode: to.Ptr("oldkroffqtkryqffpsi"),
	// 			StatusText: to.Ptr("bs"),
	// 			StatusDetails: to.Ptr("lhwvcz"),
	// 			Status: to.Ptr("jurqdrxfpowdz"),
	// 			Error: &armconnectedcache.ErrorDetail{
	// 				Code: to.Ptr("dkvgvtftpsjsbhlnapvihefxneoggs"),
	// 				Message: to.Ptr("okakgyfnmyob"),
	// 				Details: []*armconnectedcache.ErrorDetail{
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key4215": to.Ptr("zjbszvlzf"),
	// 		},
	// 		Location: to.Ptr("westus"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/rgConnectedCache/providers/Microsoft.ConnectedCache/enterpriseCustomers/MccRPTest2"),
	// 		Name: to.Ptr("MCCTPTest2"),
	// 		Type: to.Ptr("Microsoft.ConnectedCache/enterpriseCustomers"),
	// 		SystemData: &armconnectedcache.SystemData{
	// 			CreatedBy: to.Ptr("gambtqj"),
	// 			CreatedByType: to.Ptr(armconnectedcache.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.771Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("qomgaceiessgnuogz"),
	// 			LastModifiedByType: to.Ptr(armconnectedcache.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.771Z"); return t}()),
	// 		},
	// 	},
	// }
}
