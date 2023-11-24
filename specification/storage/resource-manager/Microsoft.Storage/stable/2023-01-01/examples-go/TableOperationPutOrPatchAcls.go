package armstorage_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0baf811c3c76c87b3c127d098519bd97141222dd/specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/TableOperationPutOrPatchAcls.json
func ExampleTableClient_Create_tableOperationPutOrPatchAcls() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTableClient().Create(ctx, "res3376", "sto328", "table6185", &armstorage.TableClientCreateOptions{Parameters: &armstorage.Table{
		TableProperties: &armstorage.TableProperties{
			SignedIdentifiers: []*armstorage.TableSignedIdentifier{
				{
					AccessPolicy: &armstorage.TableAccessPolicy{
						ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-20T08:49:37.000Z"); return t }()),
						Permission: to.Ptr("raud"),
						StartTime:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-17T08:49:37.000Z"); return t }()),
					},
					ID: to.Ptr("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI"),
				},
				{
					AccessPolicy: &armstorage.TableAccessPolicy{
						ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-20T08:49:37.000Z"); return t }()),
						Permission: to.Ptr("rad"),
						StartTime:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-17T08:49:37.000Z"); return t }()),
					},
					ID: to.Ptr("PTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODklMTI"),
				}},
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Table = armstorage.Table{
	// 	Name: to.Ptr("table6185"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/tableServices/tables"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res3376/providers/Microsoft.Storage/storageAccounts/sto328/tableServices/default/tables/table6185"),
	// 	TableProperties: &armstorage.TableProperties{
	// 		SignedIdentifiers: []*armstorage.TableSignedIdentifier{
	// 			{
	// 				AccessPolicy: &armstorage.TableAccessPolicy{
	// 					ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-20T08:49:37.000Z"); return t}()),
	// 					Permission: to.Ptr("raud"),
	// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-17T08:49:37.000Z"); return t}()),
	// 				},
	// 				ID: to.Ptr("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI"),
	// 			},
	// 			{
	// 				AccessPolicy: &armstorage.TableAccessPolicy{
	// 					ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-20T08:49:37.000Z"); return t}()),
	// 					Permission: to.Ptr("rad"),
	// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-17T08:49:37.000Z"); return t}()),
	// 				},
	// 				ID: to.Ptr("PTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODklMTI"),
	// 		}},
	// 		TableName: to.Ptr("table6185"),
	// 	},
	// }
}
