package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databox/resource-manager/Microsoft.DataBox/stable/2022-02-01/examples/JobsCreateDoubleEncryption.json
func ExampleJobsClient_BeginCreate_jobsCreateDoubleEncryption() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabox.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewJobsClient().BeginCreate(ctx, "SdkRg608", "SdkJob6599", armdatabox.JobResource{
		Location: to.Ptr("westus"),
		SKU: &armdatabox.SKU{
			Name: to.Ptr(armdatabox.SKUNameDataBox),
		},
		Properties: &armdatabox.JobProperties{
			TransferType: to.Ptr(armdatabox.TransferTypeImportToAzure),
			Details: &armdatabox.JobDetails{
				ContactDetails: &armdatabox.ContactDetails{
					ContactName: to.Ptr("Public SDK Test"),
					EmailList: []*string{
						to.Ptr("testing@microsoft.com")},
					Phone:          to.Ptr("1234567890"),
					PhoneExtension: to.Ptr("1234"),
				},
				DataImportDetails: []*armdatabox.DataImportDetails{
					{
						AccountDetails: &armdatabox.StorageAccountDetails{
							DataAccountType:  to.Ptr(armdatabox.DataAccountTypeStorageAccount),
							StorageAccountID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/databoxbvt/providers/Microsoft.Storage/storageAccounts/databoxbvttestaccount"),
						},
					}},
				JobDetailsType: to.Ptr(armdatabox.ClassDiscriminatorDataBox),
				Preferences: &armdatabox.Preferences{
					EncryptionPreferences: &armdatabox.EncryptionPreferences{
						DoubleEncryption: to.Ptr(armdatabox.DoubleEncryptionEnabled),
					},
				},
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
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JobResource = armdatabox.JobResource{
	// 	Location: to.Ptr("westus"),
	// 	SKU: &armdatabox.SKU{
	// 		Name: to.Ptr(armdatabox.SKUNameDataBox),
	// 	},
	// 	Tags: map[string]*string{
	// 	},
	// 	Name: to.Ptr("SdkJob6599"),
	// 	Type: to.Ptr("Microsoft.DataBox/jobs"),
	// 	ID: to.Ptr("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/SdkRg608/providers/Microsoft.DataBox/jobs/SdkJob6599"),
	// 	Properties: &armdatabox.JobProperties{
	// 		DeliveryInfo: &armdatabox.JobDeliveryInfo{
	// 			ScheduledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
	// 		},
	// 		DeliveryType: to.Ptr(armdatabox.JobDeliveryTypeNonScheduled),
	// 		IsCancellable: to.Ptr(true),
	// 		IsCancellableWithoutFee: to.Ptr(true),
	// 		IsDeletable: to.Ptr(false),
	// 		IsShippingAddressEditable: to.Ptr(true),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-21T14:49:49.0027611+05:30"); return t}()),
	// 		Status: to.Ptr(armdatabox.StageNameDeviceOrdered),
	// 		TransferType: to.Ptr(armdatabox.TransferTypeExportFromAzure),
	// 		Details: &armdatabox.JobDetails{
	// 			ContactDetails: &armdatabox.ContactDetails{
	// 				ContactName: to.Ptr("Public SDK Test"),
	// 				EmailList: []*string{
	// 					to.Ptr("testing@microsoft.com")},
	// 					NotificationPreference: []*armdatabox.NotificationPreference{
	// 						{
	// 							SendNotification: to.Ptr(true),
	// 							StageName: to.Ptr(armdatabox.NotificationStageNameDevicePrepared),
	// 						},
	// 						{
	// 							SendNotification: to.Ptr(true),
	// 							StageName: to.Ptr(armdatabox.NotificationStageNameDataCopy),
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
	// 					}},
	// 					Phone: to.Ptr("1234567890"),
	// 					PhoneExtension: to.Ptr("1234"),
	// 				},
	// 				CopyLogDetails: []armdatabox.CopyLogDetailsClassification{
	// 				},
	// 				JobDetailsType: to.Ptr(armdatabox.ClassDiscriminatorDataBox),
	// 				JobStages: []*armdatabox.JobStages{
	// 					{
	// 						DisplayName: to.Ptr("Ordered"),
	// 						StageName: to.Ptr(armdatabox.StageNameDeviceOrdered),
	// 						StageStatus: to.Ptr(armdatabox.StageStatusSucceeded),
	// 						StageTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-21T14:49:52.3995743+05:30"); return t}()),
	// 					},
	// 					{
	// 						DisplayName: to.Ptr("Processed"),
	// 						StageName: to.Ptr(armdatabox.StageNameDevicePrepared),
	// 						StageStatus: to.Ptr(armdatabox.StageStatusNone),
	// 					},
	// 					{
	// 						DisplayName: to.Ptr("Data copy in progress"),
	// 						StageName: to.Ptr(armdatabox.StageNameDataCopy),
	// 						StageStatus: to.Ptr(armdatabox.StageStatusNone),
	// 					},
	// 					{
	// 						DisplayName: to.Ptr("Dispatched"),
	// 						StageName: to.Ptr(armdatabox.StageNameDispatched),
	// 						StageStatus: to.Ptr(armdatabox.StageStatusNone),
	// 					},
	// 					{
	// 						DisplayName: to.Ptr("Delivered"),
	// 						StageName: to.Ptr(armdatabox.StageNameDelivered),
	// 						StageStatus: to.Ptr(armdatabox.StageStatusNone),
	// 					},
	// 					{
	// 						DisplayName: to.Ptr("Picked up"),
	// 						StageName: to.Ptr(armdatabox.StageNamePickedUp),
	// 						StageStatus: to.Ptr(armdatabox.StageStatusNone),
	// 					},
	// 					{
	// 						DisplayName: to.Ptr("Received"),
	// 						StageName: to.Ptr(armdatabox.StageNameAtAzureDC),
	// 						StageStatus: to.Ptr(armdatabox.StageStatusNone),
	// 					},
	// 					{
	// 						DisplayName: to.Ptr("Completed"),
	// 						StageName: to.Ptr(armdatabox.StageNameCompleted),
	// 						StageStatus: to.Ptr(armdatabox.StageStatusNone),
	// 				}},
	// 				ShippingAddress: &armdatabox.ShippingAddress{
	// 					AddressType: to.Ptr(armdatabox.AddressTypeCommercial),
	// 					City: to.Ptr("San Francisco"),
	// 					CompanyName: to.Ptr("Microsoft"),
	// 					Country: to.Ptr("US"),
	// 					PostalCode: to.Ptr("94107"),
	// 					StateOrProvince: to.Ptr("CA"),
	// 					StreetAddress1: to.Ptr("16 TOWNSEND ST"),
	// 					StreetAddress2: to.Ptr("Unit 1"),
	// 				},
	// 				CopyProgress: []*armdatabox.CopyProgress{
	// 				},
	// 			},
	// 		},
	// 	}
}
