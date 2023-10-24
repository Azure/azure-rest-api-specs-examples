package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/stable/2020-09-01/examples/getDataPolicyManifest.json
func ExampleDataPolicyManifestsClient_GetByPolicyMode() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataPolicyManifestsClient().GetByPolicyMode(ctx, "Microsoft.KeyVault.Data", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataPolicyManifest = armpolicy.DataPolicyManifest{
	// 	Name: to.Ptr("Microsoft.KeyVault.Data"),
	// 	Type: to.Ptr("Microsoft.Authorization/dataPolicyManifests"),
	// 	ID: to.Ptr("/providers/Microsoft.Authorization/dataPolicyManifests/Microsoft.KeyVault.Data"),
	// 	Properties: &armpolicy.DataPolicyManifestProperties{
	// 		Effects: []*armpolicy.DataEffect{
	// 			{
	// 				Name: to.Ptr("Audit"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Deny"),
	// 		}},
	// 		FieldValues: []*string{
	// 			to.Ptr("type")},
	// 			IsBuiltInOnly: to.Ptr(true),
	// 			Namespaces: []*string{
	// 				to.Ptr("Microsoft.KeyVault")},
	// 				PolicyMode: to.Ptr("Microsoft.KeyVault.Data"),
	// 				ResourceFunctions: &armpolicy.DataManifestResourceFunctionsDefinition{
	// 					Custom: []*armpolicy.DataManifestCustomResourceFunctionDefinition{
	// 						{
	// 							Name: to.Ptr("vault"),
	// 							AllowCustomProperties: to.Ptr(false),
	// 							DefaultProperties: []*string{
	// 								to.Ptr("location"),
	// 								to.Ptr("tags")},
	// 								FullyQualifiedResourceType: to.Ptr("Microsoft.KeyVault/vaults"),
	// 						}},
	// 						Standard: []*string{
	// 							to.Ptr("subscription"),
	// 							to.Ptr("resourceGroup")},
	// 						},
	// 						ResourceTypeAliases: []*armpolicy.ResourceTypeAliases{
	// 							{
	// 								Aliases: []*armpolicy.Alias{
	// 									{
	// 										Name: to.Ptr("Microsoft.KeyVault.Data/vaults/certificates/keyProperties.keyType"),
	// 										Paths: []*armpolicy.AliasPath{
	// 											{
	// 												Path: to.Ptr("keyProperties.keyType"),
	// 												APIVersions: []*string{
	// 													to.Ptr("2019-01-01")},
	// 											}},
	// 									}},
	// 									ResourceType: to.Ptr("vaults/certificates"),
	// 							}},
	// 						},
	// 					}
}
