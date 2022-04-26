Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorageimportexport%2Farmstorageimportexport%2Fv0.4.0/sdk/resourcemanager/storageimportexport/armstorageimportexport/README.md) on how to add the SDK to your project and authenticate.

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
		return
	}
	ctx := context.Background()
	client, err := armstorageimportexport.NewJobsClient("<subscription-id>",
		nil, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Create(ctx,
		"<job-name>",
		"<resource-group-name>",
		armstorageimportexport.PutJobParameters{
			Location: to.Ptr("<location>"),
			Properties: &armstorageimportexport.JobDetails{
				BackupDriveManifest: to.Ptr(true),
				DiagnosticsPath:     to.Ptr("<diagnostics-path>"),
				Export: &armstorageimportexport.Export{
					BlobList: &armstorageimportexport.ExportBlobList{
						BlobPathPrefix: []*string{
							to.Ptr("/")},
					},
				},
				JobType:  to.Ptr("<job-type>"),
				LogLevel: to.Ptr("<log-level>"),
				ReturnAddress: &armstorageimportexport.ReturnAddress{
					City:            to.Ptr("<city>"),
					CountryOrRegion: to.Ptr("<country-or-region>"),
					Email:           to.Ptr("<email>"),
					Phone:           to.Ptr("<phone>"),
					PostalCode:      to.Ptr("<postal-code>"),
					RecipientName:   to.Ptr("<recipient-name>"),
					StateOrProvince: to.Ptr("<state-or-province>"),
					StreetAddress1:  to.Ptr("<street-address1>"),
					StreetAddress2:  to.Ptr("<street-address2>"),
				},
				ReturnShipping: &armstorageimportexport.ReturnShipping{
					CarrierAccountNumber: to.Ptr("<carrier-account-number>"),
					CarrierName:          to.Ptr("<carrier-name>"),
				},
				StorageAccountID: to.Ptr("<storage-account-id>"),
			},
		},
		&armstorageimportexport.JobsClientCreateOptions{ClientTenantID: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
