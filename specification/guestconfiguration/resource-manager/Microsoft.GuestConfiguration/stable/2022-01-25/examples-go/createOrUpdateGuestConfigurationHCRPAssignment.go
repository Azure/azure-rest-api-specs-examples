package armguestconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2022-01-25/examples/createOrUpdateGuestConfigurationHCRPAssignment.json
func ExampleHCRPAssignmentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armguestconfiguration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewHCRPAssignmentsClient().CreateOrUpdate(ctx, "NotInstalledApplicationForWindows", "myResourceGroupName", "myMachineName", armguestconfiguration.Assignment{
		Name:     to.Ptr("NotInstalledApplicationForWindows"),
		Location: to.Ptr("westcentralus"),
		Properties: &armguestconfiguration.AssignmentProperties{
			Context: to.Ptr("Azure policy"),
			GuestConfiguration: &armguestconfiguration.Navigation{
				Name:           to.Ptr("NotInstalledApplicationForWindows"),
				AssignmentType: to.Ptr(armguestconfiguration.AssignmentTypeApplyAndAutoCorrect),
				ConfigurationParameter: []*armguestconfiguration.ConfigurationParameter{
					{
						Name:  to.Ptr("[InstalledApplication]NotInstalledApplicationResource1;Name"),
						Value: to.Ptr("NotePad,sql"),
					}},
				ContentHash: to.Ptr("123contenthash"),
				ContentURI:  to.Ptr("https://thisisfake/pacakge"),
				Version:     to.Ptr("1.*"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Assignment = armguestconfiguration.Assignment{
	// 	Name: to.Ptr("NotInstalledApplicationForWindows"),
	// 	ID: to.Ptr("/subscriptions/mysubscriptionid/resourceGroups/myResourceGroupName/providers/HybridRP.Compute/virtualMachines/myvm/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/NotInstalledApplicationForWindows"),
	// 	Location: to.Ptr("westcentralus"),
	// 	Properties: &armguestconfiguration.AssignmentProperties{
	// 		AssignmentHash: to.Ptr("abcdr453g"),
	// 		ComplianceStatus: to.Ptr(armguestconfiguration.ComplianceStatusPending),
	// 		Context: to.Ptr("Azure policy"),
	// 		GuestConfiguration: &armguestconfiguration.Navigation{
	// 			Name: to.Ptr("NotInstalledApplicationForWindows"),
	// 			ConfigurationParameter: []*armguestconfiguration.ConfigurationParameter{
	// 				{
	// 					Name: to.Ptr("[InstalledApplication]NotInstalledApplicationResource1;Name"),
	// 					Value: to.Ptr("NotePad,sql"),
	// 			}},
	// 			Version: to.Ptr("1.0.0.3"),
	// 		},
	// 		LatestReportID: to.Ptr("a2a64e5d-a1a9-4344-a866-fb9e1541f723"),
	// 		ProvisioningState: to.Ptr(armguestconfiguration.ProvisioningStateSucceeded),
	// 	},
	// }
}
