package armadvisor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor/v2"
)

// Generated from example definition: 2025-05-01-preview/CreateConfiguration_CreateInSubscription.json
func ExampleConfigurationsClient_CreateInSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armadvisor.NewClientFactory("58c3f667-7a62-4bfd-a658-846493e9a493", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConfigurationsClient().CreateInSubscription(ctx, armadvisor.ConfigurationNameDefault, armadvisor.ConfigData{
		Properties: &armadvisor.ConfigDataProperties{
			LowCPUThreshold: to.Ptr(armadvisor.CPUThresholdFive),
			Duration:        to.Ptr(armadvisor.DurationSeven),
			Exclude:         to.Ptr(true),
			Digests: []*armadvisor.DigestConfig{
				{
					Name:                  to.Ptr("digestConfigName"),
					ActionGroupResourceID: to.Ptr("/subscriptions/58c3f667-7a62-4bfd-a658-846493e9a493/resourceGroups/resourceGroup/providers/microsoft.insights/actionGroups/actionGroupName"),
					Frequency:             to.Ptr[int32](30),
					Categories: []*armadvisor.Category{
						to.Ptr(armadvisor.CategoryHighAvailability),
						to.Ptr(armadvisor.CategorySecurity),
						to.Ptr(armadvisor.CategoryPerformance),
						to.Ptr(armadvisor.CategoryCost),
						to.Ptr(armadvisor.CategoryOperationalExcellence),
					},
					Language: to.Ptr("en"),
					State:    to.Ptr(armadvisor.DigestConfigStateActive),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armadvisor.ConfigurationsClientCreateInSubscriptionResponse{
	// 	ConfigData: armadvisor.ConfigData{
	// 		ID: to.Ptr("/subscriptions/58c3f667-7a62-4bfd-a658-846493e9a493/resourceGroups/resourceGroup/providers/Microsoft.Advisor/configurations/default"),
	// 		Type: to.Ptr("Microsoft.Advisor/configurations"),
	// 		Name: to.Ptr("default"),
	// 		Properties: &armadvisor.ConfigDataProperties{
	// 			LowCPUThreshold: to.Ptr(armadvisor.CPUThresholdFive),
	// 			Duration: to.Ptr(armadvisor.DurationSeven),
	// 			Exclude: to.Ptr(true),
	// 			Digests: []*armadvisor.DigestConfig{
	// 				{
	// 					Name: to.Ptr("digestConfigName"),
	// 					ActionGroupResourceID: to.Ptr("/subscriptions/58c3f667-7a62-4bfd-a658-846493e9a493/resourceGroups/resourceGroup/providers/microsoft.insights/actionGroups/actionGroupName"),
	// 					Frequency: to.Ptr[int32](30),
	// 					Categories: []*armadvisor.Category{
	// 						to.Ptr(armadvisor.CategoryHighAvailability),
	// 						to.Ptr(armadvisor.CategorySecurity),
	// 						to.Ptr(armadvisor.CategoryPerformance),
	// 						to.Ptr(armadvisor.CategoryCost),
	// 						to.Ptr(armadvisor.CategoryOperationalExcellence),
	// 					},
	// 					Language: to.Ptr("en"),
	// 					State: to.Ptr(armadvisor.DigestConfigStateActive),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
