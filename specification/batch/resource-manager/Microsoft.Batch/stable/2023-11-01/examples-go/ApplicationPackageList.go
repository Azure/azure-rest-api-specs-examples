package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/90115af9fda46f323e5c42c274f2b376108d1d47/specification/batch/resource-manager/Microsoft.Batch/stable/2023-11-01/examples/ApplicationPackageList.json
func ExampleApplicationPackageClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApplicationPackageClient().NewListPager("default-azurebatch-japaneast", "sampleacct", "app1", &armbatch.ApplicationPackageClientListOptions{Maxresults: nil})
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
		// page.ListApplicationPackagesResult = armbatch.ListApplicationPackagesResult{
		// 	Value: []*armbatch.ApplicationPackage{
		// 		{
		// 			Name: to.Ptr("1.0"),
		// 			Type: to.Ptr("Microsoft.Batch/batchAccounts/applications/versions"),
		// 			Etag: to.Ptr("W/\"0x8D64FF0B9F47F67\""),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/applications/app1/versions/1.0"),
		// 			Properties: &armbatch.ApplicationPackageProperties{
		// 				State: to.Ptr(armbatch.PackageStatePending),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("2.0"),
		// 			Type: to.Ptr("Microsoft.Batch/batchAccounts/applications/versions"),
		// 			Etag: to.Ptr("W/\"0x8D64FF0B9F47F67\""),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/applications/app1/versions/2.0"),
		// 			Properties: &armbatch.ApplicationPackageProperties{
		// 				Format: to.Ptr("zip"),
		// 				LastActivationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-27T18:48:09.933Z"); return t}()),
		// 				State: to.Ptr(armbatch.PackageStateActive),
		// 			},
		// 	}},
		// }
	}
}
