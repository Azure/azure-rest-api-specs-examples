package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/ConfigurationGroupValueUpdateTags.json
func ExampleConfigurationGroupValuesClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConfigurationGroupValuesClient().UpdateTags(ctx, "rg1", "testConfigurationGroupValue", armhybridnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConfigurationGroupValue = armhybridnetwork.ConfigurationGroupValue{
	// 	Name: to.Ptr("testConfigurationGroupValue"),
	// 	Type: to.Ptr("Microsoft.HybridNetwork/configurationGroupValues"),
	// 	ID: to.Ptr("/subscriptions/subid/providers/Microsoft.HybridNetwork/configurationGroupValues/testConfigurationGroupValue"),
	// 	SystemData: &armhybridnetwork.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armhybridnetwork.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armhybridnetwork.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armhybridnetwork.ConfigurationValueWithoutSecrets{
	// 		ConfigurationGroupSchemaName: to.Ptr("testConfigurationGroupSchemaName"),
	// 		ConfigurationGroupSchemaOfferingLocation: to.Ptr("eastus"),
	// 		ConfigurationGroupSchemaResourceReference: &armhybridnetwork.OpenDeploymentResourceReference{
	// 			IDType: to.Ptr(armhybridnetwork.IDTypeOpen),
	// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/testRG/providers/microsoft.hybridnetwork/publishers/testPublisher/configurationGroupSchemas/testConfigurationGroupSchemaName"),
	// 		},
	// 		ConfigurationType: to.Ptr(armhybridnetwork.ConfigurationGroupValueConfigurationTypeOpen),
	// 		ProvisioningState: to.Ptr(armhybridnetwork.ProvisioningStateSucceeded),
	// 		PublisherName: to.Ptr("testPublisher"),
	// 		PublisherScope: to.Ptr(armhybridnetwork.PublisherScope("Public")),
	// 		ConfigurationValue: to.Ptr("{\"interconnect-groups\":{\"stripe-one\":{\"name\":\"Stripe one\",\"international-interconnects\":[\"france\",\"germany\"],\"domestic-interconnects\":[\"birmingham\",\"edinburgh\"]},\"stripe-two\":{\"name\":\"Stripe two\",\"international-interconnects\":[\"germany\",\"italy\"],\"domestic-interconnects\":[\"edinburgh\",\"london\"]}},\"interconnect-group-assignments\":{\"ssc-one\":{\"ssc\":\"SSC 1\",\"interconnects\":\"stripe-one\"},\"ssc-two\":{\"ssc\":\"SSC 2\",\"interconnects\":\"stripe-two\"}}}"),
	// 	},
	// }
}
