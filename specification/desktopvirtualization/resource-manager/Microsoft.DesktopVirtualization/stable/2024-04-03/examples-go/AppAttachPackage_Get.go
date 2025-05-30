package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4883fa5dbf6f2c9093fac8ce334547e9dfac68fa/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/AppAttachPackage_Get.json
func ExampleAppAttachPackageClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAppAttachPackageClient().Get(ctx, "resourceGroup1", "packagefullname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AppAttachPackage = armdesktopvirtualization.AppAttachPackage{
	// 	Name: to.Ptr("packageName"),
	// 	Type: to.Ptr("Microsoft.DesktopVirtualization/appattachpackages"),
	// 	ID: to.Ptr("/subscriptions/d15725f7-6577-4a8c-95f1-3da903b42364/resourcegroups/charlesk-southcentralus/providers/Microsoft.DesktopVirtualization/appattachpackages/ModifierPackage"),
	// 	SystemData: &armdesktopvirtualization.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-28T23:44:56.130Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-09T01:43:31.070Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user1"),
	// 		LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("southcentralus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armdesktopvirtualization.AppAttachPackageProperties{
	// 		FailHealthCheckOnStagingFailure: to.Ptr(armdesktopvirtualization.FailHealthCheckOnStagingFailureNeedsAssistance),
	// 		HostPoolReferences: []*string{
	// 		},
	// 		Image: &armdesktopvirtualization.AppAttachPackageInfoProperties{
	// 			CertificateExpiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-02T17:18:19.123Z"); return t}()),
	// 			CertificateName: to.Ptr("certName"),
	// 			DisplayName: to.Ptr("displayname"),
	// 			ImagePath: to.Ptr("imagepath"),
	// 			IsActive: to.Ptr(false),
	// 			IsRegularRegistration: to.Ptr(false),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.957Z"); return t}()),
	// 			PackageAlias: to.Ptr("msixpackagealias"),
	// 			PackageApplications: []*armdesktopvirtualization.MsixPackageApplications{
	// 				{
	// 					Description: to.Ptr("PackageApplicationDescription"),
	// 					AppID: to.Ptr("AppId"),
	// 					AppUserModelID: to.Ptr("AppUserModelId"),
	// 					FriendlyName: to.Ptr("FriendlyName"),
	// 					IconImageName: to.Ptr("Iconimagename"),
	// 					RawIcon: []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
	// 					RawPNG: []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
	// 			}},
	// 			PackageDependencies: []*armdesktopvirtualization.MsixPackageDependencies{
	// 				{
	// 					DependencyName: to.Ptr("MsixPackage_Dependency_Name"),
	// 					MinVersion: to.Ptr("packageDep_version"),
	// 					Publisher: to.Ptr("MsixPackage_Dependency_Publisher"),
	// 			}},
	// 			PackageFamilyName: to.Ptr("MsixPackage_FamilyName"),
	// 			PackageFullName: to.Ptr("MsixPackage_FullName"),
	// 			PackageName: to.Ptr("MsixPackageName"),
	// 			PackageRelativePath: to.Ptr("packagerelativepath"),
	// 			Version: to.Ptr("packageversion"),
	// 		},
	// 		KeyVaultURL: to.Ptr(""),
	// 	},
	// }
}
