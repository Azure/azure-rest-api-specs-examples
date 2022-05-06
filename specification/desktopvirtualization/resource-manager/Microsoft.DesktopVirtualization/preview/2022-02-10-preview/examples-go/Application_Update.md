Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdesktopvirtualization%2Farmdesktopvirtualization%2Fv0.4.0/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/README.md) on how to add the SDK to your project and authenticate.

```go
package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2022-02-10-preview/examples/Application_Update.json
func ExampleApplicationsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdesktopvirtualization.NewApplicationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<application-group-name>",
		"<application-name>",
		&armdesktopvirtualization.ApplicationsClientUpdateOptions{Application: &armdesktopvirtualization.ApplicationPatch{
			Properties: &armdesktopvirtualization.ApplicationPatchProperties{
				Description:          to.Ptr("<description>"),
				ApplicationType:      to.Ptr(armdesktopvirtualization.RemoteApplicationTypeInBuilt),
				CommandLineArguments: to.Ptr("<command-line-arguments>"),
				CommandLineSetting:   to.Ptr(armdesktopvirtualization.CommandLineSettingAllow),
				FilePath:             to.Ptr("<file-path>"),
				FriendlyName:         to.Ptr("<friendly-name>"),
				IconIndex:            to.Ptr[int32](1),
				IconPath:             to.Ptr("<icon-path>"),
				ShowInPortal:         to.Ptr(true),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
