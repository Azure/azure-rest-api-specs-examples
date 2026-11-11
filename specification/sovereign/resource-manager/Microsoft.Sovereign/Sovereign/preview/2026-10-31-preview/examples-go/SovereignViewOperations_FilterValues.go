package armsovereign_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sovereign/armsovereign"
)

// Generated from example definition: 2026-10-31-preview/SovereignViewOperations_FilterValues.json
func ExampleViewOperationsClient_FilterValues() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsovereign.NewClientFactory("22b62f96-c8d5-4131-97f8-f820a5a3c55f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewViewOperationsClient().FilterValues(ctx, "ExampleResourceGroup", "ExampleSV", armsovereign.ViewFilterValuesContent{
		RequestedFilter: to.Ptr(armsovereign.ViewFilterFieldResourceType),
		Filters: []*armsovereign.ViewComplianceDetailsFilters{
			{
				Field:    to.Ptr(armsovereign.ViewFilterFieldRegion),
				Operator: to.Ptr(armsovereign.ViewComplianceDetailsFilterOperatorsIn),
				Value: []*string{
					to.Ptr("northeurope"),
				},
			},
			{
				Field:    to.Ptr(armsovereign.ViewFilterFieldParentResourceName),
				Operator: to.Ptr(armsovereign.ViewComplianceDetailsFilterOperatorsIn),
				Value: []*string{
					to.Ptr("exampleParentVm"),
				},
			},
			{
				Field:    to.Ptr(armsovereign.ViewFilterFieldBusinessContinuity),
				Operator: to.Ptr(armsovereign.ViewComplianceDetailsFilterOperatorsIn),
				Value: []*string{
					to.Ptr("Constrained"),
				},
			},
		},
		SearchTerm: to.Ptr("Microsoft.Compute"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsovereign.ViewOperationsClientFilterValuesResponse{
	// 	ViewFilterValuesResult: armsovereign.ViewFilterValuesResult{
	// 		DataState: to.Ptr(armsovereign.DataStateProcessed),
	// 		RequestedFilter: to.Ptr(armsovereign.ViewFilterFieldResourceType),
	// 		TotalCount: to.Ptr[int32](3),
	// 		Values: []*armsovereign.ViewFilterValueItem{
	// 			{
	// 				Value: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 				DisplayName: to.Ptr("Virtual Machines"),
	// 				Count: to.Ptr[int32](5),
	// 			},
	// 			{
	// 				Value: to.Ptr("Microsoft.Compute/disks"),
	// 				DisplayName: to.Ptr("Disks"),
	// 				Count: to.Ptr[int32](3),
	// 			},
	// 			{
	// 				Value: to.Ptr("Microsoft.Compute/virtualMachineScaleSets"),
	// 				DisplayName: to.Ptr("Virtual Machine Scale Sets"),
	// 				Count: to.Ptr[int32](2),
	// 			},
	// 		},
	// 	},
	// }
}
