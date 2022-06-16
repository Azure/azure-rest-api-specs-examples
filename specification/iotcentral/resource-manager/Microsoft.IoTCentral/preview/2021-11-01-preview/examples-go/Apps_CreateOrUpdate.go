package armiotcentral_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotcentral/armiotcentral"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/Apps_CreateOrUpdate.json
func ExampleAppsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armiotcentral.NewAppsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armiotcentral.App{
			Location: to.Ptr("<location>"),
			Identity: &armiotcentral.SystemAssignedServiceIdentity{
				Type: to.Ptr(armiotcentral.SystemAssignedServiceIdentityTypeSystemAssigned),
			},
			Properties: &armiotcentral.AppProperties{
				DisplayName: to.Ptr("<display-name>"),
				Subdomain:   to.Ptr("<subdomain>"),
				Template:    to.Ptr("<template>"),
			},
			SKU: &armiotcentral.AppSKUInfo{
				Name: to.Ptr(armiotcentral.AppSKUST2),
			},
		},
		&armiotcentral.AppsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
