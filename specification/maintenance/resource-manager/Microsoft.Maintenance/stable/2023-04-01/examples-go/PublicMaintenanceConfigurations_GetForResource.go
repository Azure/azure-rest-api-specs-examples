package armmaintenance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2023-04-01/examples/PublicMaintenanceConfigurations_GetForResource.json
func ExamplePublicMaintenanceConfigurationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaintenance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPublicMaintenanceConfigurationsClient().Get(ctx, "configuration1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Configuration = armmaintenance.Configuration{
	// 	Name: to.Ptr("configuration1"),
	// 	ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/configuration1"),
	// 	Location: to.Ptr("westus2"),
	// 	Properties: &armmaintenance.ConfigurationProperties{
	// 		MaintenanceScope: to.Ptr(armmaintenance.MaintenanceScopeSQLDB),
	// 		MaintenanceWindow: &armmaintenance.Window{
	// 			Duration: to.Ptr("05:00"),
	// 			ExpirationDateTime: to.Ptr("9999-12-31 00:00"),
	// 			RecurEvery: to.Ptr("2Weeks"),
	// 			StartDateTime: to.Ptr("2020-04-30 08:00"),
	// 			TimeZone: to.Ptr("Pacific Standard Time"),
	// 		},
	// 		Namespace: to.Ptr("Microsoft.Maintenance"),
	// 		Visibility: to.Ptr(armmaintenance.VisibilityPublic),
	// 	},
	// }
}
