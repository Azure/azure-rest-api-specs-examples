Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatabox%2Farmdatabox%2Fv0.2.0/sdk/resourcemanager/databox/armdatabox/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatabox_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/JobsCreate.json
func ExampleJobsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewJobsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<job-name>",
		armdatabox.JobResource{
			Location: to.StringPtr("<location>"),
			SKU: &armdatabox.SKU{
				Name: armdatabox.SKUNameDataBox.ToPtr(),
			},
			Properties: &armdatabox.JobProperties{
				TransferType: armdatabox.TransferTypeImportToAzure.ToPtr(),
				Details: &armdatabox.JobDetails{
					ContactDetails: &armdatabox.ContactDetails{
						ContactName: to.StringPtr("<contact-name>"),
						EmailList: []*string{
							to.StringPtr("testing@microsoft.com")},
						Phone:          to.StringPtr("<phone>"),
						PhoneExtension: to.StringPtr("<phone-extension>"),
					},
					DataImportDetails: []*armdatabox.DataImportDetails{
						{
							AccountDetails: &armdatabox.StorageAccountDetails{
								DataAccountType:  armdatabox.DataAccountTypeStorageAccount.ToPtr(),
								StorageAccountID: to.StringPtr("<storage-account-id>"),
							},
						}},
					JobDetailsType: armdatabox.ClassDiscriminatorDataBox.ToPtr(),
					ShippingAddress: &armdatabox.ShippingAddress{
						AddressType:     armdatabox.AddressTypeCommercial.ToPtr(),
						City:            to.StringPtr("<city>"),
						CompanyName:     to.StringPtr("<company-name>"),
						Country:         to.StringPtr("<country>"),
						PostalCode:      to.StringPtr("<postal-code>"),
						StateOrProvince: to.StringPtr("<state-or-province>"),
						StreetAddress1:  to.StringPtr("<street-address1>"),
						StreetAddress2:  to.StringPtr("<street-address2>"),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.JobsClientCreateResult)
}
```
