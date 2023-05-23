package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/JobsGetWaitingForAction.json
func ExampleJobsClient_Get_jobsGetWaitingForAction() {
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
	// 		Error: &armdatabox.CloudError{
	// 			AdditionalInfo: []*armdatabox.AdditionalErrorInfo{
	// 			},
	// 			Code: to.Ptr("SsemUserCopyErrorWaitingForAction"),
	// 			Message: to.Ptr("Job has an error in copy stage and is waiting for customer action.  Please review the error and select one of the actions provided in the job's properties.details.actions"),
	// 			Target: to.Ptr("CopyIntervention"),
	// 			Details: []*armdatabox.CloudError{
	// 			},
	// 		},
	// 		IsCancellable: to.Ptr(false),
	// 		IsCancellableWithoutFee: to.Ptr(false),
	// 		IsDeletable: to.Ptr(false),
	// 		IsPrepareToShipEnabled: to.Ptr(true),
	// 		IsShippingAddressEditable: to.Ptr(false),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-14T23:49:49.5177298+00:00"); return t}()),
	// 		Status: to.Ptr(armdatabox.StageNameDataCopy),
	// 		TransferType: to.Ptr(armdatabox.TransferTypeImportToAzure),
	// 		Details: &armdatabox.JobDetails{
	// 			Actions: []*armdatabox.CustomerResolutionCode{
	// 				to.Ptr(armdatabox.CustomerResolutionCodeMoveToCleanUpDevice)},
	// 				ContactDetails: &armdatabox.ContactDetails{
	// 					ContactName: to.Ptr("XXXX XXXX"),
	// 					EmailList: []*string{
	// 						to.Ptr("xxxx@xxxx.xxx")},
	// 						NotificationPreference: []*armdatabox.NotificationPreference{
	// 							{
	// 								SendNotification: to.Ptr(true),
	// 								StageName: to.Ptr(armdatabox.NotificationStageNameDevicePrepared),
	// 							},
	// 							{
	// 								SendNotification: to.Ptr(true),
	// 								StageName: to.Ptr(armdatabox.NotificationStageNameDispatched),
	// 							},
	// 							{
	// 								SendNotification: to.Ptr(true),
	// 								StageName: to.Ptr(armdatabox.NotificationStageNameDelivered),
	// 							},
	// 							{
	// 								SendNotification: to.Ptr(true),
	// 								StageName: to.Ptr(armdatabox.NotificationStageNamePickedUp),
	// 							},
	// 							{
	// 								SendNotification: to.Ptr(true),
	// 								StageName: to.Ptr(armdatabox.NotificationStageNameAtAzureDC),
	// 							},
	// 							{
	// 								SendNotification: to.Ptr(true),
	// 								StageName: to.Ptr(armdatabox.NotificationStageNameDataCopy),
	// 						}},
	// 						Phone: to.Ptr("0000000000"),
	// 						PhoneExtension: to.Ptr(""),
	// 					},
	// 					CopyLogDetails: []armdatabox.CopyLogDetailsClassification{
	// 						&armdatabox.AccountCopyLogDetails{
	// 							CopyLogDetailsType: to.Ptr(armdatabox.ClassDiscriminatorDataBox),
	// 							AccountName: to.Ptr("YourStorageAccountName"),
	// 							CopyLogLink: to.Ptr("databoxcopylog/xxx.xml"),
	// 						},
	// 						&armdatabox.AccountCopyLogDetails{
	// 							CopyLogDetailsType: to.Ptr(armdatabox.ClassDiscriminatorDataBox),
	// 							AccountName: to.Ptr("YourResourceGroupName"),
	// 							CopyLogLink: to.Ptr("databoxcopylog/xxx.xml"),
	// 					}},
	// 					DataImportDetails: []*armdatabox.DataImportDetails{
	// 						{
	// 							AccountDetails: &armdatabox.StorageAccountDetails{
	// 								DataAccountType: to.Ptr(armdatabox.DataAccountTypeStorageAccount),
	// 								StorageAccountID: to.Ptr("/subscriptions/YourSubscriptionId/resourcegroups/YourResourceGroupName/providers/Microsoft.Storage/storageAccounts/YourStorageAccountName"),
	// 							},
	// 						},
	// 						{
	// 							AccountDetails: &armdatabox.ManagedDiskDetails{
	// 								DataAccountType: to.Ptr(armdatabox.DataAccountTypeManagedDisk),
	// 								ResourceGroupID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName"),
	// 								StagingStorageAccountID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Storage/storageAccounts/YourStorageAccountName"),
	// 							},
	// 					}},
	// 					DeliveryPackage: &armdatabox.PackageShippingDetails{
	// 						CarrierName: to.Ptr("Ups"),
	// 						TrackingID: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
	// 						TrackingURL: to.Ptr("https://xxx.xxx.xx"),
	// 					},
	// 					ExpectedDataSizeInTeraBytes: to.Ptr[int32](0),
	// 					JobDetailsType: to.Ptr(armdatabox.ClassDiscriminatorDataBox),
	// 					JobStages: []*armdatabox.JobStages{
	// 						{
	// 							DisplayName: to.Ptr("Ordered"),
	// 							StageName: to.Ptr(armdatabox.StageNameDeviceOrdered),
	// 							StageStatus: to.Ptr(armdatabox.StageStatusSucceeded),
	// 							StageTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-14T23:49:53.81509+00:00"); return t}()),
	// 						},
	// 						{
	// 							DisplayName: to.Ptr("Processed"),
	// 							StageName: to.Ptr(armdatabox.StageNameDevicePrepared),
	// 							StageStatus: to.Ptr(armdatabox.StageStatusSucceeded),
	// 							StageTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-15T00:05:18.7350971+00:00"); return t}()),
	// 						},
	// 						{
	// 							DisplayName: to.Ptr("Dispatched"),
	// 							StageName: to.Ptr(armdatabox.StageNameDispatched),
	// 							StageStatus: to.Ptr(armdatabox.StageStatusSucceeded),
	// 							StageTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-15T00:25:19+00:00"); return t}()),
	// 						},
	// 						{
	// 							DisplayName: to.Ptr("Delivered"),
	// 							StageName: to.Ptr(armdatabox.StageNameDelivered),
	// 							StageStatus: to.Ptr(armdatabox.StageStatusSucceeded),
	// 							StageTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-15T00:25:26+00:00"); return t}()),
	// 						},
	// 						{
	// 							DisplayName: to.Ptr("Picked up"),
	// 							StageName: to.Ptr(armdatabox.StageNamePickedUp),
	// 							StageStatus: to.Ptr(armdatabox.StageStatusSucceeded),
	// 							StageTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-15T00:30:22+00:00"); return t}()),
	// 						},
	// 						{
	// 							DisplayName: to.Ptr("Received"),
	// 							StageName: to.Ptr(armdatabox.StageNameAtAzureDC),
	// 							StageStatus: to.Ptr(armdatabox.StageStatusSucceeded),
	// 							StageTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-15T00:30:29+00:00"); return t}()),
	// 						},
	// 						{
	// 							DisplayName: to.Ptr("Data copy in progress"),
	// 							StageName: to.Ptr(armdatabox.StageNameDataCopy),
	// 							StageStatus: to.Ptr(armdatabox.StageStatusWaitingForCustomerAction),
	// 							StageTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-15T00:45:26.2679756+00:00"); return t}()),
	// 						},
	// 						{
	// 							DisplayName: to.Ptr("Completed"),
	// 							StageName: to.Ptr(armdatabox.StageNameCompleted),
	// 							StageStatus: to.Ptr(armdatabox.StageStatusNone),
	// 					}},
	// 					KeyEncryptionKey: &armdatabox.KeyEncryptionKey{
	// 						KekType: to.Ptr(armdatabox.KekTypeMicrosoftManaged),
	// 					},
	// 					LastMitigationActionOnJob: &armdatabox.LastMitigationActionOnJob{
	// 						ActionDateTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-12T05:00:22.5047578Z"); return t}()),
	// 						CustomerResolution: to.Ptr(armdatabox.CustomerResolutionCodeMoveToCleanUpDevice),
	// 						IsPerformedByCustomer: to.Ptr(true),
	// 					},
	// 					Preferences: &armdatabox.Preferences{
	// 						EncryptionPreferences: &armdatabox.EncryptionPreferences{
	// 							DoubleEncryption: to.Ptr(armdatabox.DoubleEncryptionDisabled),
	// 						},
	// 					},
	// 					ReturnPackage: &armdatabox.PackageShippingDetails{
	// 						CarrierName: to.Ptr("Ups"),
	// 						TrackingID: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
	// 						TrackingURL: to.Ptr("https://xxx.xxx.xx"),
	// 					},
	// 					ReverseShipmentLabelSasKey: to.Ptr("https://xxx.xxx.xx"),
	// 					ShippingAddress: &armdatabox.ShippingAddress{
	// 						AddressType: to.Ptr(armdatabox.AddressTypeNone),
	// 						City: to.Ptr("XXXX XXXX"),
	// 						CompanyName: to.Ptr("XXXX XXXX"),
	// 						Country: to.Ptr("XX"),
	// 						PostalCode: to.Ptr("00000"),
	// 						StateOrProvince: to.Ptr("XX"),
	// 						StreetAddress1: to.Ptr("164 TOWNSEND ST"),
	// 						StreetAddress2: to.Ptr("XXXX XXXX"),
	// 					},
	// 					CopyProgress: []*armdatabox.CopyProgress{
	// 						{
	// 							AccountID: to.Ptr("/subscriptions/YourSubscriptionId/resourcegroups/YourResourceGroupName/providers/Microsoft.Storage/storageAccounts/YourStorageAccountName"),
	// 							BytesProcessed: to.Ptr[int64](2000),
	// 							DataAccountType: to.Ptr(armdatabox.DataAccountTypeStorageAccount),
	// 							DirectoriesErroredOut: to.Ptr[int64](0),
	// 							FilesErroredOut: to.Ptr[int64](0),
	// 							FilesProcessed: to.Ptr[int64](100),
	// 							InvalidDirectoriesProcessed: to.Ptr[int64](0),
	// 							InvalidFileBytesUploaded: to.Ptr[int64](10),
	// 							InvalidFilesProcessed: to.Ptr[int64](10),
	// 							IsEnumerationInProgress: to.Ptr(false),
	// 							RenamedContainerCount: to.Ptr[int64](60),
	// 							StorageAccountName: to.Ptr("YourStorageAccountName"),
	// 							TotalBytesToProcess: to.Ptr[int64](2000),
	// 							TotalFilesToProcess: to.Ptr[int64](110),
	// 							TransferType: to.Ptr(armdatabox.TransferTypeImportToAzure),
	// 						},
	// 						{
	// 							AccountID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName"),
	// 							BytesProcessed: to.Ptr[int64](2000),
	// 							DataAccountType: to.Ptr(armdatabox.DataAccountTypeManagedDisk),
	// 							DirectoriesErroredOut: to.Ptr[int64](0),
	// 							FilesErroredOut: to.Ptr[int64](0),
	// 							FilesProcessed: to.Ptr[int64](100),
	// 							InvalidDirectoriesProcessed: to.Ptr[int64](0),
	// 							InvalidFileBytesUploaded: to.Ptr[int64](10),
	// 							InvalidFilesProcessed: to.Ptr[int64](10),
	// 							IsEnumerationInProgress: to.Ptr(false),
	// 							RenamedContainerCount: to.Ptr[int64](60),
	// 							TotalBytesToProcess: to.Ptr[int64](2000),
	// 							TotalFilesToProcess: to.Ptr[int64](110),
	// 							TransferType: to.Ptr(armdatabox.TransferTypeImportToAzure),
	// 					}},
	// 				},
	// 			},
	// 		}
}
