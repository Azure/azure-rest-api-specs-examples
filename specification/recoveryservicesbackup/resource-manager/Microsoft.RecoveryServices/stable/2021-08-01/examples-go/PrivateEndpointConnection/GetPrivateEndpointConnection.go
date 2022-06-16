package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/PrivateEndpointConnection/GetPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewPrivateEndpointConnectionClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<vault-name>",
		"<resource-group-name>",
		"<private-endpoint-connection-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PrivateEndpointConnectionResource.ID: %s\n", *res.ID)
}
