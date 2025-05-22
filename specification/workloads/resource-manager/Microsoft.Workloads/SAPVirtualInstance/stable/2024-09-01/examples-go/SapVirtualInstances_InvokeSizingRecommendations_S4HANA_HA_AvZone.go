package armworkloadssapvirtualinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
)

// Generated from example definition: 2024-09-01/SapVirtualInstances_InvokeSizingRecommendations_S4HANA_HA_AvZone.json
func ExampleSAPVirtualInstancesClient_GetSizingRecommendations_sapSizingRecommendationsForHaWithAvailabilityZone() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("8e17e36c-42e9-4cd5-a078-7b44883414e0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSAPVirtualInstancesClient().GetSizingRecommendations(ctx, "centralus", armworkloadssapvirtualinstance.SAPSizingRecommendationRequest{
		AppLocation:          to.Ptr("eastus"),
		Environment:          to.Ptr(armworkloadssapvirtualinstance.SAPEnvironmentTypeProd),
		SapProduct:           to.Ptr(armworkloadssapvirtualinstance.SAPProductTypeS4HANA),
		DeploymentType:       to.Ptr(armworkloadssapvirtualinstance.SAPDeploymentTypeThreeTier),
		Saps:                 to.Ptr[int64](75000),
		DbMemory:             to.Ptr[int64](1024),
		DatabaseType:         to.Ptr(armworkloadssapvirtualinstance.SAPDatabaseTypeHANA),
		DbScaleMethod:        to.Ptr(armworkloadssapvirtualinstance.SAPDatabaseScaleMethodScaleUp),
		HighAvailabilityType: to.Ptr(armworkloadssapvirtualinstance.SAPHighAvailabilityTypeAvailabilityZone),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armworkloadssapvirtualinstance.SAPVirtualInstancesClientGetSizingRecommendationsResponse{
	// 	ThreeTierRecommendationResult: &armworkloadssapvirtualinstance.ThreeTierRecommendationResult{
	// 		DeploymentType: to.Ptr(armworkloadssapvirtualinstance.SAPDeploymentTypeThreeTier),
	// 		ApplicationServerVMSKU: to.Ptr("Standard_E8ds_v4"),
	// 		ApplicationServerInstanceCount: to.Ptr[int64](6),
	// 		CentralServerVMSKU: to.Ptr("Standard_E4ds_v4"),
	// 		CentralServerInstanceCount: to.Ptr[int64](2),
	// 		DbVMSKU: to.Ptr("Standard_M64s"),
	// 		DatabaseInstanceCount: to.Ptr[int64](2),
	// 	},
	// }
}
