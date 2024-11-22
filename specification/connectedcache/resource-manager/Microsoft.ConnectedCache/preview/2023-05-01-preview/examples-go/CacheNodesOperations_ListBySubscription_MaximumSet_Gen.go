package armconnectedcache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedcache/armconnectedcache"
)

// Generated from example definition: 2023-05-01-preview/CacheNodesOperations_ListBySubscription_MaximumSet_Gen.json
func ExampleCacheNodesOperationsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedcache.NewClientFactory("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCacheNodesOperationsClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armconnectedcache.CacheNodesOperationsClientListBySubscriptionResponse{
		// 	CacheNodePreviewResourceListResult: armconnectedcache.CacheNodePreviewResourceListResult{
		// 		Value: []*armconnectedcache.CacheNodePreviewResource{
		// 			{
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/DoTest/providers/Microsoft.ConnectedCache/cacheNodes/MccRPTest2"),
		// 				Name: to.Ptr("MCCTPTest2"),
		// 				Type: to.Ptr("Microsoft.ConnectedCache/cacheNodes"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armconnectedcache.CacheNodeOldResponse{
		// 					ProvisioningState: to.Ptr(armconnectedcache.ProvisioningStateSucceeded),
		// 					StatusCode: to.Ptr("oldkroffqtkryqffpsi"),
		// 					StatusText: to.Ptr("bs"),
		// 					StatusDetails: to.Ptr("lhwvcz"),
		// 					Status: to.Ptr("jurqdrxfpowdz"),
		// 					Error: &armconnectedcache.ErrorDetail{
		// 						Code: to.Ptr("dkvgvtftpsjsbhlnapvihefxneoggs"),
		// 						Message: to.Ptr("okakgyfnmyob"),
		// 						Details: []*armconnectedcache.ErrorDetail{
		// 						},
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"key4711": to.Ptr("zfrkvxymuupxgxrkkmbm"),
		// 				},
		// 				SystemData: &armconnectedcache.SystemData{
		// 					CreatedBy: to.Ptr("gambtqj"),
		// 					CreatedByType: to.Ptr(armconnectedcache.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.771Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("qomgaceiessgnuogz"),
		// 					LastModifiedByType: to.Ptr(armconnectedcache.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.771Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/DoTest/providers/Microsoft.ConnectedCache/cacheNodes/MccRPTest1"),
		// 				Name: to.Ptr("MCCTPTest1"),
		// 				Type: to.Ptr("Microsoft.ConnectedCache/cacheNodes"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armconnectedcache.CacheNodeOldResponse{
		// 					ProvisioningState: to.Ptr(armconnectedcache.ProvisioningStateSucceeded),
		// 					StatusCode: to.Ptr("gtqfmwgkfp"),
		// 					StatusText: to.Ptr("ztbgggsdeihkputd"),
		// 					StatusDetails: to.Ptr("odhpdelbuqrfkiaolqrrzpdaokctz"),
		// 					Status: to.Ptr("mobullauhrxnpyocr"),
		// 					Error: &armconnectedcache.ErrorDetail{
		// 						Code: to.Ptr("dkvgvtftpsjsbhlnapvihefxneoggs"),
		// 						Message: to.Ptr("okakgyfnmyob"),
		// 						Details: []*armconnectedcache.ErrorDetail{
		// 						},
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"key8585": to.Ptr("jgoeqypqcf"),
		// 				},
		// 				SystemData: &armconnectedcache.SystemData{
		// 					CreatedBy: to.Ptr("gambtqj"),
		// 					CreatedByType: to.Ptr(armconnectedcache.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.771Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("qomgaceiessgnuogz"),
		// 					LastModifiedByType: to.Ptr(armconnectedcache.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.771Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
