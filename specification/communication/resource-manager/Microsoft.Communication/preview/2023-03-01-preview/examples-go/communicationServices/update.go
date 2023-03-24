package armcommunication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ddcde53422ff186f5b69fb32338ecac3d11c3bca/specification/communication/resource-manager/Microsoft.Communication/preview/2023-03-01-preview/examples/communicationServices/update.json
func ExampleServicesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().Update(ctx, "MyResourceGroup", "MyCommunicationResource", armcommunication.ServiceResourceUpdate{
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
	// res.ServiceResource = armcommunication.ServiceResource{
	// 	Name: to.Ptr("MyCommunicationResource"),
	// 	Type: to.Ptr("Microsoft.Communication/CommunicationServices"),
	// 	ID: to.Ptr("/subscriptions/11112222-3333-4444-5555-666677778888/resourceGroups/MyResourceGroup/providers/Microsoft.Communication/CommunicationServices/MyCommunicationResource"),
	// 	Location: to.Ptr("Global"),
	// 	Tags: map[string]*string{
	// 		"newTag": to.Ptr("newVal"),
	// 	},
	// 	Properties: &armcommunication.ServiceProperties{
	// 		DataLocation: to.Ptr("United States"),
	// 		HostName: to.Ptr("mycommunicationresource.communications.azure.com"),
	// 		ProvisioningState: to.Ptr(armcommunication.CommunicationServicesProvisioningStateSucceeded),
	// 		Version: to.Ptr("0.2.0"),
	// 	},
	// }
}
