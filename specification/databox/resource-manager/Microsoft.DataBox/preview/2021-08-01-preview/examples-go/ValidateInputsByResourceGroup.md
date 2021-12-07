Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatabox%2Farmdatabox%2Fv0.1.0/sdk/resourcemanager/databox/armdatabox/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/ValidateInputsByResourceGroup.json
func ExampleServiceClient_ValidateInputsByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewServiceClient("<subscription-id>", cred, nil)
	_, err = client.ValidateInputsByResourceGroup(ctx,
		"<resource-group-name>",
		"<location>",
		&armdatabox.CreateJobValidations{
			ValidationRequest: armdatabox.ValidationRequest{
				IndividualRequestDetails: []armdatabox.ValidationInputRequestClassification{
					&armdatabox.DataTransferDetailsValidationRequest{
						ValidationInputRequest: armdatabox.ValidationInputRequest{
							ValidationType: armdatabox.ValidationInputDiscriminatorValidateDataTransferDetails.ToPtr(),
						},
						DataImportDetails: []*armdatabox.DataImportDetails{
							{
								AccountDetails: &armdatabox.StorageAccountDetails{
									DataAccountDetails: armdatabox.DataAccountDetails{
										DataAccountType: armdatabox.DataAccountTypeStorageAccount.ToPtr(),
									},
									StorageAccountID: to.StringPtr("<storage-account-id>"),
								},
							}},
						DeviceType:   armdatabox.SKUNameDataBox.ToPtr(),
						TransferType: armdatabox.TransferTypeImportToAzure.ToPtr(),
					},
					&armdatabox.ValidateAddress{
						ValidationInputRequest: armdatabox.ValidationInputRequest{
							ValidationType: armdatabox.ValidationInputDiscriminatorValidateAddress.ToPtr(),
						},
						DeviceType: armdatabox.SKUNameDataBox.ToPtr(),
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
						TransportPreferences: &armdatabox.TransportPreferences{
							PreferredShipmentType: armdatabox.TransportShipmentTypesMicrosoftManaged.ToPtr(),
						},
					},
					&armdatabox.SubscriptionIsAllowedToCreateJobValidationRequest{
						ValidationInputRequest: armdatabox.ValidationInputRequest{
							ValidationType: armdatabox.ValidationInputDiscriminatorValidateSubscriptionIsAllowedToCreateJob.ToPtr(),
						},
					},
					&armdatabox.SKUAvailabilityValidationRequest{
						ValidationInputRequest: armdatabox.ValidationInputRequest{
							ValidationType: armdatabox.ValidationInputDiscriminatorValidateSKUAvailability.ToPtr(),
						},
						Country:      to.StringPtr("<country>"),
						DeviceType:   armdatabox.SKUNameDataBox.ToPtr(),
						Location:     to.StringPtr("<location>"),
						TransferType: armdatabox.TransferTypeImportToAzure.ToPtr(),
					},
					&armdatabox.CreateOrderLimitForSubscriptionValidationRequest{
						ValidationInputRequest: armdatabox.ValidationInputRequest{
							ValidationType: armdatabox.ValidationInputDiscriminatorValidateCreateOrderLimit.ToPtr(),
						},
						DeviceType: armdatabox.SKUNameDataBox.ToPtr(),
					},
					&armdatabox.PreferencesValidationRequest{
						ValidationInputRequest: armdatabox.ValidationInputRequest{
							ValidationType: armdatabox.ValidationInputDiscriminatorValidatePreferences.ToPtr(),
						},
						DeviceType: armdatabox.SKUNameDataBox.ToPtr(),
						Preference: &armdatabox.Preferences{
							TransportPreferences: &armdatabox.TransportPreferences{
								PreferredShipmentType: armdatabox.TransportShipmentTypesMicrosoftManaged.ToPtr(),
							},
						},
					}},
				ValidationCategory: to.StringPtr("<validation-category>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
