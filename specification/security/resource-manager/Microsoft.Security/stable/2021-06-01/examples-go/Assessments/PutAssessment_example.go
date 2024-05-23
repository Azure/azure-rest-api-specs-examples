package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/Assessments/PutAssessment_example.json
func ExampleAssessmentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssessmentsClient().CreateOrUpdate(ctx, "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss2", "8bb8be0a-6010-4789-812f-e4d661c4ed0e", armsecurity.Assessment{
		Properties: &armsecurity.AssessmentProperties{
			ResourceDetails: &armsecurity.AzureResourceDetails{
				Source: to.Ptr(armsecurity.SourceAzure),
			},
			Status: &armsecurity.AssessmentStatus{
				Code: to.Ptr(armsecurity.AssessmentStatusCodeHealthy),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AssessmentResponse = armsecurity.AssessmentResponse{
	// 	Name: to.Ptr("8bb8be0a-6010-4789-812f-e4d661c4ed0e"),
	// 	Type: to.Ptr("Microsoft.Security/assessments"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss1/providers/Microsoft.Security/assessments/8bb8be0a-6010-4789-812f-e4d661c4ed0e"),
	// 	Properties: &armsecurity.AssessmentPropertiesResponse{
	// 		DisplayName: to.Ptr("Install internal agent on VM"),
	// 		ResourceDetails: &armsecurity.AzureResourceDetails{
	// 			Source: to.Ptr(armsecurity.SourceAzure),
	// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss1/providers/Microsoft.Security/assessments/8bb8be0a-6010-4789-812f-e4d661c4ed0e"),
	// 		},
	// 		Status: &armsecurity.AssessmentStatusResponse{
	// 			Code: to.Ptr(armsecurity.AssessmentStatusCodeHealthy),
	// 		},
	// 	},
	// }
}
