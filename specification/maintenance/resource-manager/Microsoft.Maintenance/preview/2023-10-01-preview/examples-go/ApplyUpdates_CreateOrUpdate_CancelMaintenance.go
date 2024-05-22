package armmaintenance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2023-10-01-preview/examples/ApplyUpdates_CreateOrUpdate_CancelMaintenance.json
func ExampleApplyUpdatesClient_CreateOrUpdateOrCancel_applyUpdatesCreateOrUpdateOrCancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaintenance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplyUpdatesClient().CreateOrUpdateOrCancel(ctx, "examplerg", "Microsoft.Maintenance", "maintenanceConfigurations", "maintenanceConfig1", "20230901121200", armmaintenance.ApplyUpdate{
		Properties: &armmaintenance.ApplyUpdateProperties{
			Status: to.Ptr(armmaintenance.UpdateStatusCancel),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ApplyUpdate = armmaintenance.ApplyUpdate{
	// 	Name: to.Ptr("maintenanceConfig1"),
	// 	Type: to.Ptr("Microsoft.Maintenance/applyUpdates"),
	// 	ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/microsoft.maintenance/maintenanceconfigurations/maintenanceconfig1/providers/microsoft.maintenance/applyupdates/20230901121200"),
	// 	Properties: &armmaintenance.ApplyUpdateProperties{
	// 		ResourceID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/microsoft.maintenance/maintenanceconfigurations/maintenanceconfig1"),
	// 		Status: to.Ptr(armmaintenance.UpdateStatusCancelled),
	// 	},
	// }
}
