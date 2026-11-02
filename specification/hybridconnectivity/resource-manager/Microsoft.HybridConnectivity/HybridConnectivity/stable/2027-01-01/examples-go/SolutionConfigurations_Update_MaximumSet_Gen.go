package armhybridconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
)

// Generated from example definition: 2027-01-01/SolutionConfigurations_Update_MaximumSet_Gen.json
func ExampleSolutionConfigurationsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSolutionConfigurationsClient().Update(ctx, "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine", "abc", armhybridconnectivity.SolutionConfigurationUpdate{
		Properties: &armhybridconnectivity.SolutionConfigurationPropertiesUpdate{
			SolutionType:     to.Ptr("myzljlstvmgkp"),
			SolutionSettings: &armhybridconnectivity.SolutionSettings{},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhybridconnectivity.SolutionConfigurationsClientUpdateResponse{
	// 	SolutionConfiguration: armhybridconnectivity.SolutionConfiguration{
	// 		Properties: &armhybridconnectivity.SolutionConfigurationProperties{
	// 			SolutionType: to.Ptr("nmtqllkyohwtsthxaimsye"),
	// 			SolutionSettings: &armhybridconnectivity.SolutionSettings{
	// 			},
	// 			ProvisioningState: to.Ptr(armhybridconnectivity.ResourceProvisioningStateSucceeded),
	// 			Status: to.Ptr(armhybridconnectivity.SolutionConfigurationStatus("AuthorizationPending")),
	// 			StatusDetails: to.Ptr("Aws authorization validation pending in Aws account"),
	// 			LastSyncTime: to.Ptr(time.Date(2024, time.January, 20, 3, 24, 15, 820000000, time.UTC)),
	// 		},
	// 		ID: to.Ptr("/subscriptions/testSubcrptions/resourceGroups/testResourceGroup/providers/Microsoft.HybridConnectivity/SolutionConfigurations/solutionconfigurationtest"),
	// 		Name: to.Ptr("solutionconfigurationtest"),
	// 		Type: to.Ptr("microsoft.hybridconnectivity/solutionconfigurations"),
	// 		SystemData: &armhybridconnectivity.SystemData{
	// 			CreatedBy: to.Ptr("rpxzkcrobprrdvuoqxz"),
	// 			CreatedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2024, time.January, 18, 22, 52, 7, 890000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("jidegyskxi"),
	// 			LastModifiedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2024, time.January, 18, 22, 52, 7, 890000000, time.UTC)),
	// 		},
	// 	},
	// }
}
