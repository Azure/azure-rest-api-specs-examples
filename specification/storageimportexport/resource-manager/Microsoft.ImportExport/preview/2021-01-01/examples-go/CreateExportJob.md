Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorageimportexport%2Farmstorageimportexport%2Fv0.2.1/sdk/resourcemanager/storageimportexport/armstorageimportexport/README.md) on how to add the SDK to your project and authenticate.

```go
package armstorageimportexport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storageimportexport/armstorageimportexport"
)

// x-ms-original-file: specification/storageimportexport/resource-manager/Microsoft.ImportExport/preview/2021-01-01/examples/CreateExportJob.json
func ExampleJobsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorageimportexport.NewJobsClient("<subscription-id>",
		nil, cred, nil)
	res, err := client.Create(ctx,
		"<job-name>",
		"<resource-group-name>",
		armstorageimportexport.PutJobParameters{
			Location: to.StringPtr("<location>"),
			Properties: &armstorageimportexport.JobDetails{
				BackupDriveManifest: to.BoolPtr(true),
				DiagnosticsPath:     to.StringPtr("<diagnostics-path>"),
				Export: &armstorageimportexport.Export{
					BlobList: &armstorageimportexport.ExportBlobList{
						BlobPathPrefix: []*string{
							to.StringPtr("/")},
					},
				},
				JobType:  to.StringPtr("<job-type>"),
				LogLevel: to.StringPtr("<log-level>"),
				ReturnAddress: &armstorageimportexport.ReturnAddress{
					City:            to.StringPtr("<city>"),
					CountryOrRegion: to.StringPtr("<country-or-region>"),
					Email:           to.StringPtr("<email>"),
					Phone:           to.StringPtr("<phone>"),
					PostalCode:      to.StringPtr("<postal-code>"),
					RecipientName:   to.StringPtr("<recipient-name>"),
					StateOrProvince: to.StringPtr("<state-or-province>"),
					StreetAddress1:  to.StringPtr("<street-address1>"),
					StreetAddress2:  to.StringPtr("<street-address2>"),
				},
				ReturnShipping: &armstorageimportexport.ReturnShipping{
					CarrierAccountNumber: to.StringPtr("<carrier-account-number>"),
					CarrierName:          to.StringPtr("<carrier-name>"),
				},
				StorageAccountID: to.StringPtr("<storage-account-id>"),
			},
		},
		&armstorageimportexport.JobsClientCreateOptions{ClientTenantID: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.JobsClientCreateResult)
}
```
