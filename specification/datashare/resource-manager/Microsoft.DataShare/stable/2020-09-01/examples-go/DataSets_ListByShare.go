package armdatashare_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSets_ListByShare.json
func ExampleDataSetsClient_NewListBySharePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataSetsClient().NewListBySharePager("SampleResourceGroup", "Account1", "Share1", &armdatashare.DataSetsClientListByShareOptions{SkipToken: nil,
		Filter:  nil,
		Orderby: nil,
	})
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
		// page.DataSetList = armdatashare.DataSetList{
		// 	Value: []armdatashare.DataSetClassification{
		// 		&armdatashare.BlobDataSet{
		// 			Name: to.Ptr("Dataset1"),
		// 			Type: to.Ptr("Microsoft.DataShare/accounts/shares/dataSets"),
		// 			ID: to.Ptr("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.DataShare/accounts/Account1/shares/Share1/datasets/Dataset1"),
		// 			Kind: to.Ptr(armdatashare.DataSetKindBlob),
		// 			Properties: &armdatashare.BlobProperties{
		// 				ContainerName: to.Ptr("C1"),
		// 				FilePath: to.Ptr("file22"),
		// 				ResourceGroup: to.Ptr("SampleResourceGroup"),
		// 				StorageAccountName: to.Ptr("storage1"),
		// 				SubscriptionID: to.Ptr("433a8dfd-e5d5-4e77-ad86-90acdc75eb1a"),
		// 			},
		// 		},
		// 		&armdatashare.BlobDataSet{
		// 			Name: to.Ptr("Dataset1"),
		// 			Type: to.Ptr("Microsoft.DataShare/accounts/shares/dataSets"),
		// 			ID: to.Ptr("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.DataShare/accounts/Account1/shares/Share1/datasets/Dataset2"),
		// 			Kind: to.Ptr(armdatashare.DataSetKindBlob),
		// 			Properties: &armdatashare.BlobProperties{
		// 				ContainerName: to.Ptr("C1"),
		// 				FilePath: to.Ptr("file21"),
		// 				ResourceGroup: to.Ptr("SampleResourceGroup"),
		// 				StorageAccountName: to.Ptr("storage2"),
		// 				SubscriptionID: to.Ptr("433a8dfd-e5d5-4e77-ad86-90acdc75eb1a"),
		// 			},
		// 	}},
		// }
	}
}
