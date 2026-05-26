package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox/v3"
)

// Generated from example definition: 2025-07-01/ValidateInputs.json
func ExampleServiceClient_ValidateInputs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabox.NewClientFactory("YourSubscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceClient().ValidateInputs(ctx, "westus", &armdatabox.CreateJobValidations{
		IndividualRequestDetails: []armdatabox.ValidationInputRequestClassification{
			&armdatabox.DataTransferDetailsValidationRequest{
				DataImportDetails: []*armdatabox.DataImportDetails{
					{
						AccountDetails: &armdatabox.StorageAccountDetails{
							DataAccountType:  to.Ptr(armdatabox.DataAccountTypeStorageAccount),
							StorageAccountID: to.Ptr("/subscriptions/YourSubscriptionId/resourcegroups/YourResourceGroupName/providers/Microsoft.Storage/storageAccounts/YourStorageAccountName"),
						},
					},
				},
				DeviceType:     to.Ptr(armdatabox.SKUNameDataBox),
				Model:          to.Ptr(armdatabox.ModelNameDataBox),
				TransferType:   to.Ptr(armdatabox.TransferTypeImportToAzure),
				ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidateDataTransferDetails),
			},
			&armdatabox.ValidateAddress{
				DeviceType: to.Ptr(armdatabox.SKUNameDataBox),
				Model:      to.Ptr(armdatabox.ModelNameDataBox),
				ShippingAddress: &armdatabox.ShippingAddress{
					AddressType:     to.Ptr(armdatabox.AddressTypeCommercial),
					City:            to.Ptr("XXXX XXXX"),
					CompanyName:     to.Ptr("XXXX XXXX"),
					Country:         to.Ptr("XX"),
					PostalCode:      to.Ptr("00000"),
					StateOrProvince: to.Ptr("XX"),
					StreetAddress1:  to.Ptr("XXXX XXXX"),
					StreetAddress2:  to.Ptr("XXXX XXXX"),
				},
				TransportPreferences: &armdatabox.TransportPreferences{
					PreferredShipmentType: to.Ptr(armdatabox.TransportShipmentTypesMicrosoftManaged),
				},
				ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidateAddress),
			},
			&armdatabox.SubscriptionIsAllowedToCreateJobValidationRequest{
				ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidateSubscriptionIsAllowedToCreateJob),
			},
			&armdatabox.SKUAvailabilityValidationRequest{
				Country:        to.Ptr("XX"),
				DeviceType:     to.Ptr(armdatabox.SKUNameDataBox),
				Location:       to.Ptr("westus"),
				Model:          to.Ptr(armdatabox.ModelNameDataBox),
				TransferType:   to.Ptr(armdatabox.TransferTypeImportToAzure),
				ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidateSKUAvailability),
			},
			&armdatabox.CreateOrderLimitForSubscriptionValidationRequest{
				DeviceType:     to.Ptr(armdatabox.SKUNameDataBox),
				Model:          to.Ptr(armdatabox.ModelNameDataBox),
				ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidateCreateOrderLimit),
			},
			&armdatabox.PreferencesValidationRequest{
				DeviceType: to.Ptr(armdatabox.SKUNameDataBox),
				Model:      to.Ptr(armdatabox.ModelNameDataBox),
				Preference: &armdatabox.Preferences{
					TransportPreferences: &armdatabox.TransportPreferences{
						PreferredShipmentType: to.Ptr(armdatabox.TransportShipmentTypesMicrosoftManaged),
					},
				},
				ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidatePreferences),
			},
		},
		ValidationCategory: to.Ptr("JobCreationValidation"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatabox.ServiceClientValidateInputsResponse{
	// 	ValidationResponse: armdatabox.ValidationResponse{
	// 		Properties: &armdatabox.ValidationResponseProperties{
	// 			IndividualResponseDetails: []armdatabox.ValidationInputResponseClassification{
	// 				&armdatabox.DataTransferDetailsValidationResponseProperties{
	// 					Error: &armdatabox.CloudError{
	// 						AdditionalInfo: []*armdatabox.AdditionalErrorInfo{
	// 						},
	// 						Code: to.Ptr("Success"),
	// 						Target: to.Ptr("KeyEncryptionKey"),
	// 						Details: []*armdatabox.CloudError{
	// 						},
	// 					},
	// 					Status: to.Ptr(armdatabox.ValidationStatusValid),
	// 					ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidateDataTransferDetails),
	// 				},
	// 				&armdatabox.AddressValidationProperties{
	// 					AlternateAddresses: []*armdatabox.ShippingAddress{
	// 						{
	// 							AddressType: to.Ptr(armdatabox.AddressTypeNone),
	// 							City: to.Ptr("XXXX XXXX"),
	// 							Country: to.Ptr("XX"),
	// 							PostalCode: to.Ptr("00000"),
	// 							StateOrProvince: to.Ptr("XX"),
	// 							StreetAddress1: to.Ptr("XXXX XXXX"),
	// 							StreetAddress2: to.Ptr("XXXX XXXX"),
	// 							StreetAddress3: to.Ptr(""),
	// 						},
	// 					},
	// 					Error: &armdatabox.CloudError{
	// 						AdditionalInfo: []*armdatabox.AdditionalErrorInfo{
	// 						},
	// 						Code: to.Ptr("Success"),
	// 						Target: to.Ptr("KeyEncryptionKey"),
	// 						Details: []*armdatabox.CloudError{
	// 						},
	// 					},
	// 					ValidationStatus: to.Ptr(armdatabox.AddressValidationStatusValid),
	// 					ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidateAddress),
	// 				},
	// 				&armdatabox.SubscriptionIsAllowedToCreateJobValidationResponseProperties{
	// 					Error: &armdatabox.CloudError{
	// 						AdditionalInfo: []*armdatabox.AdditionalErrorInfo{
	// 						},
	// 						Code: to.Ptr("Success"),
	// 						Target: to.Ptr("KeyEncryptionKey"),
	// 						Details: []*armdatabox.CloudError{
	// 						},
	// 					},
	// 					Status: to.Ptr(armdatabox.ValidationStatusValid),
	// 					ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidateSubscriptionIsAllowedToCreateJob),
	// 				},
	// 				&armdatabox.SKUAvailabilityValidationResponseProperties{
	// 					Error: &armdatabox.CloudError{
	// 						AdditionalInfo: []*armdatabox.AdditionalErrorInfo{
	// 						},
	// 						Code: to.Ptr("Success"),
	// 						Target: to.Ptr("KeyEncryptionKey"),
	// 						Details: []*armdatabox.CloudError{
	// 						},
	// 					},
	// 					Status: to.Ptr(armdatabox.ValidationStatusValid),
	// 					ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidateSKUAvailability),
	// 				},
	// 				&armdatabox.CreateOrderLimitForSubscriptionValidationResponseProperties{
	// 					Error: &armdatabox.CloudError{
	// 						AdditionalInfo: []*armdatabox.AdditionalErrorInfo{
	// 						},
	// 						Code: to.Ptr("Success"),
	// 						Target: to.Ptr("KeyEncryptionKey"),
	// 						Details: []*armdatabox.CloudError{
	// 						},
	// 					},
	// 					Status: to.Ptr(armdatabox.ValidationStatusValid),
	// 					ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidateCreateOrderLimit),
	// 				},
	// 				&armdatabox.PreferencesValidationResponseProperties{
	// 					Error: &armdatabox.CloudError{
	// 						AdditionalInfo: []*armdatabox.AdditionalErrorInfo{
	// 						},
	// 						Code: to.Ptr("Success"),
	// 						Target: to.Ptr("KeyEncryptionKey"),
	// 						Details: []*armdatabox.CloudError{
	// 						},
	// 					},
	// 					Status: to.Ptr(armdatabox.ValidationStatusValid),
	// 					ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidatePreferences),
	// 				},
	// 			},
	// 			Status: to.Ptr(armdatabox.OverallValidationStatusAllValidToProceed),
	// 		},
	// 	},
	// }
}
