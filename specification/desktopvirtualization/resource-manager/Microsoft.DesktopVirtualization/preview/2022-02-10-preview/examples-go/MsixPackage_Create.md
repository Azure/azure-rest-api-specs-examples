```go
package armdesktopvirtualization_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2022-02-10-preview/examples/MsixPackage_Create.json
func ExampleMSIXPackagesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdesktopvirtualization.NewMSIXPackagesClient("daefabc0-95b4-48b3-b645-8a753a63c4fa", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"resourceGroup1",
		"hostpool1",
		"msixpackagefullname",
		armdesktopvirtualization.MSIXPackage{
			Properties: &armdesktopvirtualization.MSIXPackageProperties{
				DisplayName:           to.Ptr("displayname"),
				ImagePath:             to.Ptr("imagepath"),
				IsActive:              to.Ptr(false),
				IsRegularRegistration: to.Ptr(false),
				LastUpdated:           to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.9571247Z"); return t }()),
				PackageApplications: []*armdesktopvirtualization.MsixPackageApplications{
					{
						Description:    to.Ptr("application-desc"),
						AppID:          to.Ptr("ApplicationId"),
						AppUserModelID: to.Ptr("AppUserModelId"),
						FriendlyName:   to.Ptr("friendlyname"),
						IconImageName:  to.Ptr("Apptile"),
						RawIcon:        []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
						RawPNG:         []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
					}},
				PackageDependencies: []*armdesktopvirtualization.MsixPackageDependencies{
					{
						DependencyName: to.Ptr("MsixTest_Dependency_Name"),
						MinVersion:     to.Ptr("version"),
						Publisher:      to.Ptr("PublishedName"),
					}},
				PackageFamilyName:   to.Ptr("MsixPackage_FamilyName"),
				PackageName:         to.Ptr("MsixPackage_name"),
				PackageRelativePath: to.Ptr("packagerelativepath"),
				Version:             to.Ptr("version"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdesktopvirtualization%2Farmdesktopvirtualization%2Fv2.0.0-beta.1/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/README.md) on how to add the SDK to your project and authenticate.
