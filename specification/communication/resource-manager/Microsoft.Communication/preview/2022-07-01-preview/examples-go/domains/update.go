package armcommunication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/communication/resource-manager/Microsoft.Communication/preview/2022-07-01-preview/examples/domains/update.json
func ExampleDomainsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcommunication.NewDomainsClient("12345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx, "MyResourceGroup", "MyEmailServiceResource", "mydomain.com", armcommunication.UpdateDomainRequestParameters{
		Properties: &armcommunication.UpdateDomainProperties{
			UserEngagementTracking: to.Ptr(armcommunication.UserEngagementTrackingEnabled),
			ValidSenderUsernames: map[string]*string{
				"info":   to.Ptr("MyDomain Info"),
				"alerts": to.Ptr("MyDomain Alerts"),
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
