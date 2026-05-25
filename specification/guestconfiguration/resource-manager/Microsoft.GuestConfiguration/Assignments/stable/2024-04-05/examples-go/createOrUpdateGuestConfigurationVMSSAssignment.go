package armguestconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration/v2"
)

// Generated from example definition: 2024-04-05/createOrUpdateGuestConfigurationVMSSAssignment.json
func ExampleAssignmentsVMSSClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armguestconfiguration.NewClientFactory("mySubscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssignmentsVMSSClient().CreateOrUpdate(ctx, "myResourceGroupName", "myVMSSName", "NotInstalledApplicationForWindows", armguestconfiguration.Assignment{
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
					},
				},
				ContentHash:            to.Ptr("123contenthash"),
				ContentManagedIdentity: to.Ptr("test_identity"),
				ContentURI:             to.Ptr("https://thisisfake/pacakge"),
				Version:                to.Ptr("1.0.0.3"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armguestconfiguration.AssignmentsVMSSClientCreateOrUpdateResponse{
	// 	Assignment: armguestconfiguration.Assignment{
	// 		Name: to.Ptr("NotInstalledApplicationForWindows"),
	// 		ID: to.Ptr("/subscriptions/mysubscriptionid/resourceGroups/myResourceGroupName/providers/Microsoft.Compute/virtualmachinescalesets/myvmssname/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/NotInstalledApplicationForWindows"),
	// 		Location: to.Ptr("westcentralus"),
	// 		Properties: &armguestconfiguration.AssignmentProperties{
	// 			AssignmentHash: nil,
	// 			ComplianceStatus: to.Ptr(armguestconfiguration.ComplianceStatusPending),
	// 			Context: to.Ptr("Azure policy"),
	// 			GuestConfiguration: &armguestconfiguration.Navigation{
	// 				Name: to.Ptr("NotInstalledApplicationForWindows"),
	// 				AssignmentSource: to.Ptr("AzurePolicy"),
	// 				ConfigurationParameter: []*armguestconfiguration.ConfigurationParameter{
	// 					{
	// 						Name: to.Ptr("[InstalledApplication]NotInstalledApplicationResource1;Name"),
	// 						Value: to.Ptr("NotePad,sql"),
	// 					},
	// 				},
	// 				ContentManagedIdentity: to.Ptr("test_identity"),
	// 				ContentType: nil,
	// 				Version: to.Ptr("1.0.0.3"),
	// 			},
	// 			LastComplianceStatusChecked: nil,
	// 			LatestReportID: nil,
	// 			ProvisioningState: to.Ptr(armguestconfiguration.ProvisioningStateSucceeded),
	// 			ResourceType: nil,
	// 		},
	// 	},
	// }
}
