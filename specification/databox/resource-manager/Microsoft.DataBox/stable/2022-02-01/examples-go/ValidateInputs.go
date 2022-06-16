package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/databox/resource-manager/Microsoft.DataBox/stable/2022-02-01/examples/ValidateInputs.json
func ExampleServiceClient_ValidateInputs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatabox.NewServiceClient("fa68082f-8ff7-4a25-95c7-ce9da541242f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ValidateInputs(ctx,
		"westus",
		&armdatabox.CreateJobValidations{
			IndividualRequestDetails: []armdatabox.ValidationInputRequestClassification{
				&armdatabox.DataTransferDetailsValidationRequest{
					ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidateDataTransferDetails),
					DataImportDetails: []*armdatabox.DataImportDetails{
						{
							AccountDetails: &armdatabox.StorageAccountDetails{
								DataAccountType:  to.Ptr(armdatabox.DataAccountTypeStorageAccount),
								StorageAccountID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/databoxbvt/providers/Microsoft.Storage/storageAccounts/databoxbvttestaccount"),
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
						City:            to.Ptr("San Francisco"),
						CompanyName:     to.Ptr("Microsoft"),
						Country:         to.Ptr("US"),
						PostalCode:      to.Ptr("94107"),
						StateOrProvince: to.Ptr("CA"),
						StreetAddress1:  to.Ptr("16 TOWNSEND ST"),
						StreetAddress2:  to.Ptr("Unit 1"),
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
					Country:        to.Ptr("US"),
					DeviceType:     to.Ptr(armdatabox.SKUNameDataBox),
					Location:       to.Ptr("westus"),
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
			ValidationCategory: to.Ptr("JobCreationValidation"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
