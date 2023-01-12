package armservicenetworking_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicenetworking/armservicenetworking"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/FrontendPut.json
func ExampleFrontendsInterfaceClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicenetworking.NewFrontendsInterfaceClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "TC1", "publicIp1", armservicenetworking.Frontend{
		Location: to.Ptr("West US"),
		Properties: &armservicenetworking.FrontendProperties{
			IPAddressVersion: to.Ptr(armservicenetworking.FrontendIPAddressVersionIPv4),
			Mode:             to.Ptr("public"),
			PublicIPAddress: &armservicenetworking.FrontendPropertiesIPAddress{
				ID: to.Ptr("resourceUriAsString"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
