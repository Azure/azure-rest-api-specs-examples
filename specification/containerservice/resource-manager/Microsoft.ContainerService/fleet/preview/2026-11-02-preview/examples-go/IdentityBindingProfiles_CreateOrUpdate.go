package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet/v3"
)

// Generated from example definition: 2026-11-02-preview/IdentityBindingProfiles_CreateOrUpdate.json
func ExampleIdentityBindingProfilesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIdentityBindingProfilesClient().BeginCreateOrUpdate(ctx, "rg1", "fleet1", "identitybinding1", armcontainerservicefleet.IdentityBindingProfile{
		Properties: &armcontainerservicefleet.IdentityBindingProfileProperties{
			MemberSelector: &armcontainerservicefleet.MemberSelector{
				ByLabel: to.Ptr("env=production"),
			},
			IdentityBindingProperties: &armcontainerservicefleet.IdentityBindingProperties{
				ManagedIdentity: &armcontainerservicefleet.IdentityBindingManagedIdentityProfile{
					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1"),
				},
				AllowedSubjects: []*armcontainerservicefleet.AllowedSubject{
					{
						NamespaceSelector: &armcontainerservicefleet.IdentityBindingLabelSelector{
							MatchLabels: []*string{
								to.Ptr("kubernetes.io/metadata.name=prod"),
							},
						},
						ServiceAccountSelector: &armcontainerservicefleet.IdentityBindingLabelSelector{
							MatchLabels: []*string{
								to.Ptr("app=payments"),
							},
						},
					},
				},
			},
		},
	}, &armcontainerservicefleet.IdentityBindingProfilesClientBeginCreateOrUpdateOptions{
		IfMatch: to.Ptr("uktvayathbu")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerservicefleet.IdentityBindingProfilesClientCreateOrUpdateResponse{
	// 	IdentityBindingProfile: armcontainerservicefleet.IdentityBindingProfile{
	// 		Properties: &armcontainerservicefleet.IdentityBindingProfileProperties{
	// 			ProvisioningState: to.Ptr(armcontainerservicefleet.IdentityBindingProfileProvisioningStateSucceeded),
	// 			MemberSelector: &armcontainerservicefleet.MemberSelector{
	// 				ByLabel: to.Ptr("env=production"),
	// 			},
	// 			IdentityBindingProperties: &armcontainerservicefleet.IdentityBindingProperties{
	// 				ManagedIdentity: &armcontainerservicefleet.IdentityBindingManagedIdentityProfile{
	// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1"),
	// 					ObjectID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				},
	// 				AllowedSubjects: []*armcontainerservicefleet.AllowedSubject{
	// 					{
	// 						NamespaceSelector: &armcontainerservicefleet.IdentityBindingLabelSelector{
	// 							MatchLabels: []*string{
	// 								to.Ptr("kubernetes.io/metadata.name=prod"),
	// 							},
	// 						},
	// 						ServiceAccountSelector: &armcontainerservicefleet.IdentityBindingLabelSelector{
	// 							MatchLabels: []*string{
	// 								to.Ptr("app=payments"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 				OIDCIssuer: &armcontainerservicefleet.IdentityBindingOIDCIssuerProfile{
	// 					OIDCIssuerURL: to.Ptr("https://oidc-endpoint"),
	// 				},
	// 			},
	// 			Status: &armcontainerservicefleet.IdentityBindingProfileStatus{
	// 				LastOperationID: to.Ptr("12345678-1234-1234-1234-123456789012"),
	// 			},
	// 		},
	// 		ETag: to.Ptr("\"EtagValue\""),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/fleets/fleet1/identityBindingProfiles/identitybinding1"),
	// 		Name: to.Ptr("identitybinding1"),
	// 		Type: to.Ptr("Microsoft.ContainerService/fleets/identityBindingProfiles"),
	// 		SystemData: &armcontainerservicefleet.SystemData{
	// 			CreatedBy: to.Ptr("someUser"),
	// 			CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2022, time.March, 23, 5, 40, 40, 657000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("someOtherUser"),
	// 			LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2022, time.March, 23, 5, 40, 40, 657000000, time.UTC)),
	// 		},
	// 	},
	// }
}
