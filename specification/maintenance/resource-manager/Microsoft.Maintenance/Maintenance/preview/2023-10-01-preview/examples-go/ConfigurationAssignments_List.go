package armmaintenance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"
)

// Generated from example definition: 2023-10-01-preview/ConfigurationAssignments_List.json
func ExampleConfigurationAssignmentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaintenance.NewClientFactory("5b4b650e-28b9-4790-b3ab-ddbd88d727c4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConfigurationAssignmentsClient().NewListPager("examplerg", "Microsoft.Compute", "virtualMachineScaleSets", "smdtest1", nil)
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
		// page = armmaintenance.ConfigurationAssignmentsClientListResponse{
		// 	ListConfigurationAssignmentsResult: armmaintenance.ListConfigurationAssignmentsResult{
		// 		Value: []*armmaintenance.ConfigurationAssignment{
		// 			{
		// 				ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Compute/virtualMachineScaleSets/smdtest1/providers/Microsoft.Maintenance/configurationAssignments/workervmConfiguration"),
		// 				Name: to.Ptr("workervmConfiguration"),
		// 				Type: to.Ptr("Microsoft.Maintenance/configurationAssignments"),
		// 				Properties: &armmaintenance.ConfigurationAssignmentProperties{
		// 					MaintenanceConfigurationID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Maintenance/maintenanceConfigurations/configuration1"),
		// 					ResourceID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Compute/virtualMachineScaleSets/smdtest1"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
