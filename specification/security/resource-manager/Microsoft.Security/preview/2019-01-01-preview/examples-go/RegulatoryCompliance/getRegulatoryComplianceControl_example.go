package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/RegulatoryCompliance/getRegulatoryComplianceControl_example.json
func ExampleRegulatoryComplianceControlsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRegulatoryComplianceControlsClient().Get(ctx, "PCI-DSS-3.2", "1.1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RegulatoryComplianceControl = armsecurity.RegulatoryComplianceControl{
	// 	Name: to.Ptr("1.1"),
	// 	Type: to.Ptr("Microsoft.Security/regulatoryComplianceControl"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/regulatoryComplianceStandards/PCI-DSS-3.2/regulatoryComplianceControls/1.1"),
	// 	Properties: &armsecurity.RegulatoryComplianceControlProperties{
	// 		Description: to.Ptr("Common Criteria Related to Organization and Management"),
	// 		FailedAssessments: to.Ptr[int32](4),
	// 		PassedAssessments: to.Ptr[int32](7),
	// 		SkippedAssessments: to.Ptr[int32](0),
	// 		State: to.Ptr(armsecurity.StateFailed),
	// 	},
	// }
}
