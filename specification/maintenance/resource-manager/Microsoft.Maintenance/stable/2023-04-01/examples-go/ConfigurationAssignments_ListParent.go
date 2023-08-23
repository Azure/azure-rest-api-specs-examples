package armmaintenance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2023-04-01/examples/ConfigurationAssignments_ListParent.json
func ExampleConfigurationAssignmentsClient_NewListParentPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaintenance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConfigurationAssignmentsClient().NewListParentPager("examplerg", "Microsoft.Compute", "virtualMachineScaleSets", "smdtest1", "virtualMachines", "smdtestvm1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ListConfigurationAssignmentsResult = armmaintenance.ListConfigurationAssignmentsResult{
		// 	Value: []*armmaintenance.ConfigurationAssignment{
		// 		{
		// 			Name: to.Ptr("workervmPolicy"),
		// 			Type: to.Ptr("Microsoft.Maintenance/configurationAssignments"),
		// 			ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Compute/virtualMachineScaleSets/smdtest1/virtualMachines/smdtestvm1/providers/Microsoft.Maintenance/configurationAssignments/workervmPolicy"),
		// 			Properties: &armmaintenance.ConfigurationAssignmentProperties{
		// 				MaintenanceConfigurationID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Maintenance/maintenanceConfigurations/policy1"),
		// 				ResourceID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Compute/virtualMachineScaleSets/smdtest1/virtualMachines/smdtestvm1"),
		// 			},
		// 	}},
		// }
	}
}
