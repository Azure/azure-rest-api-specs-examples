package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"
)

// x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/HostPools_RetrieveRegistrationToken_Post.json
func ExampleHostPoolsClient_RetrieveRegistrationToken() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdesktopvirtualization.NewHostPoolsClient("<subscription-id>", cred, nil)
	res, err := client.RetrieveRegistrationToken(ctx,
		"<resource-group-name>",
		"<host-pool-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.HostPoolsClientRetrieveRegistrationTokenResult)
}
