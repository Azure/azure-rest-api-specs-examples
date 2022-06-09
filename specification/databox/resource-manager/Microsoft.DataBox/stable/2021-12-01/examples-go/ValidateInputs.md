```go
package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/databox/resource-manager/Microsoft.DataBox/stable/2021-12-01/examples/ValidateInputs.json
func ExampleServiceClient_ValidateInputs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdatabox.NewServiceClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.ValidateInputs(ctx,
		"<location>",
		&armdatabox.CreateJobValidations{
			IndividualRequestDetails: []armdatabox.ValidationInputRequestClassification{
				&armdatabox.DataTransferDetailsValidationRequest{
					ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidateDataTransferDetails),
					DataImportDetails: []*armdatabox.DataImportDetails{
						{
							AccountDetails: &armdatabox.StorageAccountDetails{
								DataAccountType:  to.Ptr(armdatabox.DataAccountTypeStorageAccount),
								StorageAccountID: to.Ptr("<storage-account-id>"),
							},
						}},
					DeviceType:   to.Ptr(armdatabox.SKUNameDataBox),
					TransferType: to.Ptr(armdatabox.TransferTypeImportToAzure),
				},
				&armdatabox.ValidateAddress{
					ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidateAddress),
					DeviceType:     to.Ptr(armdatabox.SKUNameDataBox),
					ShippingAddress: &armdatabox.ShippingAddress{
						AddressType:     to.Ptr(armdatabox.AddressTypeCommercial),
						City:            to.Ptr("<city>"),
						CompanyName:     to.Ptr("<company-name>"),
						Country:         to.Ptr("<country>"),
						PostalCode:      to.Ptr("<postal-code>"),
						StateOrProvince: to.Ptr("<state-or-province>"),
						StreetAddress1:  to.Ptr("<street-address1>"),
						StreetAddress2:  to.Ptr("<street-address2>"),
					},
					TransportPreferences: &armdatabox.TransportPreferences{
						PreferredShipmentType: to.Ptr(armdatabox.TransportShipmentTypesMicrosoftManaged),
					},
				},
				&armdatabox.SubscriptionIsAllowedToCreateJobValidationRequest{
					ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidateSubscriptionIsAllowedToCreateJob),
				},
				&armdatabox.SKUAvailabilityValidationRequest{
					ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidateSKUAvailability),
					Country:        to.Ptr("<country>"),
					DeviceType:     to.Ptr(armdatabox.SKUNameDataBox),
					Location:       to.Ptr("<location>"),
					TransferType:   to.Ptr(armdatabox.TransferTypeImportToAzure),
				},
				&armdatabox.CreateOrderLimitForSubscriptionValidationRequest{
					ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidateCreateOrderLimit),
					DeviceType:     to.Ptr(armdatabox.SKUNameDataBox),
				},
				&armdatabox.PreferencesValidationRequest{
					ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidatePreferences),
					DeviceType:     to.Ptr(armdatabox.SKUNameDataBox),
					Preference: &armdatabox.Preferences{
						TransportPreferences: &armdatabox.TransportPreferences{
							PreferredShipmentType: to.Ptr(armdatabox.TransportShipmentTypesMicrosoftManaged),
						},
					},
				}},
			ValidationCategory: to.Ptr("<validation-category>"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatabox%2Farmdatabox%2Fv0.4.0/sdk/resourcemanager/databox/armdatabox/README.md) on how to add the SDK to your project and authenticate.
