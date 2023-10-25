package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/MsixPackage_Update.json
func ExampleMSIXPackagesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMSIXPackagesClient().Update(ctx, "resourceGroup1", "hostpool1", "msixpackagefullname", &armdesktopvirtualization.MSIXPackagesClientUpdateOptions{MsixPackage: &armdesktopvirtualization.MSIXPackagePatch{
		Properties: &armdesktopvirtualization.MSIXPackagePatchProperties{
			DisplayName:           to.Ptr("displayname"),
			IsActive:              to.Ptr(true),
			IsRegularRegistration: to.Ptr(false),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MSIXPackage = armdesktopvirtualization.MSIXPackage{
	// 	Name: to.Ptr("hostpool1/MsixPackageFullName"),
	// 	Type: to.Ptr("Microsoft.DesktopVirtualization/hostpools/msixpackages"),
	// 	ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourcegroups/resourcegroup1/providers/Microsoft.DesktopVirtualization/hostpools/hostpool1/msixpackages/msixPackageFullName"),
	// 	Properties: &armdesktopvirtualization.MSIXPackageProperties{
	// 		DisplayName: to.Ptr("dis"),
	// 		ImagePath: to.Ptr("imagepath"),
	// 		IsActive: to.Ptr(true),
	// 		IsRegularRegistration: to.Ptr(false),
	// 		LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.9571247Z"); return t}()),
	// 		PackageApplications: []*armdesktopvirtualization.MsixPackageApplications{
	// 			{
	// 				Description: to.Ptr("desc"),
	// 				AppID: to.Ptr("Application_Id"),
	// 				AppUserModelID: to.Ptr("Application_ModelID"),
	// 				FriendlyName: to.Ptr("fri"),
	// 				IconImageName: to.Ptr("Apptile"),
	// 				RawIcon: []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
	// 				RawPNG: []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
	// 		}},
	// 		PackageDependencies: []*armdesktopvirtualization.MsixPackageDependencies{
	// 			{
	// 				DependencyName: to.Ptr("MsixPackage_Dependency_Name"),
	// 				MinVersion: to.Ptr("packageDep_version"),
	// 				Publisher: to.Ptr("MsixPackage_Dependency_Publisher"),
	// 		}},
	// 		PackageFamilyName: to.Ptr("MsixPackage_FamilyName"),
	// 		PackageName: to.Ptr("MsixPackage_Name"),
	// 		PackageRelativePath: to.Ptr("MsixPackage_RelativePackageRoot"),
	// 		Version: to.Ptr("version"),
	// 	},
	// 	SystemData: &armdesktopvirtualization.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 	},
	// }
}
