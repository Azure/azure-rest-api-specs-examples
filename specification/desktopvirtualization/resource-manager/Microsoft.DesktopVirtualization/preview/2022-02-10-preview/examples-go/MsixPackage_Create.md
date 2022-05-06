Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdesktopvirtualization%2Farmdesktopvirtualization%2Fv0.4.0/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/README.md) on how to add the SDK to your project and authenticate.

```go
package armdesktopvirtualization_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2022-02-10-preview/examples/MsixPackage_Create.json
func ExampleMSIXPackagesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdesktopvirtualization.NewMSIXPackagesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<host-pool-name>",
		"<msix-package-full-name>",
		armdesktopvirtualization.MSIXPackage{
			Properties: &armdesktopvirtualization.MSIXPackageProperties{
				DisplayName:           to.Ptr("<display-name>"),
				ImagePath:             to.Ptr("<image-path>"),
				IsActive:              to.Ptr(false),
				IsRegularRegistration: to.Ptr(false),
				LastUpdated:           to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.9571247Z"); return t }()),
				PackageApplications: []*armdesktopvirtualization.MsixPackageApplications{
					{
						Description:    to.Ptr("<description>"),
						AppID:          to.Ptr("<app-id>"),
						AppUserModelID: to.Ptr("<app-user-model-id>"),
						FriendlyName:   to.Ptr("<friendly-name>"),
						IconImageName:  to.Ptr("<icon-image-name>"),
						RawIcon:        []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
						RawPNG:         []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
					}},
				PackageDependencies: []*armdesktopvirtualization.MsixPackageDependencies{
					{
						DependencyName: to.Ptr("<dependency-name>"),
						MinVersion:     to.Ptr("<min-version>"),
						Publisher:      to.Ptr("<publisher>"),
					}},
				PackageFamilyName:   to.Ptr("<package-family-name>"),
				PackageName:         to.Ptr("<package-name>"),
				PackageRelativePath: to.Ptr("<package-relative-path>"),
				Version:             to.Ptr("<version>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
