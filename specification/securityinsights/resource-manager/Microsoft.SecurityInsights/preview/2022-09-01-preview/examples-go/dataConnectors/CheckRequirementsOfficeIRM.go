package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/dataConnectors/CheckRequirementsOfficeIRM.json
func ExampleDataConnectorsCheckRequirementsClient_Post_checkRequirementsForOfficeIrm() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorsCheckRequirementsClient().Post(ctx, "myRg", "myWorkspace", &armsecurityinsights.OfficeIRMCheckRequirements{
		Kind: to.Ptr(armsecurityinsights.DataConnectorKindOfficeIRM),
		Properties: &armsecurityinsights.OfficeIRMCheckRequirementsProperties{
			TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataConnectorRequirementsState = armsecurityinsights.DataConnectorRequirementsState{
	// 	AuthorizationState: to.Ptr(armsecurityinsights.DataConnectorAuthorizationStateValid),
	// 	LicenseState: to.Ptr(armsecurityinsights.DataConnectorLicenseStateValid),
	// }
}
