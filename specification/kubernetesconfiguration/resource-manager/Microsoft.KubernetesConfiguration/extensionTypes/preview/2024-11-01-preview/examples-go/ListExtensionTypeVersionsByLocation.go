package armextensiontypes_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armextensiontypes"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ba0c086df0ebe03a61579485c1c10de0d17804b2/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/extensionTypes/preview/2024-11-01-preview/examples/ListExtensionTypeVersionsByLocation.json
func ExampleClient_NewListVersionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armextensiontypes.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListVersionsPager("westus", "extensionType1", &armextensiontypes.ClientListVersionsOptions{ReleaseTrain: to.Ptr("stable"),
		ClusterType:  to.Ptr("connectedCluster"),
		MajorVersion: to.Ptr("2"),
		ShowLatest:   to.Ptr(true),
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ExtensionTypeVersionsList = armextensiontypes.ExtensionTypeVersionsList{
		// 	Value: []*armextensiontypes.ExtensionTypeVersionForReleaseTrain{
		// 		{
		// 			Name: to.Ptr("bbbb"),
		// 			Type: to.Ptr("cccc"),
		// 			ID: to.Ptr("aaaa"),
		// 			Properties: &armextensiontypes.ExtensionTypeVersionForReleaseTrainProperties{
		// 				SupportedClusterTypes: []*string{
		// 					to.Ptr("connectedCluster")},
		// 					UnsupportedKubernetesVersions: &armextensiontypes.ExtensionTypeVersionForReleaseTrainPropertiesUnsupportedKubernetesVersions{
		// 						Appliances: []*armextensiontypes.ExtensionTypeVersionUnsupportedKubernetesMatrixItem{
		// 							{
		// 								Distributions: []*string{
		// 									to.Ptr("AKS")},
		// 									UnsupportedVersions: []*string{
		// 										to.Ptr("1.18..1.21")},
		// 								}},
		// 								ConnectedCluster: []*armextensiontypes.ExtensionTypeVersionUnsupportedKubernetesMatrixItem{
		// 									{
		// 										Distributions: []*string{
		// 											to.Ptr("AKS")},
		// 											UnsupportedVersions: []*string{
		// 											},
		// 										},
		// 										{
		// 											Distributions: []*string{
		// 												to.Ptr("GKE")},
		// 												UnsupportedVersions: []*string{
		// 													to.Ptr(">1.22")},
		// 											}},
		// 											ManagedCluster: []*armextensiontypes.ExtensionTypeVersionUnsupportedKubernetesMatrixItem{
		// 												{
		// 													Distributions: []*string{
		// 														to.Ptr("AKS")},
		// 														UnsupportedVersions: []*string{
		// 														},
		// 												}},
		// 												ProvisionedCluster: []*armextensiontypes.ExtensionTypeVersionUnsupportedKubernetesMatrixItem{
		// 													{
		// 														Distributions: []*string{
		// 															to.Ptr("AKS")},
		// 															UnsupportedVersions: []*string{
		// 															},
		// 													}},
		// 												},
		// 												Version: to.Ptr("1.1.0"),
		// 											},
		// 										},
		// 										{
		// 											Name: to.Ptr("bbbb"),
		// 											Type: to.Ptr("cccc"),
		// 											ID: to.Ptr("aaaa"),
		// 											Properties: &armextensiontypes.ExtensionTypeVersionForReleaseTrainProperties{
		// 												SupportedClusterTypes: []*string{
		// 													to.Ptr("connectedCluster")},
		// 													UnsupportedKubernetesVersions: &armextensiontypes.ExtensionTypeVersionForReleaseTrainPropertiesUnsupportedKubernetesVersions{
		// 														Appliances: []*armextensiontypes.ExtensionTypeVersionUnsupportedKubernetesMatrixItem{
		// 															{
		// 																Distributions: []*string{
		// 																	to.Ptr("AKS")},
		// 																	UnsupportedVersions: []*string{
		// 																		to.Ptr("1.18..1.21")},
		// 																}},
		// 																ConnectedCluster: []*armextensiontypes.ExtensionTypeVersionUnsupportedKubernetesMatrixItem{
		// 																	{
		// 																		Distributions: []*string{
		// 																			to.Ptr("AKS")},
		// 																			UnsupportedVersions: []*string{
		// 																			},
		// 																		},
		// 																		{
		// 																			Distributions: []*string{
		// 																				to.Ptr("GKE")},
		// 																				UnsupportedVersions: []*string{
		// 																					to.Ptr(">1.22")},
		// 																			}},
		// 																			ManagedCluster: []*armextensiontypes.ExtensionTypeVersionUnsupportedKubernetesMatrixItem{
		// 																				{
		// 																					Distributions: []*string{
		// 																						to.Ptr("AKS")},
		// 																						UnsupportedVersions: []*string{
		// 																						},
		// 																				}},
		// 																				ProvisionedCluster: []*armextensiontypes.ExtensionTypeVersionUnsupportedKubernetesMatrixItem{
		// 																					{
		// 																						Distributions: []*string{
		// 																							to.Ptr("AKS")},
		// 																							UnsupportedVersions: []*string{
		// 																							},
		// 																					}},
		// 																				},
		// 																				Version: to.Ptr("1.2.0"),
		// 																			},
		// 																	}},
		// 																}
	}
}
