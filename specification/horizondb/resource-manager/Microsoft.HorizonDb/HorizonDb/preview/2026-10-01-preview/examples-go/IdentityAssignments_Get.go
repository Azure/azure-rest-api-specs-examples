package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-10-01-preview/IdentityAssignments_Get.json
func ExampleIdentityAssignmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIdentityAssignmentsClient().Get(ctx, "example-resource-group", "example-cluster", "fabric-mirroring", "fabric-writer", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhorizondb.IdentityAssignmentsClientGetResponse{
	// 	IdentityAssignment: armhorizondb.IdentityAssignment{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster/changeDataCaptureDestinations/fabric-mirroring/identityAssignments/fabric-writer"),
	// 		Name: to.Ptr("fabric-writer"),
	// 		Type: to.Ptr("Microsoft.HorizonDb/clusters/changeDataCaptureDestinations/identityAssignments"),
	// 		ETag: to.Ptr("\"00000000-0000-0000-0000-000000000020\""),
	// 		Properties: &armhorizondb.ChangeDataCaptureIdentityAssignmentProperties{
	// 			IdentityResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/identity-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/fabric-writer"),
	// 			EffectiveState: to.Ptr(armhorizondb.ChangeDataCaptureIdentityAssignmentStateAssigned),
	// 			ActivationState: to.Ptr(armhorizondb.ChangeDataCaptureIdentityActivationStateActive),
	// 			PrincipalID: to.Ptr("11111111-2222-3333-4444-555555555555"),
	// 			TenantID: to.Ptr("66666666-7777-8888-9999-000000000000"),
	// 			ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
	// 			Validation: &armhorizondb.IdentityValidationStatus{
	// 				State: to.Ptr(armhorizondb.ChangeDataCaptureIdentityValidationStateSucceeded),
	// 				ValidatedAtUTC: to.Ptr(time.Date(2026, time.September, 15, 17, 30, 0, 0, time.UTC)),
	// 			},
	// 			Status: &armhorizondb.ChangeDataCaptureStatus{
	// 				Code: to.Ptr("Ready"),
	// 				Message: to.Ptr("The identity assignment is active and valid."),
	// 			},
	// 		},
	// 	},
	// }
}
