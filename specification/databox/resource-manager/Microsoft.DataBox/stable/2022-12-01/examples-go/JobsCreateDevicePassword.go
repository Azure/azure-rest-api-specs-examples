package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/JobsCreateDevicePassword.json
func ExampleJobsClient_BeginCreate_jobsCreateDevicePassword() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabox.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewJobsClient().BeginCreate(ctx, "YourResourceGroupName", "TestJobName1", armdatabox.JobResource{
		Location: to.Ptr("westus"),
		SKU: &armdatabox.SKU{
			Name: to.Ptr(armdatabox.SKUNameDataBox),
		},
		Properties: &armdatabox.JobProperties{
			TransferType: to.Ptr(armdatabox.TransferTypeImportToAzure),
			Details: &armdatabox.JobDetails{
				ContactDetails: &armdatabox.ContactDetails{
					ContactName: to.Ptr("XXXX XXXX"),
					EmailList: []*string{
						to.Ptr("xxxx@xxxx.xxx")},
					Phone:          to.Ptr("0000000000"),
					PhoneExtension: to.Ptr(""),
				},
				DataImportDetails: []*armdatabox.DataImportDetails{
					{
						AccountDetails: &armdatabox.StorageAccountDetails{
							DataAccountType:  to.Ptr(armdatabox.DataAccountTypeStorageAccount),
							SharePassword:    to.Ptr("<sharePassword>"),
							StorageAccountID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Storage/storageAccounts/YourStorageAccountName"),
						},
					}},
				JobDetailsType: to.Ptr(armdatabox.ClassDiscriminatorDataBox),
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
				DevicePassword: to.Ptr("<devicePassword>"),
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
	// 	Name: to.Ptr("TestJobName1"),
	// 	Type: to.Ptr("Microsoft.DataBox/jobs"),
	// 	ID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.DataBox/jobs/TestJobName1"),
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
	// 		TransferType: to.Ptr(armdatabox.TransferTypeImportToAzure),
	// 		Details: &armdatabox.JobDetails{
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
	// 						DisplayName: to.Ptr("Data copy in progress"),
	// 						StageName: to.Ptr(armdatabox.StageNameDataCopy),
	// 						StageStatus: to.Ptr(armdatabox.StageStatusNone),
	// 					},
	// 					{
	// 						DisplayName: to.Ptr("Completed"),
	// 						StageName: to.Ptr(armdatabox.StageNameCompleted),
	// 						StageStatus: to.Ptr(armdatabox.StageStatusNone),
	// 				}},
	// 				ShippingAddress: &armdatabox.ShippingAddress{
	// 					AddressType: to.Ptr(armdatabox.AddressTypeCommercial),
	// 					City: to.Ptr("XXXX XXXX"),
	// 					CompanyName: to.Ptr("XXXX XXXX"),
	// 					Country: to.Ptr("XX"),
	// 					PostalCode: to.Ptr("00000"),
	// 					StateOrProvince: to.Ptr("XX"),
	// 					StreetAddress1: to.Ptr("XXXX XXXX"),
	// 					StreetAddress2: to.Ptr("XXXX XXXX"),
	// 				},
	// 				CopyProgress: []*armdatabox.CopyProgress{
	// 				},
	// 			},
	// 		},
	// 	}
}
