package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2021-07-12/examples/Application_Update.json
func ExampleApplicationsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdesktopvirtualization.NewApplicationsClient("daefabc0-95b4-48b3-b645-8a753a63c4fa", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"resourceGroup1",
		"applicationGroup1",
		"application1",
		&armdesktopvirtualization.ApplicationsClientUpdateOptions{Application: &armdesktopvirtualization.ApplicationPatch{
			Properties: &armdesktopvirtualization.ApplicationPatchProperties{
				Description:          to.Ptr("des1"),
				ApplicationType:      to.Ptr(armdesktopvirtualization.RemoteApplicationTypeInBuilt),
				CommandLineArguments: to.Ptr("arguments"),
				CommandLineSetting:   to.Ptr(armdesktopvirtualization.CommandLineSettingAllow),
				FilePath:             to.Ptr("path"),
				FriendlyName:         to.Ptr("friendly"),
				IconIndex:            to.Ptr[int32](1),
				IconPath:             to.Ptr("icon"),
				ShowInPortal:         to.Ptr(true),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
