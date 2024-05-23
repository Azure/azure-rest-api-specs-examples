package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/RegulatoryCompliance/getRegulatoryComplianceStandard_example.json
func ExampleRegulatoryComplianceStandardsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRegulatoryComplianceStandardsClient().Get(ctx, "PCI-DSS-3.2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RegulatoryComplianceStandard = armsecurity.RegulatoryComplianceStandard{
	// 	Name: to.Ptr("PCI-DSS-3.2"),
	// 	Type: to.Ptr("Microsoft.Security/regulatoryComplianceStandard"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/regulatoryComplianceStandards/PCI-DSS-3.2"),
	// 	Properties: &armsecurity.RegulatoryComplianceStandardProperties{
	// 		FailedControls: to.Ptr[int32](4),
	// 		PassedControls: to.Ptr[int32](7),
	// 		SkippedControls: to.Ptr[int32](0),
	// 		State: to.Ptr(armsecurity.StateFailed),
	// 		UnsupportedControls: to.Ptr[int32](0),
	// 	},
	// }
}
