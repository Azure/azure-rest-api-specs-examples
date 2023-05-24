package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/JobsGetExport.json
func ExampleJobsClient_Get_jobsGetExport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabox.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobsClient().Get(ctx, "YourResourceGroupName", "TestJobName1", &armdatabox.JobsClientGetOptions{Expand: to.Ptr("details")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JobResource = armdatabox.JobResource{
	// 	Identity: &armdatabox.ResourceIdentity{
	// 		Type: to.Ptr("None"),
	// 	},
	// 	Location: to.Ptr("westus"),
	// 	SKU: &armdatabox.SKU{
	// 		Name: to.Ptr(armdatabox.SKUNameDataBox),
	// 	},
	// 	Tags: map[string]*string{
	// 	},
	// 	Name: to.Ptr("TestJobName1"),
	// 	Type: to.Ptr("Microsoft.DataBox/jobs"),
	// 	ID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.DataBox/jobs/TestJobName1"),
	// 	Properties: &armdatabox.JobProperties{
	// 		CancellationReason: to.Ptr("CancelTest"),
	// 		DeliveryInfo: &armdatabox.JobDeliveryInfo{
	// 			ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
	// 		},
	// 		DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
	// 		IsCancellable: to.Ptr(false),
	// 		IsCancellableWithoutFee: to.Ptr(false),
	// 		IsDeletable: to.Ptr(true),
	// 		IsPrepareToShipEnabled: to.Ptr(true),
	// 		IsShippingAddressEditable: to.Ptr(false),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-08-07T10:49:08.7195419+05:30"); return t}()),
	// 		Status: to.Ptr(armdatabox.StageNameCancelled),
	// 		TransferType: to.Ptr(armdatabox.TransferTypeExportFromAzure),
	// 		Details: &armdatabox.JobDetails{
	// 			ChainOfCustodySasKey: to.Ptr("https://xxx.xxx.xx"),
	// 			ContactDetails: &armdatabox.ContactDetails{
	// 				ContactName: to.Ptr(" "),
	// 				EmailList: []*string{
	// 				},
	// 				NotificationPreference: []*armdatabox.NotificationPreference{
	// 				},
	// 				Phone: to.Ptr("0000000000"),
	// 				PhoneExtension: to.Ptr(""),
	// 			},
	// 			CopyLogDetails: []armdatabox.CopyLogDetailsClassification{
	// 			},
	// 			DataExportDetails: []*armdatabox.DataExportDetails{
	// 				{
	// 					AccountDetails: &armdatabox.StorageAccountDetails{
	// 						DataAccountType: to.Ptr(armdatabox.DataAccountTypeStorageAccount),
	// 						StorageAccountID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Storage/storageAccounts/YourStorageAccountName"),
	// 					},
	// 					LogCollectionLevel: to.Ptr(armdatabox.LogCollectionLevelError),
	// 					TransferConfiguration: &armdatabox.TransferConfiguration{
	// 						TransferAllDetails: &armdatabox.TransferConfigurationTransferAllDetails{
	// 							Include: &armdatabox.TransferAllDetails{
	// 								DataAccountType: to.Ptr(armdatabox.DataAccountTypeStorageAccount),
	// 								TransferAllBlobs: to.Ptr(true),
	// 								TransferAllFiles: to.Ptr(true),
	// 							},
	// 						},
	// 						TransferConfigurationType: to.Ptr(armdatabox.TransferConfigurationTypeTransferAll),
	// 						TransferFilterDetails: &armdatabox.TransferConfigurationTransferFilterDetails{
	// 						},
	// 					},
	// 			}},
	// 			DataImportDetails: []*armdatabox.DataImportDetails{
	// 			},
	// 			DeliveryPackage: &armdatabox.PackageShippingDetails{
	// 				CarrierName: to.Ptr(""),
	// 				TrackingID: to.Ptr(""),
	// 				TrackingURL: to.Ptr(""),
	// 			},
	// 			JobDetailsType: to.Ptr(armdatabox.ClassDiscriminatorDataBox),
	// 			JobStages: []*armdatabox.JobStages{
	// 				{
	// 					DisplayName: to.Ptr("Ordered"),
	// 					StageName: to.Ptr(armdatabox.StageNameDeviceOrdered),
	// 					StageStatus: to.Ptr(armdatabox.StageStatusSucceeded),
	// 					StageTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-08-07T10:49:12.7675644+05:30"); return t}()),
	// 				},
	// 				{
	// 					DisplayName: to.Ptr("Canceled"),
	// 					StageName: to.Ptr(armdatabox.StageNameCancelled),
	// 					StageStatus: to.Ptr(armdatabox.StageStatusSucceeded),
	// 					StageTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-08-07T10:49:57.2572438+05:30"); return t}()),
	// 			}},
	// 			KeyEncryptionKey: &armdatabox.KeyEncryptionKey{
	// 				KekType: to.Ptr(armdatabox.KekTypeMicrosoftManaged),
	// 			},
	// 			ReturnPackage: &armdatabox.PackageShippingDetails{
	// 				CarrierName: to.Ptr(""),
	// 				TrackingID: to.Ptr(""),
	// 				TrackingURL: to.Ptr(""),
	// 			},
	// 			ReverseShipmentLabelSasKey: to.Ptr("https://xxx.xxx.xx"),
	// 			ShippingAddress: &armdatabox.ShippingAddress{
	// 				AddressType: to.Ptr(armdatabox.AddressTypeCommercial),
	// 				City: to.Ptr("XXXX XXXX"),
	// 				CompanyName: to.Ptr("XXXX XXXX"),
	// 				Country: to.Ptr("XX"),
	// 				PostalCode: to.Ptr("00000"),
	// 				StateOrProvince: to.Ptr("XX"),
	// 				StreetAddress1: to.Ptr("XXXX XXXX"),
	// 				StreetAddress2: to.Ptr("XXXX XXXX"),
	// 			},
	// 			CopyProgress: []*armdatabox.CopyProgress{
	// 				{
	// 					AccountID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Storage/storageAccounts/YourStorageAccountName"),
	// 					BytesProcessed: to.Ptr[int64](0),
	// 					DataAccountType: to.Ptr(armdatabox.DataAccountTypeStorageAccount),
	// 					IsEnumerationInProgress: to.Ptr(false),
	// 					StorageAccountName: to.Ptr("YourStorageAccountName"),
	// 					TotalBytesToProcess: to.Ptr[int64](0),
	// 					TransferType: to.Ptr(armdatabox.TransferTypeExportFromAzure),
	// 			}},
	// 		},
	// 	},
	// }
}
