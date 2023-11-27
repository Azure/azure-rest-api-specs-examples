package armdatashare_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/Shares_ListSynchronizationDetails.json
func ExampleSharesClient_NewListSynchronizationDetailsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSharesClient().NewListSynchronizationDetailsPager("SampleResourceGroup", "Account1", "Share1", armdatashare.ShareSynchronization{
		SynchronizationID: to.Ptr("7d0536a6-3fa5-43de-b152-3d07c4f6b2bb"),
	}, &armdatashare.SharesClientListSynchronizationDetailsOptions{SkipToken: nil,
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
		// page.SynchronizationDetailsList = armdatashare.SynchronizationDetailsList{
		// 	Value: []*armdatashare.SynchronizationDetails{
		// 		{
		// 			Name: to.Ptr("dataset1"),
		// 			DataSetID: to.Ptr("7d0536a6-3fa5-43de-b152-3d07c4f6b2bb"),
		// 			DataSetType: to.Ptr(armdatashare.DataSetTypeBlob),
		// 			DurationMs: to.Ptr[int32](2000),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-14T04:47:52.961Z"); return t}()),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-14T04:47:52.961Z"); return t}()),
		// 			Status: to.Ptr("Completed"),
		// 	}},
		// }
	}
}
