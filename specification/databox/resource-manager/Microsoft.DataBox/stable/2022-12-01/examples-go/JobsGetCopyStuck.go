package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/JobsGetCopyStuck.json
func ExampleJobsClient_Get_jobsGetCopyStuck() {
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
	// 		"defaultTagsKey": to.Ptr("defaultTagsValue"),
	// 	},
	// 	Name: to.Ptr("TestJobName1"),
	// 	Type: to.Ptr("Microsoft.DataBox/jobs"),
	// 	ID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.DataBox/jobs/TestJobName1"),
	// 	Properties: &armdatabox.JobProperties{
	// 		DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
	// 		IsCancellable: to.Ptr(false),
	// 		IsCancellableWithoutFee: to.Ptr(false),
	// 		IsDeletable: to.Ptr(true),
	// 		IsPrepareToShipEnabled: to.Ptr(true),
	// 		IsShippingAddressEditable: to.Ptr(false),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-05T12:56:12.6384315+05:30"); return t}()),
	// 		Status: to.Ptr(armdatabox.StageNameCompletedWithWarnings),
	// 		TransferType: to.Ptr(armdatabox.TransferTypeImportToAzure),
	// 		Details: &armdatabox.JobDetails{
	// 			Actions: []*armdatabox.CustomerResolutionCode{
	// 			},
	// 			ContactDetails: &armdatabox.ContactDetails{
	// 				ContactName: to.Ptr("XXXX XXXX"),
	// 				EmailList: []*string{
	// 					to.Ptr("xxxx@xxxx.xxx")},
	// 					NotificationPreference: []*armdatabox.NotificationPreference{
	// 						{
	// 							SendNotification: to.Ptr(true),
	// 							StageName: to.Ptr(armdatabox.NotificationStageNameDevicePrepared),
	// 						},
	// 						{
	// 							SendNotification: to.Ptr(true),
	// 							StageName: to.Ptr(armdatabox.NotificationStageNameDispatched),
	// 						},
	// 						{
	// 							SendNotification: to.Ptr(true),
	// 							StageName: to.Ptr(armdatabox.NotificationStageNameDelivered),
	// 						},
	// 						{
	// 							SendNotification: to.Ptr(true),
	// 							StageName: to.Ptr(armdatabox.NotificationStageNamePickedUp),
	// 						},
	// 						{
	// 							SendNotification: to.Ptr(true),
	// 							StageName: to.Ptr(armdatabox.NotificationStageNameAtAzureDC),
	// 						},
	// 						{
	// 							SendNotification: to.Ptr(true),
	// 							StageName: to.Ptr(armdatabox.NotificationStageNameDataCopy),
	// 					}},
	// 					Phone: to.Ptr("0000000000"),
	// 					PhoneExtension: to.Ptr(""),
	// 				},
	// 				CopyLogDetails: []armdatabox.CopyLogDetailsClassification{
	// 					&armdatabox.AccountCopyLogDetails{
	// 						CopyLogDetailsType: to.Ptr(armdatabox.ClassDiscriminatorDataBox),
	// 						AccountName: to.Ptr("YourStorageAccountName"),
	// 						CopyLogLink: to.Ptr("databoxcopylog/xxx.xml"),
	// 				}},
	// 				DataImportDetails: []*armdatabox.DataImportDetails{
	// 					{
	// 						AccountDetails: &armdatabox.StorageAccountDetails{
	// 							DataAccountType: to.Ptr(armdatabox.DataAccountTypeStorageAccount),
	// 							StorageAccountID: to.Ptr("/subscriptions/YourSubscriptionId/resourcegroups/YourResourceGroupName/providers/Microsoft.Storage/storageAccounts/YourStorageAccountName"),
	// 						},
	// 				}},
	// 				DeliveryPackage: &armdatabox.PackageShippingDetails{
	// 					CarrierName: to.Ptr("Ups"),
	// 					TrackingID: to.Ptr("5fb6965e-0b80-4f38-b21a-32673fed2d84"),
	// 					TrackingURL: to.Ptr("https://xxx.xxx.xx"),
	// 				},
	// 				ExpectedDataSizeInTeraBytes: to.Ptr[int32](0),
	// 				JobDetailsType: to.Ptr(armdatabox.ClassDiscriminatorDataBox),
	// 				JobStages: []*armdatabox.JobStages{
	// 					{
	// 						DisplayName: to.Ptr("Ordered"),
	// 						StageName: to.Ptr(armdatabox.StageNameDeviceOrdered),
	// 						StageStatus: to.Ptr(armdatabox.StageStatusSucceeded),
	// 						StageTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-05T12:56:25.3911023+05:30"); return t}()),
	// 					},
	// 					{
	// 						DisplayName: to.Ptr("Processed"),
	// 						StageName: to.Ptr(armdatabox.StageNameDevicePrepared),
	// 						StageStatus: to.Ptr(armdatabox.StageStatusSucceeded),
	// 						StageTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-05T13:07:49.676421+05:30"); return t}()),
	// 					},
	// 					{
	// 						DisplayName: to.Ptr("Dispatched"),
	// 						StageName: to.Ptr(armdatabox.StageNameDispatched),
	// 						StageStatus: to.Ptr(armdatabox.StageStatusSucceeded),
	// 						StageTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-05T02:21:06+05:30"); return t}()),
	// 					},
	// 					{
	// 						DisplayName: to.Ptr("Delivered"),
	// 						StageName: to.Ptr(armdatabox.StageNameDelivered),
	// 						StageStatus: to.Ptr(armdatabox.StageStatusSucceeded),
	// 						StageTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-05T02:21:13+05:30"); return t}()),
	// 					},
	// 					{
	// 						DisplayName: to.Ptr("Picked up"),
	// 						StageName: to.Ptr(armdatabox.StageNamePickedUp),
	// 						StageStatus: to.Ptr(armdatabox.StageStatusSucceeded),
	// 						StageTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-05T02:23:10+05:30"); return t}()),
	// 					},
	// 					{
	// 						DisplayName: to.Ptr("Received"),
	// 						StageName: to.Ptr(armdatabox.StageNameAtAzureDC),
	// 						StageStatus: to.Ptr(armdatabox.StageStatusSucceeded),
	// 						StageTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-05T02:23:17+05:30"); return t}()),
	// 					},
	// 					{
	// 						DisplayName: to.Ptr("Data copy in progress"),
	// 						StageName: to.Ptr(armdatabox.StageNameDataCopy),
	// 						StageStatus: to.Ptr(armdatabox.StageStatusSucceeded),
	// 						StageTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-05T13:34:20.1497635+05:30"); return t}()),
	// 					},
	// 					{
	// 						StageName: to.Ptr(armdatabox.StageNameCompletedWithWarnings),
	// 						StageStatus: to.Ptr(armdatabox.StageStatusSucceeded),
	// 						StageTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-05T14:08:38.250614+05:30"); return t}()),
	// 				}},
	// 				KeyEncryptionKey: &armdatabox.KeyEncryptionKey{
	// 					KekType: to.Ptr(armdatabox.KekTypeMicrosoftManaged),
	// 				},
	// 				LastMitigationActionOnJob: &armdatabox.LastMitigationActionOnJob{
	// 					ActionDateTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-05T08:28:18.8107951Z"); return t}()),
	// 					CustomerResolution: to.Ptr(armdatabox.CustomerResolutionCodeMoveToCleanUpDevice),
	// 					IsPerformedByCustomer: to.Ptr(true),
	// 				},
	// 				Preferences: &armdatabox.Preferences{
	// 				},
	// 				ReturnPackage: &armdatabox.PackageShippingDetails{
	// 					CarrierName: to.Ptr("Ups"),
	// 					TrackingID: to.Ptr("b3875a34-aad6-4dbe-83a9-3f39cb21b0e8"),
	// 					TrackingURL: to.Ptr("https://xxx.xxx.xx"),
	// 				},
	// 				ReverseShipmentLabelSasKey: to.Ptr("http://xxx.xxx.xxx"),
	// 				ShippingAddress: &armdatabox.ShippingAddress{
	// 					AddressType: to.Ptr(armdatabox.AddressTypeNone),
	// 					City: to.Ptr("XXXX XXXX"),
	// 					CompanyName: to.Ptr("XXXX XXXX"),
	// 					Country: to.Ptr("XX"),
	// 					PostalCode: to.Ptr("00000"),
	// 					StateOrProvince: to.Ptr("XX"),
	// 					StreetAddress1: to.Ptr("164 TOWNSEND ST"),
	// 					StreetAddress2: to.Ptr("XXXX XXXX"),
	// 				},
	// 				CopyProgress: []*armdatabox.CopyProgress{
	// 					{
	// 						AccountID: to.Ptr("/subscriptions/YourSubscriptionId/resourcegroups/YourResourceGroupName/providers/Microsoft.Storage/storageAccounts/YourStorageAccountName"),
	// 						BytesProcessed: to.Ptr[int64](2000),
	// 						DataAccountType: to.Ptr(armdatabox.DataAccountTypeStorageAccount),
	// 						DirectoriesErroredOut: to.Ptr[int64](0),
	// 						FilesErroredOut: to.Ptr[int64](0),
	// 						FilesProcessed: to.Ptr[int64](100),
	// 						InvalidDirectoriesProcessed: to.Ptr[int64](0),
	// 						InvalidFileBytesUploaded: to.Ptr[int64](10),
	// 						InvalidFilesProcessed: to.Ptr[int64](10),
	// 						IsEnumerationInProgress: to.Ptr(false),
	// 						RenamedContainerCount: to.Ptr[int64](60),
	// 						StorageAccountName: to.Ptr("YourStorageAccountName"),
	// 						TotalBytesToProcess: to.Ptr[int64](2000),
	// 						TotalFilesToProcess: to.Ptr[int64](110),
	// 						TransferType: to.Ptr(armdatabox.TransferTypeImportToAzure),
	// 				}},
	// 			},
	// 		},
	// 	}
}
