package armmaintenance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2023-10-01-preview/examples/ConfigurationAssignments_CreateOrUpdateParent.json
func ExampleConfigurationAssignmentsClient_CreateOrUpdateParent() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaintenance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConfigurationAssignmentsClient().CreateOrUpdateParent(ctx, "examplerg", "Microsoft.Compute", "virtualMachineScaleSets", "smdtest1", "virtualMachines", "smdvm1", "workervmPolicy", armmaintenance.ConfigurationAssignment{
		Properties: &armmaintenance.ConfigurationAssignmentProperties{
			MaintenanceConfigurationID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Maintenance/maintenanceConfigurations/policy1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConfigurationAssignment = armmaintenance.ConfigurationAssignment{
	// 	Name: to.Ptr("workervmPolicy"),
	// 	Type: to.Ptr("Microsoft.Maintenance/configurationAssignments"),
	// 	ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Compute/virtualMachineScaleSets/smdtest1/virtualMachines/smdvm1/providers/Microsoft.Maintenance/configurationAssignments/workervmPolicy"),
	// 	Properties: &armmaintenance.ConfigurationAssignmentProperties{
	// 		MaintenanceConfigurationID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Maintenance/maintenanceConfigurations/policy1"),
	// 		ResourceID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Compute/virtualMachineScaleSets/smdtest1/virtualMachines/smdvm1"),
	// 	},
	// }
}
