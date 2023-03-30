package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_ManagedApis_Put.json
func ExampleIntegrationServiceEnvironmentManagedApisClient_BeginPut() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIntegrationServiceEnvironmentManagedApisClient().BeginPut(ctx, "testResourceGroup", "testIntegrationServiceEnvironment", "servicebus", armlogic.IntegrationServiceEnvironmentManagedAPI{
		Location:   to.Ptr("brazilsouth"),
		Properties: &armlogic.IntegrationServiceEnvironmentManagedAPIProperties{},
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
	// res.IntegrationServiceEnvironmentManagedAPI = armlogic.IntegrationServiceEnvironmentManagedAPI{
	// 	Name: to.Ptr("servicebus"),
	// 	Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments/managedApis"),
	// 	ID: to.Ptr("/subscriptions/80d4fe69-c95b-4dd2-a938-9250f1c8ab03/resourceGroups/rohithah-ise/providers/Microsoft.Logic/integrationServiceEnvironments/tes-ise-ga/managedApis/servicebus"),
	// 	Location: to.Ptr("brazilsouth"),
	// 	Properties: &armlogic.IntegrationServiceEnvironmentManagedAPIProperties{
	// 		Category: to.Ptr(armlogic.APITierStandard),
	// 		GeneralInformation: &armlogic.APIResourceGeneralInformation{
	// 			Description: to.Ptr("Connect to Azure Service Bus to send and receive messages. You can perform actions such as send to queue, send to topic, receive from queue, receive from subscription, etc."),
	// 			DisplayName: to.Ptr("Service Bus"),
	// 			IconURL: to.Ptr("https://connectoricons-df.azureedge.net/servicebus/icon_1.0.1223.1623.png"),
	// 			ReleaseTag: to.Ptr("Production"),
	// 			Tier: to.Ptr(armlogic.APITierStandard),
	// 		},
	// 		IntegrationServiceEnvironment: &armlogic.ResourceReference{
	// 			Name: to.Ptr("tes-ise-ga"),
	// 			Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments"),
	// 			ID: to.Ptr("/subscriptions/80d4fe69-c95b-4dd2-a938-9250f1c8ab03/resourceGroups/rohithah-ise/providers/Microsoft.Logic/integrationServiceEnvironments/tes-ise-ga"),
	// 		},
	// 		ProvisioningState: to.Ptr(armlogic.WorkflowProvisioningStateSucceeded),
	// 	},
	// }
}
