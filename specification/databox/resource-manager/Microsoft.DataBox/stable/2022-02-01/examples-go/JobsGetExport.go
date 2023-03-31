package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databox/resource-manager/Microsoft.DataBox/stable/2022-02-01/examples/JobsGetExport.json
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
	res, err := clientFactory.NewJobsClient().Get(ctx, "SdkRg8091", "SdkJob6429", &armdatabox.JobsClientGetOptions{Expand: to.Ptr("details")})
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
	// 	Name: to.Ptr("SdkJob6429"),
	// 	Type: to.Ptr("Microsoft.DataBox/jobs"),
	// 	ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/SdkRg8091/providers/Microsoft.DataBox/jobs/SdkJob6429"),
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
	// 			ChainOfCustodySasKey: to.Ptr("https://wusbeta.blob.core.windows.net/chainofcustody/9a6ee052-bcff-4b5e-a478-7dcbfb86e9fb.txt?sv=2018-03-28&sr=b&sig=JbBBXZKharvfg35ZfmrhowO1DuOpzcZCNUdeFzklvWs%3D&st=2020-08-07T05%3A10%3A05Z&se=2020-08-07T05%3A40%3A05Z&sp=r"),
	// 			ContactDetails: &armdatabox.ContactDetails{
	// 				ContactName: to.Ptr(" "),
	// 				EmailList: []*string{
	// 				},
	// 				NotificationPreference: []*armdatabox.NotificationPreference{
	// 				},
	// 				Phone: to.Ptr("1234567890"),
	// 				PhoneExtension: to.Ptr("1234"),
	// 			},
	// 			CopyLogDetails: []armdatabox.CopyLogDetailsClassification{
	// 			},
	// 			DataExportDetails: []*armdatabox.DataExportDetails{
	// 				{
	// 					AccountDetails: &armdatabox.StorageAccountDetails{
	// 						DataAccountType: to.Ptr(armdatabox.DataAccountTypeStorageAccount),
	// 						StorageAccountID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.Storage/storageAccounts/aaaaaa2"),
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
	// 			ReverseShipmentLabelSasKey: to.Ptr("https://wusbeta.blob.core.windows.net/customer-reverse-shipment-instructions/CustomerShipmentInstructions_MicrosoftManaged.PDF?sv=2018-03-28&sr=b&sig=04aL%2FTe7998qQJlhmGI3C0Q%2FxsCVo1t0B4uQUIx7TmQ%3D&st=2020-08-07T05%3A10%3A05Z&se=2020-08-08T05%3A20%3A05Z&sp=r"),
	// 			ShippingAddress: &armdatabox.ShippingAddress{
	// 				AddressType: to.Ptr(armdatabox.AddressTypeCommercial),
	// 				City: to.Ptr("San Francisco"),
	// 				CompanyName: to.Ptr("Microsoft"),
	// 				Country: to.Ptr("US"),
	// 				PostalCode: to.Ptr("94107"),
	// 				StateOrProvince: to.Ptr("CA"),
	// 				StreetAddress1: to.Ptr("16 TOWNSEND ST"),
	// 				StreetAddress2: to.Ptr("Unit 1"),
	// 			},
	// 			CopyProgress: []*armdatabox.CopyProgress{
	// 				{
	// 					AccountID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.Storage/storageAccounts/aaaaaa2"),
	// 					BytesProcessed: to.Ptr[int64](0),
	// 					DataAccountType: to.Ptr(armdatabox.DataAccountTypeStorageAccount),
	// 					IsEnumerationInProgress: to.Ptr(false),
	// 					StorageAccountName: to.Ptr("aaaaaa2"),
	// 					TotalBytesToProcess: to.Ptr[int64](0),
	// 					TransferType: to.Ptr(armdatabox.TransferTypeExportFromAzure),
	// 			}},
	// 		},
	// 	},
	// }
}
