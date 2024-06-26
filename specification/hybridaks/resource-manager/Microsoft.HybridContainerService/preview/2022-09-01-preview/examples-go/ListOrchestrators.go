package armhybridcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/ListOrchestrators.json
func ExampleClient_ListOrchestrators() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().ListOrchestrators(ctx, "subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OrchestratorVersionProfileListResult = armhybridcontainerservice.OrchestratorVersionProfileListResult{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("microsoft.hybridcontainerservice/customLocations/orchestrators"),
	// 	ID: to.Ptr("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation/providers/Microsoft.HybridContainerService/orchestrators"),
	// 	Orchestrators: []*armhybridcontainerservice.OrchestratorVersionProfile{
	// 		{
	// 			OrchestratorType: to.Ptr("Kubernetes"),
	// 			OrchestratorVersion: to.Ptr("1.8.1"),
	// 			Upgrades: []*armhybridcontainerservice.OrchestratorProfile{
	// 				{
	// 					OrchestratorVersion: to.Ptr("1.8.4"),
	// 				},
	// 				{
	// 					OrchestratorVersion: to.Ptr("1.8.2"),
	// 			}},
	// 		},
	// 		{
	// 			OrchestratorType: to.Ptr("Kubernetes"),
	// 			OrchestratorVersion: to.Ptr("1.9.10"),
	// 			Upgrades: []*armhybridcontainerservice.OrchestratorProfile{
	// 				{
	// 					IsPreview: to.Ptr(false),
	// 					OrchestratorType: to.Ptr(""),
	// 					OrchestratorVersion: to.Ptr("1.9.11"),
	// 				},
	// 				{
	// 					IsPreview: to.Ptr(false),
	// 					OrchestratorType: to.Ptr(""),
	// 					OrchestratorVersion: to.Ptr("1.10.12"),
	// 				},
	// 				{
	// 					IsPreview: to.Ptr(false),
	// 					OrchestratorType: to.Ptr(""),
	// 					OrchestratorVersion: to.Ptr("1.10.13"),
	// 			}},
	// 	}},
	// }
}
