package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/ResourceGuardProxyCRUD/DeleteResourceGuardProxy.json
func ExampleDppResourceGuardProxyClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewDppResourceGuardProxyClient("5e13b949-1218-4d18-8b99-7e12155ec4f7", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "SampleResourceGroup", "sampleVault", "swaggerExample", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
