package armmaintenance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2023-04-01/examples/ConfigurationAssignmentsForResourceGroup_UpdateForResource.json
func ExampleConfigurationAssignmentsForResourceGroupClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaintenance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConfigurationAssignmentsForResourceGroupClient().Update(ctx, "examplerg", "workervmConfiguration", armmaintenance.ConfigurationAssignment{
		Properties: &armmaintenance.ConfigurationAssignmentProperties{
			Filter: &armmaintenance.ConfigurationAssignmentFilterProperties{
				Locations: []*string{
					to.Ptr("Japan East"),
					to.Ptr("UK South")},
				ResourceTypes: []*string{
					to.Ptr("Microsoft.HybridCompute/machines"),
					to.Ptr("Microsoft.Compute/virtualMachines")},
				TagSettings: &armmaintenance.TagSettingsProperties{
					FilterOperator: to.Ptr(armmaintenance.TagOperatorsAny),
					Tags: map[string][]*string{
						"tag1": {
							to.Ptr("tag1Value1"),
							to.Ptr("tag1Value2"),
							to.Ptr("tag1Value3")},
						"tag2": {
							to.Ptr("tag2Value1"),
							to.Ptr("tag2Value2"),
							to.Ptr("tag2Value3")},
					},
				},
			},
			MaintenanceConfigurationID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Maintenance/maintenanceConfigurations/configuration1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConfigurationAssignment = armmaintenance.ConfigurationAssignment{
	// 	Name: to.Ptr("workervmConfiguration"),
	// 	Type: to.Ptr("Microsoft.Maintenance/configurationAssignments"),
	// 	ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Maintenance/configurationAssignments/workervmConfiguration"),
	// 	Properties: &armmaintenance.ConfigurationAssignmentProperties{
	// 		Filter: &armmaintenance.ConfigurationAssignmentFilterProperties{
	// 			Locations: []*string{
	// 				to.Ptr("Japan East"),
	// 				to.Ptr("UK South")},
	// 				ResourceTypes: []*string{
	// 					to.Ptr("Microsoft.HybridCompute/machines"),
	// 					to.Ptr("Microsoft.Compute/virtualMachines")},
	// 					TagSettings: &armmaintenance.TagSettingsProperties{
	// 						FilterOperator: to.Ptr(armmaintenance.TagOperatorsAny),
	// 						Tags: map[string][]*string{
	// 							"tag1": []*string{
	// 								to.Ptr("tag1Value1"),
	// 								to.Ptr("tag1Value2"),
	// 								to.Ptr("tag1Value3")},
	// 								"tag2": []*string{
	// 									to.Ptr("tag2Value1"),
	// 									to.Ptr("tag2Value2"),
	// 									to.Ptr("tag2Value3")},
	// 								},
	// 							},
	// 						},
	// 						MaintenanceConfigurationID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Maintenance/maintenanceConfigurations/configuration1"),
	// 						ResourceID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg"),
	// 					},
	// 				}
}
