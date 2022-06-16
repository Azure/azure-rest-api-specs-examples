package armm365securityandcompliance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/SecurityCenterServiceDelete.json
func ExamplePrivateLinkServicesForM365SecurityCenterClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armm365securityandcompliance.NewPrivateLinkServicesForM365SecurityCenterClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"rg1",
		"service1",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
