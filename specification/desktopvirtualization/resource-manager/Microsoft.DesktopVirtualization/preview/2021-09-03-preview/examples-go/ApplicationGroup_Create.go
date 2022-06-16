package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"
)

// x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/ApplicationGroup_Create.json
func ExampleApplicationGroupsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdesktopvirtualization.NewApplicationGroupsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<application-group-name>",
		armdesktopvirtualization.ApplicationGroup{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"tag1": to.StringPtr("value1"),
				"tag2": to.StringPtr("value2"),
			},
			Properties: &armdesktopvirtualization.ApplicationGroupProperties{
				Description:          to.StringPtr("<description>"),
				ApplicationGroupType: armdesktopvirtualization.ApplicationGroupType("RemoteApp").ToPtr(),
				FriendlyName:         to.StringPtr("<friendly-name>"),
				HostPoolArmPath:      to.StringPtr("<host-pool-arm-path>"),
				MigrationRequest: &armdesktopvirtualization.MigrationRequestProperties{
					MigrationPath: to.StringPtr("<migration-path>"),
					Operation:     armdesktopvirtualization.Operation("Start").ToPtr(),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ApplicationGroupsClientCreateOrUpdateResult)
}
