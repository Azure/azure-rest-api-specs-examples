package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPSizingRecommendations_S4HANA_HA_AvSet.json
func ExampleClient_SAPSizingRecommendations_sapSizingRecommendationsS4HanaDistributedHaAvSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloads.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().SAPSizingRecommendations(ctx, "centralus", &armworkloads.ClientSAPSizingRecommendationsOptions{SAPSizingRecommendation: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armworkloads.ClientSAPSizingRecommendationsResponse{
	// 	                            SAPSizingRecommendationResultClassification: &armworkloads.ThreeTierRecommendationResult{
	// 		DeploymentType: to.Ptr(armworkloads.SAPDeploymentTypeThreeTier),
	// 		ApplicationServerInstanceCount: to.Ptr[int64](3),
	// 		ApplicationServerVMSKU: to.Ptr("Standard_E16ds_v4"),
	// 		CentralServerInstanceCount: to.Ptr[int64](2),
	// 		CentralServerVMSKU: to.Ptr("Standard_E8ds_v4"),
	// 		DatabaseInstanceCount: to.Ptr[int64](2),
	// 		DbVMSKU: to.Ptr("Standard_M64s"),
	// 	},
	// 	                        }
}
