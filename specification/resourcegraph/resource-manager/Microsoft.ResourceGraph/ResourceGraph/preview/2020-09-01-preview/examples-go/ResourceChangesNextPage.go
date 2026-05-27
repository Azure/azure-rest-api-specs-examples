package armresourcegraph_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
)

// Generated from example definition: 2020-09-01-preview/ResourceChangesNextPage.json
func ExampleClient_ResourceChanges_nextPageQuery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcegraph.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().ResourceChanges(ctx, armresourcegraph.ResourceChangesRequestParameters{
		SkipToken: to.Ptr("ew0KICAiJGlkIjogIjEiLA0KICAiRW5kVGltZSI6ICJcL0RhdGUoMTU1MDc0NT"),
		Top:       to.Ptr[int32](2),
		Interval: &armresourcegraph.ResourceChangesRequestParametersInterval{
			End:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-31T12:09:03.141Z"); return t }()),
			Start: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-30T12:09:03.141Z"); return t }()),
		},
		ResourceIDs: []*string{
			to.Ptr("/subscriptions/4d962866-1e3f-47f2-bd18-450c08f914c1/resourceGroups/MyResourceGroup/providers/Microsoft.Storage/storageAccounts/mystorageaccount"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armresourcegraph.ClientResourceChangesResponse{
	// 	ResourceChangeList: armresourcegraph.ResourceChangeList{
	// 		SkipToken: "kVGltZSI6ICJcL0RhdGUoMTU1MDc0NTew0KICAiJGlkIjogIjEiLA0KICAiRW5",
	// 		Changes: []*armresourcegraph.ResourceChangeData{
	// 			{
	// 				AfterSnapshot: &armresourcegraph.ResourceChangeDataAfterSnapshot{
	// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-31T05:12:32.087Z"); return t}()),
	// 				},
	// 				BeforeSnapshot: &armresourcegraph.ResourceChangeDataBeforeSnapshot{
	// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-31T03:43:08.629Z"); return t}()),
	// 				},
	// 				ChangeID: to.Ptr("55f458c4-f1c0-4963-bc6c-af275cd47702"),
	// 				ChangeType: to.Ptr(armresourcegraph.ChangeTypeUpdate),
	// 				ResourceID: to.Ptr("/subscriptions/4d962866-1e3f-47f2-bd18-450c08f914c1/resourceGroups/MyResourceGroup/providers/Microsoft.Storage/storageAccounts/mystorageaccount"),
	// 			},
	// 			{
	// 				AfterSnapshot: &armresourcegraph.ResourceChangeDataAfterSnapshot{
	// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-31T02:01:02.163Z"); return t}()),
	// 				},
	// 				BeforeSnapshot: &armresourcegraph.ResourceChangeDataBeforeSnapshot{
	// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-31T01:54:24.42Z"); return t}()),
	// 				},
	// 				ChangeID: to.Ptr("0495b929-b86d-46cc-a972-939145feed90"),
	// 				ChangeType: to.Ptr(armresourcegraph.ChangeTypeCreate),
	// 				ResourceID: to.Ptr("/subscriptions/4d962866-1e3f-47f2-bd18-450c08f914c1/resourceGroups/MyResourceGroup/providers/Microsoft.Storage/storageAccounts/mystorageaccount"),
	// 			},
	// 		},
	// 	},
	// }
}
