package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d74afb775446d7f0bc1810fdc5a128c56289e854/specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/ClassifyProblemClassificationsForSubscription.json
func ExampleProblemClassificationsClient_ClassifyProblems() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProblemClassificationsClient().ClassifyProblems(ctx, "serviceId1", armsupport.ProblemClassificationsClassificationInput{
		IssueSummary: to.Ptr("Can not connect to Windows VM"),
		ResourceID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgname/providers/Microsoft.Compute/virtualMachines/vmname"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProblemClassificationsClassificationOutput = armsupport.ProblemClassificationsClassificationOutput{
	// 	ProblemClassificationResults: []*armsupport.ProblemClassificationsClassificationResult{
	// 		{
	// 			Description: to.Ptr("Problem classification description"),
	// 			ProblemClassificationID: to.Ptr("problemClassificationId1"),
	// 			ProblemID: to.Ptr("problemId1"),
	// 			RelatedService: &armsupport.ClassificationService{
	// 				DisplayName: to.Ptr("SQL Server in VM - Linux"),
	// 				ResourceTypes: []*string{
	// 					to.Ptr("MICROSOFT.CLASSICCOMPUTE/VIRTUALMACHINES"),
	// 					to.Ptr("MICROSOFT.COMPUTE/VIRTUALMACHINES")},
	// 					ServiceID: to.Ptr("/providers/Microsoft.Support/services/40ef020e-8ae7-8d57-b538-9153c47cee69"),
	// 				},
	// 				ServiceID: to.Ptr("serviceId1"),
	// 				Title: to.Ptr("Problem classification title"),
	// 		}},
	// 	}
}
