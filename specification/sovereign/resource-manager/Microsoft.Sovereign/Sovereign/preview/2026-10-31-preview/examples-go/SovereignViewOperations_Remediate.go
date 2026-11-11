package armsovereign_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sovereign/armsovereign"
)

// Generated from example definition: 2026-10-31-preview/SovereignViewOperations_Remediate.json
func ExampleViewOperationsClient_BeginRemediate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsovereign.NewClientFactory("22b62f96-c8d5-4131-97f8-f820a5a3c55f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewViewOperationsClient().BeginRemediate(ctx, "ExampleResourceGroup", "ExampleSV", armsovereign.ViewRemediateContent{
		DimensionType: to.Ptr(armsovereign.ComplianceDimensionDataResidency),
		ResidencyRemediationRequest: &armsovereign.ViewResidencyRemediationContent{
			Configuration: &armsovereign.ViewResidencyRemediationConfiguration{
				TargetRegion:            to.Ptr("swedencentral"),
				TargetResourceGroupName: to.Ptr("RemediatedResourceGroup"),
			},
			Resources: []*armsovereign.ViewResidencyRemediationResource{
				{
					ID: to.Ptr("/subscriptions/22b62f96-c8d5-4131-97f8-f820a5a3c55f/resourceGroups/ExampleResourceGroup/providers/Microsoft.Storage/storageAccounts/examplestorage"),
					Configuration: &armsovereign.ViewResidencyRemediationResourceConfiguration{
						SourceRegion:            to.Ptr("northeurope"),
						TargetResourceGroupName: to.Ptr("RemediatedStorageResourceGroup"),
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsovereign.ViewOperationsClientRemediateResponse{
	// 	ViewRemediateResult: armsovereign.ViewRemediateResult{
	// 		RemediationID: to.Ptr("2f30c9e8-1b4d-4b4f-a17d-54db32e132f4"),
	// 	},
	// }
}
