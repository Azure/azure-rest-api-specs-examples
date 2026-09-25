package armstorage_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2026-06-01/TableOperationPutOrPatchAcls.json
func ExampleTableClient_Update_tableOperationPutOrPatchAcls() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTableClient().Update(ctx, "res3376", "sto328", "table6185", armstorage.Table{
		TableProperties: &armstorage.TableProperties{
			SignedIdentifiers: []*armstorage.TableSignedIdentifier{
				{
					AccessPolicy: &armstorage.TableAccessPolicy{
						ExpiryTime: to.Ptr(time.Date(2022, time.March, 20, 8, 49, 37, 0, time.UTC)),
						Permission: to.Ptr("raud"),
						StartTime:  to.Ptr(time.Date(2022, time.March, 17, 8, 49, 37, 0, time.UTC)),
					},
					ID: to.Ptr("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI"),
				},
				{
					AccessPolicy: &armstorage.TableAccessPolicy{
						ExpiryTime: to.Ptr(time.Date(2022, time.March, 20, 8, 49, 37, 0, time.UTC)),
						Permission: to.Ptr("rad"),
						StartTime:  to.Ptr(time.Date(2022, time.March, 17, 8, 49, 37, 0, time.UTC)),
					},
					ID: to.Ptr("PTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODklMTI"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstorage.TableClientUpdateResponse{
	// 	Table: armstorage.Table{
	// 		Name: to.Ptr("table6185"),
	// 		Type: to.Ptr("Microsoft.Storage/storageAccounts/tableServices/tables"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res3376/providers/Microsoft.Storage/storageAccounts/sto328/tableServices/default/tables/table6185"),
	// 		TableProperties: &armstorage.TableProperties{
	// 			SignedIdentifiers: []*armstorage.TableSignedIdentifier{
	// 				{
	// 					AccessPolicy: &armstorage.TableAccessPolicy{
	// 						ExpiryTime: to.Ptr(time.Date(2022, time.March, 20, 8, 49, 37, 0, time.UTC)),
	// 						Permission: to.Ptr("raud"),
	// 						StartTime: to.Ptr(time.Date(2022, time.March, 17, 8, 49, 37, 0, time.UTC)),
	// 					},
	// 					ID: to.Ptr("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI"),
	// 				},
	// 				{
	// 					AccessPolicy: &armstorage.TableAccessPolicy{
	// 						ExpiryTime: to.Ptr(time.Date(2022, time.March, 20, 8, 49, 37, 0, time.UTC)),
	// 						Permission: to.Ptr("rad"),
	// 						StartTime: to.Ptr(time.Date(2022, time.March, 17, 8, 49, 37, 0, time.UTC)),
	// 					},
	// 					ID: to.Ptr("PTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODklMTI"),
	// 				},
	// 			},
	// 			TableName: to.Ptr("table6185"),
	// 		},
	// 	},
	// }
}
