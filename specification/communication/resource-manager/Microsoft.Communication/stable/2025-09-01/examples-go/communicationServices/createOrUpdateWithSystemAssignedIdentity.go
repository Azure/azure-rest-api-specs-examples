package armcommunication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7832c9e47b8998a1c994b98345eea24dbc2ac5b8/specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/communicationServices/createOrUpdateWithSystemAssignedIdentity.json
func ExampleServicesClient_BeginCreateOrUpdate_createOrUpdateResourceWithManagedIdentity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServicesClient().BeginCreateOrUpdate(ctx, "MyResourceGroup", "MyCommunicationResource", armcommunication.ServiceResource{
		Location: to.Ptr("Global"),
		Identity: &armcommunication.ManagedServiceIdentity{
			Type: to.Ptr(armcommunication.ManagedServiceIdentityTypeSystemAssigned),
		},
		Properties: &armcommunication.ServiceProperties{
			DataLocation: to.Ptr("United States"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceResource = armcommunication.ServiceResource{
	// 	Name: to.Ptr("MyCommunicationResource"),
	// 	Type: to.Ptr("Microsoft.Communication/CommunicationServices"),
	// 	ID: to.Ptr("/subscriptions/11112222-3333-4444-5555-666677778888/resourceGroups/MyResourceGroup/providers/Microsoft.Communication/CommunicationServices/MyCommunicationResource"),
	// 	Location: to.Ptr("Global"),
	// 	Identity: &armcommunication.ManagedServiceIdentity{
	// 		Type: to.Ptr(armcommunication.ManagedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 		TenantID: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 	},
	// 	Properties: &armcommunication.ServiceProperties{
	// 		DataLocation: to.Ptr("United States"),
	// 		HostName: to.Ptr("mycommunicationresource.communications.azure.com"),
	// 		ProvisioningState: to.Ptr(armcommunication.CommunicationServicesProvisioningStateSucceeded),
	// 	},
	// }
}
