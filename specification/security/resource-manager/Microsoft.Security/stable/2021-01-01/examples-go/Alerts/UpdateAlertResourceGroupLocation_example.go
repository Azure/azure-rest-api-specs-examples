package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-01-01/examples/Alerts/UpdateAlertResourceGroupLocation_example.json
func ExampleAlertsClient_UpdateResourceGroupLevelStateToActivate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewAlertsClient("<subscription-id>",
		"<asc-location>", cred, nil)
	_, err = client.UpdateResourceGroupLevelStateToActivate(ctx,
		"<alert-name>",
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
