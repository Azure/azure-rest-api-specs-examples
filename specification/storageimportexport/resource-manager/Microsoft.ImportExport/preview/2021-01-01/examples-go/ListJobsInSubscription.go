package armstorageimportexport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storageimportexport/armstorageimportexport"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storageimportexport/resource-manager/Microsoft.ImportExport/preview/2021-01-01/examples/ListJobsInSubscription.json
func ExampleJobsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorageimportexport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobsClient().NewListBySubscriptionPager(&armstorageimportexport.JobsClientListBySubscriptionOptions{Top: nil,
		Filter:         nil,
		AcceptLanguage: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ListJobsResponse = armstorageimportexport.ListJobsResponse{
		// 	Value: []*armstorageimportexport.JobResponse{
		// 		{
		// 			Name: to.Ptr("test-bn1-import-cpu100-01"),
		// 			Type: to.Ptr("Microsoft.ImportExport/jobs"),
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/testrg/providers/Microsoft.ImportExport/jobs/test-bn1-import-cpu100-01"),
		// 			Identity: &armstorageimportexport.IdentityDetails{
		// 				Type: to.Ptr(armstorageimportexport.IdentityTypeNone),
		// 			},
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armstorageimportexport.JobDetails{
		// 				BackupDriveManifest: to.Ptr(true),
		// 				CancelRequested: to.Ptr(false),
		// 				DeliveryPackage: &armstorageimportexport.DeliveryPackageInformation{
		// 					CarrierName: to.Ptr("FedEx"),
		// 					DriveCount: to.Ptr[int64](1),
		// 					ShipDate: to.Ptr("8/11/2017 9:05:00 PM"),
		// 					TrackingNumber: to.Ptr("992016102507"),
		// 				},
		// 				DiagnosticsPath: to.Ptr("waimportexport"),
		// 				DriveList: []*armstorageimportexport.DriveStatus{
		// 					{
		// 						DriveID: to.Ptr("7PHR882C"),
		// 						ManifestFile: to.Ptr("\\DriveManifest.xml"),
		// 						ManifestHash: to.Ptr("E5D632DB047C74B3B17C8F3359950ADB"),
		// 						State: to.Ptr(armstorageimportexport.DriveStateReceived),
		// 				}},
		// 				EncryptionKey: &armstorageimportexport.EncryptionKeyDetails{
		// 					KekType: to.Ptr(armstorageimportexport.EncryptionKekTypeMicrosoftManaged),
		// 				},
		// 				JobType: to.Ptr("Import"),
		// 				LogLevel: to.Ptr("Verbose"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ReturnAddress: &armstorageimportexport.ReturnAddress{
		// 					City: to.Ptr("city"),
		// 					CountryOrRegion: to.Ptr("USA"),
		// 					Email: to.Ptr("test@contoso.com"),
		// 					Phone: to.Ptr("425000000"),
		// 					PostalCode: to.Ptr("98007"),
		// 					RecipientName: to.Ptr("La"),
		// 					StateOrProvince: to.Ptr("wa"),
		// 					StreetAddress1: to.Ptr("Street1"),
		// 					StreetAddress2: to.Ptr("street2"),
		// 				},
		// 				ReturnShipping: &armstorageimportexport.ReturnShipping{
		// 					CarrierAccountNumber: to.Ptr("989ffff"),
		// 					CarrierName: to.Ptr("FedEx"),
		// 				},
		// 				ShippingInformation: &armstorageimportexport.ShippingInformation{
		// 					AdditionalInformation: to.Ptr(""),
		// 					City: to.Ptr("Boydton"),
		// 					CountryOrRegion: to.Ptr("USA"),
		// 					Phone: to.Ptr("+1-434-738-9443"),
		// 					PostalCode: to.Ptr("23917"),
		// 					RecipientName: to.Ptr("Windows Azure Import Export Service"),
		// 					StateOrProvince: to.Ptr("VA"),
		// 					StreetAddress1: to.Ptr("Boydton 1 / BLDG # 1 101 Herbert Drive"),
		// 					StreetAddress2: to.Ptr(""),
		// 				},
		// 				State: to.Ptr("Received"),
		// 				StorageAccountID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/testrg/providers/Microsoft.Storage/storageAccounts/xtprodtestarmos2"),
		// 			},
		// 			SystemData: &armstorageimportexport.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.107Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armstorageimportexport.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.197Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armstorageimportexport.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-bn1-import-cpu100-02"),
		// 			Type: to.Ptr("Microsoft.ImportExport/jobs"),
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/testrg/providers/Microsoft.ImportExport/jobs/test-bn1-import-cpu100-02"),
		// 			Identity: &armstorageimportexport.IdentityDetails{
		// 				Type: to.Ptr(armstorageimportexport.IdentityTypeNone),
		// 			},
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armstorageimportexport.JobDetails{
		// 				BackupDriveManifest: to.Ptr(true),
		// 				CancelRequested: to.Ptr(false),
		// 				DeliveryPackage: &armstorageimportexport.DeliveryPackageInformation{
		// 					CarrierName: to.Ptr("FedEx"),
		// 					DriveCount: to.Ptr[int64](1),
		// 					ShipDate: to.Ptr("8/13/2017 7:32:23 PM"),
		// 					TrackingNumber: to.Ptr("992016102508"),
		// 				},
		// 				DiagnosticsPath: to.Ptr("waimportexport"),
		// 				DriveList: []*armstorageimportexport.DriveStatus{
		// 					{
		// 						CopyStatus: to.Ptr("InProgress"),
		// 						DriveID: to.Ptr("KV6H86XP"),
		// 						ErrorLogURI: to.Ptr("https://foo.blob.core.windows.net/waimportexport/waies/test-bn1-import-cpu100-02_KV6H86XP_20170813-194620-021_error.xml"),
		// 						ManifestFile: to.Ptr("\\DriveManifest.xml"),
		// 						ManifestHash: to.Ptr("F6A488A65AF0CCA7D050B7F9F43A197C"),
		// 						ManifestURI: to.Ptr("https://foo.blob.core.windows.net/waimportexport/waies/test-bn1-import-cpu100-02_KV6H86XP_20170813-194620-574_manifest.xml"),
		// 						State: to.Ptr(armstorageimportexport.DriveStateCompletedMoreInfo),
		// 						VerboseLogURI: to.Ptr("https://foo.blob.core.windows.net/waimportexport/waies/test-bn1-import-cpu100-02_KV6H86XP_20170813-194618-350_verbose.xml"),
		// 				}},
		// 				EncryptionKey: &armstorageimportexport.EncryptionKeyDetails{
		// 					KekType: to.Ptr(armstorageimportexport.EncryptionKekTypeMicrosoftManaged),
		// 				},
		// 				JobType: to.Ptr("Import"),
		// 				LogLevel: to.Ptr("Verbose"),
		// 				PercentComplete: to.Ptr[int64](6),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ReturnAddress: &armstorageimportexport.ReturnAddress{
		// 					City: to.Ptr("city"),
		// 					CountryOrRegion: to.Ptr("USA"),
		// 					Email: to.Ptr("test@contoso.com"),
		// 					Phone: to.Ptr("4256150991"),
		// 					PostalCode: to.Ptr("98007"),
		// 					RecipientName: to.Ptr("La"),
		// 					StateOrProvince: to.Ptr("wa"),
		// 					StreetAddress1: to.Ptr("Street1"),
		// 					StreetAddress2: to.Ptr("street2"),
		// 				},
		// 				ShippingInformation: &armstorageimportexport.ShippingInformation{
		// 					AdditionalInformation: to.Ptr(""),
		// 					City: to.Ptr("Boydton"),
		// 					CountryOrRegion: to.Ptr("USA"),
		// 					Phone: to.Ptr("+1-434-738-9443"),
		// 					PostalCode: to.Ptr("23917"),
		// 					RecipientName: to.Ptr("Windows Azure Import Export Service"),
		// 					StateOrProvince: to.Ptr("VA"),
		// 					StreetAddress1: to.Ptr("Boydton 1 / BLDG # 1 101 Herbert Drive"),
		// 					StreetAddress2: to.Ptr(""),
		// 				},
		// 				State: to.Ptr("Transferring"),
		// 				StorageAccountID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/testrg/providers/Microsoft.Storage/storageAccounts/xtprodtestarmos2"),
		// 			},
		// 			SystemData: &armstorageimportexport.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.107Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armstorageimportexport.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.197Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armstorageimportexport.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
