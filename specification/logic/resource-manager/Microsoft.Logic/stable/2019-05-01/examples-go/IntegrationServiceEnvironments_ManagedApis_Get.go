package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_ManagedApis_Get.json
func ExampleIntegrationServiceEnvironmentManagedApisClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIntegrationServiceEnvironmentManagedApisClient().Get(ctx, "testResourceGroup", "testIntegrationServiceEnvironment", "servicebus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IntegrationServiceEnvironmentManagedAPI = armlogic.IntegrationServiceEnvironmentManagedAPI{
	// 	Name: to.Ptr("servicebus"),
	// 	Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments/managedApis"),
	// 	ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment/managedApis/servicebus"),
	// 	Location: to.Ptr("northeurope"),
	// 	Properties: &armlogic.IntegrationServiceEnvironmentManagedAPIProperties{
	// 		Name: to.Ptr("servicebus"),
	// 		Capabilities: []*string{
	// 			to.Ptr("actions")},
	// 			ConnectionParameters: map[string]any{
	// 				"connectionString": map[string]any{
	// 					"type": "securestring",
	// 					"uiDefinition":map[string]any{
	// 						"description": "Azure Service Bus Connection String",
	// 						"constraints":map[string]any{
	// 							"required": "true",
	// 						},
	// 						"displayName": "Connection String",
	// 						"tooltip": "Provide Azure Service Bus Connection String",
	// 					},
	// 				},
	// 			},
	// 			GeneralInformation: &armlogic.APIResourceGeneralInformation{
	// 				Description: to.Ptr("Connect to Azure Service Bus to send and receive messages. You can perform actions such as send to queue, send to topic, receive from queue, receive from subscription, etc."),
	// 				DisplayName: to.Ptr("Service Bus"),
	// 				IconURL: to.Ptr("https://cpgeneralstore.blob.core.windows.net/officialicons/servicebus/icon_1.0.1206.1574.png"),
	// 				ReleaseTag: to.Ptr("Production"),
	// 				Tier: to.Ptr(armlogic.APITierStandard),
	// 			},
	// 			IntegrationServiceEnvironment: &armlogic.ResourceReference{
	// 				Name: to.Ptr("testIntegrationServiceEnvironment"),
	// 				Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments"),
	// 				ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment"),
	// 			},
	// 			Metadata: &armlogic.APIResourceMetadata{
	// 				BrandColor: to.Ptr("#c4d5ff"),
	// 				Source: to.Ptr("marketplace"),
	// 			},
	// 			ProvisioningState: to.Ptr(armlogic.WorkflowProvisioningStateSucceeded),
	// 			RuntimeUrls: []*string{
	// 				to.Ptr("https://flow-weiroa6odksti-db-apim-runtime.northeurope.environments.microsoftazurelogicapps.net/apim/servicebus")},
	// 			},
	// 		}
}
