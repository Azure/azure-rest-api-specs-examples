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
	res, err := client.ValidateInputsByResourceGroup(ctx,
		"<resource-group-name>",
		"<location>",
		&armdatabox.CreateJobValidations{
			IndividualRequestDetails: []armdatabox.ValidationInputRequestClassification{
				&armdatabox.DataTransferDetailsValidationRequest{
					ValidationType: armdatabox.ValidationInputDiscriminatorValidateDataTransferDetails.ToPtr(),
					DataImportDetails: []*armdatabox.DataImportDetails{
						{
							AccountDetails: &armdatabox.StorageAccountDetails{
								DataAccountType:  armdatabox.DataAccountTypeStorageAccount.ToPtr(),
								StorageAccountID: to.StringPtr("<storage-account-id>"),
							},
						}},
					DeviceType:   armdatabox.SKUNameDataBox.ToPtr(),
					TransferType: armdatabox.TransferTypeImportToAzure.ToPtr(),
				},
				&armdatabox.ValidateAddress{
					ValidationType: armdatabox.ValidationInputDiscriminatorValidateAddress.ToPtr(),
					DeviceType:     armdatabox.SKUNameDataBox.ToPtr(),
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
					ValidationType: armdatabox.ValidationInputDiscriminatorValidateSubscriptionIsAllowedToCreateJob.ToPtr(),
				},
				&armdatabox.SKUAvailabilityValidationRequest{
					ValidationType: armdatabox.ValidationInputDiscriminatorValidateSKUAvailability.ToPtr(),
					Country:        to.StringPtr("<country>"),
					DeviceType:     armdatabox.SKUNameDataBox.ToPtr(),
					Location:       to.StringPtr("<location>"),
					TransferType:   armdatabox.TransferTypeImportToAzure.ToPtr(),
				},
				&armdatabox.CreateOrderLimitForSubscriptionValidationRequest{
					ValidationType: armdatabox.ValidationInputDiscriminatorValidateCreateOrderLimit.ToPtr(),
					DeviceType:     armdatabox.SKUNameDataBox.ToPtr(),
				},
				&armdatabox.PreferencesValidationRequest{
					ValidationType: armdatabox.ValidationInputDiscriminatorValidatePreferences.ToPtr(),
					DeviceType:     armdatabox.SKUNameDataBox.ToPtr(),
					Preference: &armdatabox.Preferences{
						TransportPreferences: &armdatabox.TransportPreferences{
							PreferredShipmentType: armdatabox.TransportShipmentTypesMicrosoftManaged.ToPtr(),
						},
					},
				}},
			ValidationCategory: to.StringPtr("<validation-category>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ServiceClientValidateInputsByResourceGroupResult)
}
