package armadvisor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/CreateConfiguration.json
func ExampleConfigurationsClient_CreateInSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armadvisor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConfigurationsClient().CreateInSubscription(ctx, armadvisor.ConfigurationNameDefault, armadvisor.ConfigData{
		Properties: &armadvisor.ConfigDataProperties{
			Digests: []*armadvisor.DigestConfig{
				{
					Name:                  to.Ptr("digestConfigName"),
					ActionGroupResourceID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup/providers/microsoft.insights/actionGroups/actionGroupName"),
					Categories: []*armadvisor.Category{
						to.Ptr(armadvisor.CategoryHighAvailability),
						to.Ptr(armadvisor.CategorySecurity),
						to.Ptr(armadvisor.CategoryPerformance),
						to.Ptr(armadvisor.CategoryCost),
						to.Ptr(armadvisor.CategoryOperationalExcellence)},
					Frequency: to.Ptr[int32](30),
					State:     to.Ptr(armadvisor.DigestConfigStateActive),
					Language:  to.Ptr("en"),
				}},
			Exclude:         to.Ptr(true),
			LowCPUThreshold: to.Ptr(armadvisor.CPUThresholdFive),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConfigData = armadvisor.ConfigData{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Advisor/configurations"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup/providers/Microsoft.Advisor/configurations/default"),
	// 	Properties: &armadvisor.ConfigDataProperties{
	// 		Digests: []*armadvisor.DigestConfig{
	// 			{
	// 				Name: to.Ptr("digestConfigName"),
	// 				ActionGroupResourceID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup/providers/microsoft.insights/actionGroups/actionGroupName"),
	// 				Categories: []*armadvisor.Category{
	// 					to.Ptr(armadvisor.CategoryHighAvailability),
	// 					to.Ptr(armadvisor.CategorySecurity),
	// 					to.Ptr(armadvisor.CategoryPerformance),
	// 					to.Ptr(armadvisor.CategoryCost),
	// 					to.Ptr(armadvisor.CategoryOperationalExcellence)},
	// 					Frequency: to.Ptr[int32](30),
	// 					State: to.Ptr(armadvisor.DigestConfigStateActive),
	// 					Language: to.Ptr("en"),
	// 			}},
	// 			Exclude: to.Ptr(true),
	// 			LowCPUThreshold: to.Ptr(armadvisor.CPUThresholdFive),
	// 		},
	// 	}
}
