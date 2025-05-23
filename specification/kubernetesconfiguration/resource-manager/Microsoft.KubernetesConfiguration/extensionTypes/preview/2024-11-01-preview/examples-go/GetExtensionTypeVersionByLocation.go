package armextensiontypes_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armextensiontypes"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ba0c086df0ebe03a61579485c1c10de0d17804b2/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/extensionTypes/preview/2024-11-01-preview/examples/GetExtensionTypeVersionByLocation.json
func ExampleClient_GetVersion() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armextensiontypes.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().GetVersion(ctx, "westus", "extensionType1", "1.20.0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExtensionTypeVersionForReleaseTrain = armextensiontypes.ExtensionTypeVersionForReleaseTrain{
	// 	Name: to.Ptr("bbbb"),
	// 	Type: to.Ptr("cccc"),
	// 	ID: to.Ptr("aaaa"),
	// 	Properties: &armextensiontypes.ExtensionTypeVersionForReleaseTrainProperties{
	// 		SupportedClusterTypes: []*string{
	// 			to.Ptr("connectedCluster")},
	// 			UnsupportedKubernetesVersions: &armextensiontypes.ExtensionTypeVersionForReleaseTrainPropertiesUnsupportedKubernetesVersions{
	// 				Appliances: []*armextensiontypes.ExtensionTypeVersionUnsupportedKubernetesMatrixItem{
	// 					{
	// 						Distributions: []*string{
	// 							to.Ptr("AKS")},
	// 							UnsupportedVersions: []*string{
	// 								to.Ptr("1.18..1.21")},
	// 						}},
	// 						ConnectedCluster: []*armextensiontypes.ExtensionTypeVersionUnsupportedKubernetesMatrixItem{
	// 							{
	// 								Distributions: []*string{
	// 									to.Ptr("AKS")},
	// 									UnsupportedVersions: []*string{
	// 									},
	// 								},
	// 								{
	// 									Distributions: []*string{
	// 										to.Ptr("GKE")},
	// 										UnsupportedVersions: []*string{
	// 											to.Ptr(">1.22")},
	// 									}},
	// 									ManagedCluster: []*armextensiontypes.ExtensionTypeVersionUnsupportedKubernetesMatrixItem{
	// 										{
	// 											Distributions: []*string{
	// 												to.Ptr("AKS")},
	// 												UnsupportedVersions: []*string{
	// 												},
	// 										}},
	// 										ProvisionedCluster: []*armextensiontypes.ExtensionTypeVersionUnsupportedKubernetesMatrixItem{
	// 											{
	// 												Distributions: []*string{
	// 													to.Ptr("AKS")},
	// 													UnsupportedVersions: []*string{
	// 													},
	// 											}},
	// 										},
	// 										Version: to.Ptr("1.1.0"),
	// 									},
	// 								}
}
