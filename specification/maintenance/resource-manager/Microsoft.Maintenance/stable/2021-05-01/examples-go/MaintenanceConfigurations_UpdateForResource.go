package armmaintenance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2021-05-01/examples/MaintenanceConfigurations_UpdateForResource.json
func ExampleConfigurationsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmaintenance.NewConfigurationsClient("5b4b650e-28b9-4790-b3ab-ddbd88d727c4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"examplerg",
		"configuration1",
		armmaintenance.Configuration{
			Location: to.Ptr("westus2"),
			Properties: &armmaintenance.ConfigurationProperties{
				MaintenanceScope: to.Ptr(armmaintenance.MaintenanceScopeHost),
				MaintenanceWindow: &armmaintenance.Window{
					Duration:           to.Ptr("05:00"),
					ExpirationDateTime: to.Ptr("9999-12-31 00:00"),
					RecurEvery:         to.Ptr("Month Third Sunday"),
					StartDateTime:      to.Ptr("2020-04-30 08:00"),
					TimeZone:           to.Ptr("Pacific Standard Time"),
				},
				Namespace:  to.Ptr("Microsoft.Maintenance"),
				Visibility: to.Ptr(armmaintenance.VisibilityCustom),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
