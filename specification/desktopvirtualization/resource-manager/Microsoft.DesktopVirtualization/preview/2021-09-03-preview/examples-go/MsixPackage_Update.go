package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"
)

// x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/MsixPackage_Update.json
func ExampleMSIXPackagesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdesktopvirtualization.NewMSIXPackagesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<host-pool-name>",
		"<msix-package-full-name>",
		&armdesktopvirtualization.MSIXPackagesClientUpdateOptions{MsixPackage: &armdesktopvirtualization.MSIXPackagePatch{
			Properties: &armdesktopvirtualization.MSIXPackagePatchProperties{
				DisplayName:           to.StringPtr("<display-name>"),
				IsActive:              to.BoolPtr(true),
				IsRegularRegistration: to.BoolPtr(false),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MSIXPackagesClientUpdateResult)
}
