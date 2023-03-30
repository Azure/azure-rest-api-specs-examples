package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databox/resource-manager/Microsoft.DataBox/stable/2022-02-01/examples/JobsGetCopyStuck.json
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
	res, err := clientFactory.NewJobsClient().Get(ctx, "dmstestresource", "TJx-637505258985313014", &armdatabox.JobsClientGetOptions{Expand: to.Ptr("details")})
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
	// 	Name: to.Ptr("TJx-637505258985313014"),
	// 	Type: to.Ptr("Microsoft.DataBox/jobs"),
	// 	ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/dmstestresource/providers/Microsoft.DataBox/jobs/TJx-637505258985313014"),
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
	// 				ContactName: to.Ptr("Andrew Tribone"),
	// 				EmailList: []*string{
	// 					to.Ptr("ssemmail@microsoft.com"),
	// 					to.Ptr("vishwamdir@microsoft.com")},
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
	// 					Phone: to.Ptr("1234567890"),
	// 					PhoneExtension: to.Ptr("1234"),
	// 				},
	// 				CopyLogDetails: []armdatabox.CopyLogDetailsClassification{
	// 					&armdatabox.AccountCopyLogDetails{
	// 						CopyLogDetailsType: to.Ptr(armdatabox.ClassDiscriminatorDataBox),
	// 						AccountName: to.Ptr("databoxbvttestaccount"),
	// 						CopyLogLink: to.Ptr("databoxcopylog/strtrinidad01_ZTS18520041_CopyLog_8425fb8f9c2b447288caa4dd6f5d34bb.xml"),
	// 				}},
	// 				DataImportDetails: []*armdatabox.DataImportDetails{
	// 					{
	// 						AccountDetails: &armdatabox.StorageAccountDetails{
	// 							DataAccountType: to.Ptr(armdatabox.DataAccountTypeStorageAccount),
	// 							StorageAccountID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/databoxbvt/providers/Microsoft.Storage/storageAccounts/databoxbvttestaccount"),
	// 						},
	// 				}},
	// 				DeliveryPackage: &armdatabox.PackageShippingDetails{
	// 					CarrierName: to.Ptr("Ups"),
	// 					TrackingID: to.Ptr("5fb6965e-0b80-4f38-b21a-32673fed2d84"),
	// 					TrackingURL: to.Ptr("https://wwwapps.ups.com/WebTracking/track?track=yes&trackNums=5fb6965e-0b80-4f38-b21a-32673fed2d84"),
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
	// 					TrackingURL: to.Ptr("https://wwwapps.ups.com/WebTracking/track?track=yes&trackNums=b3875a34-aad6-4dbe-83a9-3f39cb21b0e8"),
	// 				},
	// 				ReverseShipmentLabelSasKey: to.Ptr("http://wusbeta.blob.core.windows.net/devstoreaccount1/f6be6ea8-77da-419b-9f6b-3043c28b04e4/ReverseShipment/b3875a34-aad6-4dbe-83a9-3f39cb21b0e8.PDF?sv=2018-03-28&sr=b&sig=%2BJmOyIkG0ALrKQp%2F7NCOBYccQ80DRLvN%2FHtb6Z97heg%3D&st=2021-03-05T09%3A45%3A05Z&se=2021-03-05T10%3A15%3A05Z&sp=r"),
	// 				ShippingAddress: &armdatabox.ShippingAddress{
	// 					AddressType: to.Ptr(armdatabox.AddressTypeNone),
	// 					City: to.Ptr("San Francisco"),
	// 					CompanyName: to.Ptr("Microsoft"),
	// 					Country: to.Ptr("US"),
	// 					PostalCode: to.Ptr("94107"),
	// 					StateOrProvince: to.Ptr("CA"),
	// 					StreetAddress1: to.Ptr("164 TOWNSEND ST"),
	// 					StreetAddress2: to.Ptr("UNIT 1"),
	// 				},
	// 				CopyProgress: []*armdatabox.CopyProgress{
	// 					{
	// 						AccountID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/databoxbvt/providers/Microsoft.Storage/storageAccounts/databoxbvttestaccount"),
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
	// 						StorageAccountName: to.Ptr("databoxbvttestaccount"),
	// 						TotalBytesToProcess: to.Ptr[int64](2000),
	// 						TotalFilesToProcess: to.Ptr[int64](110),
	// 						TransferType: to.Ptr(armdatabox.TransferTypeImportToAzure),
	// 				}},
	// 			},
	// 		},
	// 	}
}
