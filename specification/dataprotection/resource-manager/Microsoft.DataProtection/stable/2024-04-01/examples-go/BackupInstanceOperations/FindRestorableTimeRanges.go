package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/BackupInstanceOperations/FindRestorableTimeRanges.json
func ExampleRestorableTimeRangesClient_Find() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRestorableTimeRangesClient().Find(ctx, "Blob-Backup", "ZBlobBackupVaultBVTD3", "zblobbackuptestsa58", armdataprotection.AzureBackupFindRestorableTimeRangesRequest{
		EndTime:             to.Ptr("2021-02-24T00:35:17.6829685Z"),
		SourceDataStoreType: to.Ptr(armdataprotection.RestoreSourceDataStoreTypeOperationalStore),
		StartTime:           to.Ptr("2020-10-17T23:28:17.6829685Z"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AzureBackupFindRestorableTimeRangesResponseResource = armdataprotection.AzureBackupFindRestorableTimeRangesResponseResource{
	// 	Type: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances/findRestorableTimeRanges"),
	// 	ID: to.Ptr("zblobbackuptestsa58"),
	// 	Properties: &armdataprotection.AzureBackupFindRestorableTimeRangesResponse{
	// 		ObjectType: to.Ptr("AzureBackupFindRestorableTimeRangesResponse"),
	// 		RestorableTimeRanges: []*armdataprotection.RestorableTimeRange{
	// 			{
	// 				EndTime: to.Ptr("2021-02-24T00:35:17.0000000Z"),
	// 				ObjectType: to.Ptr("RestorableTimeRange"),
	// 				StartTime: to.Ptr("2021-02-23T18:33:51.6349708Z"),
	// 		}},
	// 	},
	// }
}
