package armdesktopvirtualization_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"
)

// x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/MsixPackage_Create.json
func ExampleMSIXPackagesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdesktopvirtualization.NewMSIXPackagesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<host-pool-name>",
		"<msix-package-full-name>",
		armdesktopvirtualization.MSIXPackage{
			Properties: &armdesktopvirtualization.MSIXPackageProperties{
				DisplayName:           to.StringPtr("<display-name>"),
				ImagePath:             to.StringPtr("<image-path>"),
				IsActive:              to.BoolPtr(false),
				IsRegularRegistration: to.BoolPtr(false),
				LastUpdated:           to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.9571247Z"); return t }()),
				PackageApplications: []*armdesktopvirtualization.MsixPackageApplications{
					{
						Description:    to.StringPtr("<description>"),
						AppID:          to.StringPtr("<app-id>"),
						AppUserModelID: to.StringPtr("<app-user-model-id>"),
						FriendlyName:   to.StringPtr("<friendly-name>"),
						IconImageName:  to.StringPtr("<icon-image-name>"),
						RawIcon:        []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
						RawPNG:         []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
					}},
				PackageDependencies: []*armdesktopvirtualization.MsixPackageDependencies{
					{
						DependencyName: to.StringPtr("<dependency-name>"),
						MinVersion:     to.StringPtr("<min-version>"),
						Publisher:      to.StringPtr("<publisher>"),
					}},
				PackageFamilyName:   to.StringPtr("<package-family-name>"),
				PackageName:         to.StringPtr("<package-name>"),
				PackageRelativePath: to.StringPtr("<package-relative-path>"),
				Version:             to.StringPtr("<version>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MSIXPackagesClientCreateOrUpdateResult)
}
