package armhybridconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
)

// Generated from example definition: 2024-12-01/SolutionConfigurations_CreateOrUpdate.json
func ExampleSolutionConfigurationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSolutionConfigurationsClient().CreateOrUpdate(ctx, "ymuj", "keebwujt", armhybridconnectivity.SolutionConfiguration{
		Properties: &armhybridconnectivity.SolutionConfigurationProperties{
			SolutionType:     to.Ptr("nmtqllkyohwtsthxaimsye"),
			SolutionSettings: &armhybridconnectivity.SolutionSettings{},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhybridconnectivity.SolutionConfigurationsClientCreateOrUpdateResponse{
	// 	SolutionConfiguration: &armhybridconnectivity.SolutionConfiguration{
	// 		Properties: &armhybridconnectivity.SolutionConfigurationProperties{
	// 			SolutionType: to.Ptr("nmtqllkyohwtsthxaimsye"),
	// 			SolutionSettings: &armhybridconnectivity.SolutionSettings{
	// 			},
	// 			ProvisioningState: to.Ptr(armhybridconnectivity.ResourceProvisioningStateSucceeded),
	// 			Status: to.Ptr(armhybridconnectivity.SolutionConfigurationStatusNew),
	// 			StatusDetails: to.Ptr("rqbrzildwecankrpukkbjjqrczxboz"),
	// 			LastSyncTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-20T03:24:15.820Z"); return t}()),
	// 		},
	// 		ID: to.Ptr("/subscriptions/testSubcrptions/resourceGroups/testResourceGroup/providers/Microsoft.HybridConnectivity/SolutionConfigurations/qpwubemzmootxmtlxaerir"),
	// 		Name: to.Ptr("qpwubemzmootxmtlxaerir"),
	// 		Type: to.Ptr("uknrk"),
	// 		SystemData: &armhybridconnectivity.SystemData{
	// 			CreatedBy: to.Ptr("rpxzkcrobprrdvuoqxz"),
	// 			CreatedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-18T22:52:07.890Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("jidegyskxi"),
	// 			LastModifiedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-18T22:52:07.890Z"); return t}()),
	// 		},
	// 	},
	// }
}
