Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatabox%2Farmdatabox%2Fv1.0.0/sdk/resourcemanager/databox/armdatabox/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/databox/resource-manager/Microsoft.DataBox/stable/2022-02-01/examples/JobsCreate.json
func ExampleJobsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatabox.NewJobsClient("fa68082f-8ff7-4a25-95c7-ce9da541242f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"SdkRg5154",
		"SdkJob952",
		armdatabox.JobResource{
			Location: to.Ptr("westus"),
			SKU: &armdatabox.SKU{
				Name: to.Ptr(armdatabox.SKUNameDataBox),
			},
			Properties: &armdatabox.JobProperties{
				TransferType: to.Ptr(armdatabox.TransferTypeImportToAzure),
				Details: &armdatabox.JobDetails{
					ContactDetails: &armdatabox.ContactDetails{
						ContactName: to.Ptr("Public SDK Test"),
						EmailList: []*string{
							to.Ptr("testing@microsoft.com")},
						Phone:          to.Ptr("1234567890"),
						PhoneExtension: to.Ptr("1234"),
					},
					DataImportDetails: []*armdatabox.DataImportDetails{
						{
							AccountDetails: &armdatabox.StorageAccountDetails{
								DataAccountType:  to.Ptr(armdatabox.DataAccountTypeStorageAccount),
								StorageAccountID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/databoxbvt/providers/Microsoft.Storage/storageAccounts/databoxbvttestaccount"),
							},
						}},
					JobDetailsType: to.Ptr(armdatabox.ClassDiscriminatorDataBox),
					ShippingAddress: &armdatabox.ShippingAddress{
						AddressType:     to.Ptr(armdatabox.AddressTypeCommercial),
						City:            to.Ptr("San Francisco"),
						CompanyName:     to.Ptr("Microsoft"),
						Country:         to.Ptr("US"),
						PostalCode:      to.Ptr("94107"),
						StateOrProvince: to.Ptr("CA"),
						StreetAddress1:  to.Ptr("16 TOWNSEND ST"),
						StreetAddress2:  to.Ptr("Unit 1"),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
