package armcommunication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication/v3"
)

// Generated from example definition: 2025-09-01/communicationServices/updateWithSystemAndUserIdentity.json
func ExampleServicesClient_Update_updateResourceToAddSystemAndUserManagedIdentities() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunication.NewClientFactory("11112222-3333-4444-5555-666677778888", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().Update(ctx, "MyResourceGroup", "MyCommunicationResource", armcommunication.ServiceResourceUpdate{
		Identity: &armcommunication.ManagedServiceIdentity{
			Type: to.Ptr(armcommunication.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
			UserAssignedIdentities: map[string]*armcommunication.UserAssignedIdentity{
				"/user/assigned/resource/id": {},
			},
		},
		Tags: map[string]*string{
			"newTag": to.Ptr("newVal"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcommunication.ServicesClientUpdateResponse{
	// 	ServiceResource: &armcommunication.ServiceResource{
	// 		Name: to.Ptr("MyCommunicationResource"),
	// 		Type: to.Ptr("Microsoft.Communication/CommunicationServices"),
	// 		ID: to.Ptr("/subscriptions/11112222-3333-4444-5555-666677778888/resourceGroups/MyResourceGroup/providers/Microsoft.Communication/CommunicationServices/MyCommunicationResource"),
	// 		Identity: &armcommunication.ManagedServiceIdentity{
	// 			Type: to.Ptr(armcommunication.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
	// 			PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 			TenantID: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 			UserAssignedIdentities: map[string]*armcommunication.UserAssignedIdentity{
	// 				"/user/assigned/resource/id": &armcommunication.UserAssignedIdentity{
	// 					ClientID: to.Ptr("11112222-3333-4444-5555-666677778888"),
	// 					PrincipalID: to.Ptr("11112222-3333-4444-5555-666677778888"),
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("Global"),
	// 		Properties: &armcommunication.ServiceProperties{
	// 			DataLocation: to.Ptr("United States"),
	// 			HostName: to.Ptr("mycommunicationresource.communications.azure.com"),
	// 			ProvisioningState: to.Ptr(armcommunication.CommunicationServicesProvisioningStateSucceeded),
	// 			Version: to.Ptr("0.2.0"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"newTag": to.Ptr("newVal"),
	// 		},
	// 	},
	// }
}
