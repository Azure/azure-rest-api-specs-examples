package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databox/resource-manager/Microsoft.DataBox/stable/2022-02-01/examples/ValidateInputs.json
func ExampleServiceClient_ValidateInputs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabox.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceClient().ValidateInputs(ctx, "westus", &armdatabox.CreateJobValidations{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ValidationResponse = armdatabox.ValidationResponse{
	// 	Properties: &armdatabox.ValidationResponseProperties{
	// 		IndividualResponseDetails: []armdatabox.ValidationInputResponseClassification{
	// 			&armdatabox.DataTransferDetailsValidationResponseProperties{
	// 				Error: &armdatabox.CloudError{
	// 					AdditionalInfo: []*armdatabox.AdditionalErrorInfo{
	// 					},
	// 					Code: to.Ptr("Success"),
	// 					Target: to.Ptr("KeyEncryptionKey"),
	// 					Details: []*armdatabox.CloudError{
	// 					},
	// 				},
	// 				ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidateDataTransferDetails),
	// 				Status: to.Ptr(armdatabox.ValidationStatusValid),
	// 			},
	// 			&armdatabox.AddressValidationProperties{
	// 				Error: &armdatabox.CloudError{
	// 					AdditionalInfo: []*armdatabox.AdditionalErrorInfo{
	// 					},
	// 					Code: to.Ptr("Success"),
	// 					Target: to.Ptr("KeyEncryptionKey"),
	// 					Details: []*armdatabox.CloudError{
	// 					},
	// 				},
	// 				ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidateAddress),
	// 				AlternateAddresses: []*armdatabox.ShippingAddress{
	// 					{
	// 						AddressType: to.Ptr(armdatabox.AddressTypeNone),
	// 						City: to.Ptr("SAN FRANCISCO"),
	// 						Country: to.Ptr("US"),
	// 						PostalCode: to.Ptr("94107"),
	// 						StateOrProvince: to.Ptr("CA"),
	// 						StreetAddress1: to.Ptr("16 TOWNSEND ST"),
	// 						StreetAddress2: to.Ptr("Unit 1"),
	// 						StreetAddress3: to.Ptr(""),
	// 				}},
	// 				ValidationStatus: to.Ptr(armdatabox.AddressValidationStatusValid),
	// 			},
	// 			&armdatabox.SubscriptionIsAllowedToCreateJobValidationResponseProperties{
	// 				Error: &armdatabox.CloudError{
	// 					AdditionalInfo: []*armdatabox.AdditionalErrorInfo{
	// 					},
	// 					Code: to.Ptr("Success"),
	// 					Target: to.Ptr("KeyEncryptionKey"),
	// 					Details: []*armdatabox.CloudError{
	// 					},
	// 				},
	// 				ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidateSubscriptionIsAllowedToCreateJob),
	// 				Status: to.Ptr(armdatabox.ValidationStatusValid),
	// 			},
	// 			&armdatabox.SKUAvailabilityValidationResponseProperties{
	// 				Error: &armdatabox.CloudError{
	// 					AdditionalInfo: []*armdatabox.AdditionalErrorInfo{
	// 					},
	// 					Code: to.Ptr("Success"),
	// 					Target: to.Ptr("KeyEncryptionKey"),
	// 					Details: []*armdatabox.CloudError{
	// 					},
	// 				},
	// 				ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidateSKUAvailability),
	// 				Status: to.Ptr(armdatabox.ValidationStatusValid),
	// 			},
	// 			&armdatabox.CreateOrderLimitForSubscriptionValidationResponseProperties{
	// 				Error: &armdatabox.CloudError{
	// 					AdditionalInfo: []*armdatabox.AdditionalErrorInfo{
	// 					},
	// 					Code: to.Ptr("Success"),
	// 					Target: to.Ptr("KeyEncryptionKey"),
	// 					Details: []*armdatabox.CloudError{
	// 					},
	// 				},
	// 				ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidateCreateOrderLimit),
	// 				Status: to.Ptr(armdatabox.ValidationStatusValid),
	// 			},
	// 			&armdatabox.PreferencesValidationResponseProperties{
	// 				Error: &armdatabox.CloudError{
	// 					AdditionalInfo: []*armdatabox.AdditionalErrorInfo{
	// 					},
	// 					Code: to.Ptr("Success"),
	// 					Target: to.Ptr("KeyEncryptionKey"),
	// 					Details: []*armdatabox.CloudError{
	// 					},
	// 				},
	// 				ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidatePreferences),
	// 				Status: to.Ptr(armdatabox.ValidationStatusValid),
	// 		}},
	// 		Status: to.Ptr(armdatabox.OverallValidationStatusAllValidToProceed),
	// 	},
	// }
}
