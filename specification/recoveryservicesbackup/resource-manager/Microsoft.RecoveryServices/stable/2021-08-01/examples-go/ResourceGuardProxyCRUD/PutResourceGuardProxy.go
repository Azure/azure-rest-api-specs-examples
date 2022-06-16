package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/ResourceGuardProxyCRUD/PutResourceGuardProxy.json
func ExampleResourceGuardProxyClient_Put() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewResourceGuardProxyClient("<subscription-id>", cred, nil)
	res, err := client.Put(ctx,
		"<vault-name>",
		"<resource-group-name>",
		"<resource-guard-proxy-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ResourceGuardProxyBaseResource.ID: %s\n", *res.ID)
}
