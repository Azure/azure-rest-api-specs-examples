package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"
)

// x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/SessionHost_Update.json
func ExampleSessionHostsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdesktopvirtualization.NewSessionHostsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<host-pool-name>",
		"<session-host-name>",
		&armdesktopvirtualization.SessionHostsClientUpdateOptions{Force: to.BoolPtr(true),
			SessionHost: &armdesktopvirtualization.SessionHostPatch{
				Properties: &armdesktopvirtualization.SessionHostPatchProperties{
					AllowNewSession: to.BoolPtr(true),
					AssignedUser:    to.StringPtr("<assigned-user>"),
				},
			},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SessionHostsClientUpdateResult)
}
