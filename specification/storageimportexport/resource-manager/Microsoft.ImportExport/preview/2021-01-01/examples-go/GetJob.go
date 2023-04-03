package armstorageimportexport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storageimportexport/armstorageimportexport"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storageimportexport/resource-manager/Microsoft.ImportExport/preview/2021-01-01/examples/GetJob.json
func ExampleJobsClient_Get_getImportJob() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorageimportexport.NewClientFactory("<subscription-id>", nil, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobsClient().Get(ctx, "myJob", "myResourceGroup", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JobResponse = armstorageimportexport.JobResponse{
	// 	Name: to.Ptr("myJob"),
	// 	Type: to.Ptr("Microsoft.ImportExport/jobs"),
	// 	ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.ImportExport/jobs/test"),
	// 	Identity: &armstorageimportexport.IdentityDetails{
	// 		Type: to.Ptr(armstorageimportexport.IdentityTypeNone),
	// 	},
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armstorageimportexport.JobDetails{
	// 		BackupDriveManifest: to.Ptr(true),
	// 		CancelRequested: to.Ptr(false),
	// 		DiagnosticsPath: to.Ptr("waimportexport"),
	// 		DriveList: []*armstorageimportexport.DriveStatus{
	// 			{
	// 				DriveHeaderHash: to.Ptr(""),
	// 				DriveID: to.Ptr("9CA995BB"),
	// 				ManifestFile: to.Ptr("\\DriveManifest.xml"),
	// 				ManifestHash: to.Ptr("109B21108597EF36D5785F08303F3638"),
	// 				State: to.Ptr(armstorageimportexport.DriveStateSpecified),
	// 		}},
	// 		EncryptionKey: &armstorageimportexport.EncryptionKeyDetails{
	// 			KekType: to.Ptr(armstorageimportexport.EncryptionKekTypeMicrosoftManaged),
	// 		},
	// 		JobType: to.Ptr("Import"),
	// 		LogLevel: to.Ptr("Verbose"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ReturnAddress: &armstorageimportexport.ReturnAddress{
	// 			City: to.Ptr("Redmond"),
	// 			CountryOrRegion: to.Ptr("USA"),
	// 			Email: to.Ptr("Test@contoso.com"),
	// 			Phone: to.Ptr("4250000000"),
	// 			PostalCode: to.Ptr("98007"),
	// 			RecipientName: to.Ptr("Test"),
	// 			StateOrProvince: to.Ptr("wa"),
	// 			StreetAddress1: to.Ptr("Street1"),
	// 			StreetAddress2: to.Ptr("street2"),
	// 		},
	// 		ReturnShipping: &armstorageimportexport.ReturnShipping{
	// 			CarrierAccountNumber: to.Ptr("989ffff"),
	// 			CarrierName: to.Ptr("FedEx"),
	// 		},
	// 		ShippingInformation: &armstorageimportexport.ShippingInformation{
	// 			AdditionalInformation: to.Ptr(""),
	// 			City: to.Ptr("Santa Clara"),
	// 			CountryOrRegion: to.Ptr("USA"),
	// 			Phone: to.Ptr("408 0000 0000"),
	// 			PostalCode: to.Ptr("95050"),
	// 			RecipientName: to.Ptr("Microsoft Azure Import/Export Service"),
	// 			StateOrProvince: to.Ptr("CA"),
	// 			StreetAddress1: to.Ptr("2045 Lafayette Street"),
	// 			StreetAddress2: to.Ptr(""),
	// 		},
	// 		State: to.Ptr("Creating"),
	// 		StorageAccountID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.ClassicStorage/storageAccounts/test"),
	// 	},
	// 	SystemData: &armstorageimportexport.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.1075056Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armstorageimportexport.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.1974346Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armstorageimportexport.CreatedByTypeUser),
	// 	},
	// }
}
