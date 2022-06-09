```go
package armstorageimportexport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storageimportexport/armstorageimportexport"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storageimportexport/resource-manager/Microsoft.ImportExport/preview/2021-01-01/examples/CreateExportJob.json
func ExampleJobsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorageimportexport.NewJobsClient("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		nil, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"myExportJob",
		"myResourceGroup",
		armstorageimportexport.PutJobParameters{
			Location: to.Ptr("West US"),
			Properties: &armstorageimportexport.JobDetails{
				BackupDriveManifest: to.Ptr(true),
				DiagnosticsPath:     to.Ptr("waimportexport"),
				Export: &armstorageimportexport.Export{
					BlobList: &armstorageimportexport.ExportBlobList{
						BlobPathPrefix: []*string{
							to.Ptr("/")},
					},
				},
				JobType:  to.Ptr("Export"),
				LogLevel: to.Ptr("Verbose"),
				ReturnAddress: &armstorageimportexport.ReturnAddress{
					City:            to.Ptr("Redmond"),
					CountryOrRegion: to.Ptr("USA"),
					Email:           to.Ptr("Test@contoso.com"),
					Phone:           to.Ptr("4250000000"),
					PostalCode:      to.Ptr("98007"),
					RecipientName:   to.Ptr("Test"),
					StateOrProvince: to.Ptr("wa"),
					StreetAddress1:  to.Ptr("Street1"),
					StreetAddress2:  to.Ptr("street2"),
				},
				ReturnShipping: &armstorageimportexport.ReturnShipping{
					CarrierAccountNumber: to.Ptr("989ffff"),
					CarrierName:          to.Ptr("FedEx"),
				},
				StorageAccountID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.ClassicStorage/storageAccounts/test"),
			},
		},
		&armstorageimportexport.JobsClientCreateOptions{ClientTenantID: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorageimportexport%2Farmstorageimportexport%2Fv0.5.0/sdk/resourcemanager/storageimportexport/armstorageimportexport/README.md) on how to add the SDK to your project and authenticate.
