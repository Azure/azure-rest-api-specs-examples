package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-07-01-preview/dataConnectors/CheckRequirementsMicrosoftThreatIntelligence.json
func ExampleDataConnectorsCheckRequirementsClient_Post_checkRequirementsForMicrosoftThreatIntelligence() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorsCheckRequirementsClient().Post(ctx, "myRg", "myWorkspace", &armsecurityinsights.MSTICheckRequirements{
		Kind: to.Ptr(armsecurityinsights.DataConnectorKindMicrosoftThreatIntelligence),
		Properties: &armsecurityinsights.MSTICheckRequirementsProperties{
			TenantID: to.Ptr("06b3ccb8-1384-4bcc-aec7-852f6d57161b"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.DataConnectorsCheckRequirementsClientPostResponse{
	// 	DataConnectorRequirementsState: armsecurityinsights.DataConnectorRequirementsState{
	// 		AuthorizationState: to.Ptr(armsecurityinsights.DataConnectorAuthorizationStateValid),
	// 		LicenseState: to.Ptr(armsecurityinsights.DataConnectorLicenseStateValid),
	// 	},
	// }
}
