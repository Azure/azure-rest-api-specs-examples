package armhybridkubernetes_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridkubernetes/armhybridkubernetes/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ecee919199a39cc0d864410f540aa105bf7cdb64/specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/preview/2024-12-01-preview/examples/CreateCluster_KindExample.json
func ExampleConnectedClusterClient_BeginCreateOrReplace_createClusterKindExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridkubernetes.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectedClusterClient().BeginCreateOrReplace(ctx, "k8sc-rg", "testCluster", armhybridkubernetes.ConnectedCluster{
		Location: to.Ptr("East US"),
		Tags:     map[string]*string{},
		Identity: &armhybridkubernetes.ConnectedClusterIdentity{
			Type: to.Ptr(armhybridkubernetes.ResourceIdentityTypeSystemAssigned),
		},
		Kind: to.Ptr(armhybridkubernetes.ConnectedClusterKindProvisionedCluster),
		Properties: &armhybridkubernetes.ConnectedClusterProperties{
			AADProfile: &armhybridkubernetes.AADProfile{
				AdminGroupObjectIDs: []*string{
					to.Ptr("56f988bf-86f1-41af-91ab-2d7cd011db47")},
				EnableAzureRBAC: to.Ptr(true),
				TenantID:        to.Ptr("82f988bf-86f1-41af-91ab-2d7cd011db47"),
			},
			AgentPublicKeyCertificate: to.Ptr(""),
			ArcAgentProfile: &armhybridkubernetes.ArcAgentProfile{
				AgentAutoUpgrade:    to.Ptr(armhybridkubernetes.AutoUpgradeOptionsEnabled),
				DesiredAgentVersion: to.Ptr("0.1.0"),
				SystemComponents: []*armhybridkubernetes.SystemComponent{
					{
						Type:                 to.Ptr("Strato"),
						MajorVersion:         to.Ptr[int32](0),
						UserSpecifiedVersion: to.Ptr("0.1.1"),
					}},
			},
			AzureHybridBenefit:  to.Ptr(armhybridkubernetes.AzureHybridBenefitNotApplicable),
			Distribution:        to.Ptr("AKS"),
			DistributionVersion: to.Ptr("1.0"),
			OidcIssuerProfile: &armhybridkubernetes.OidcIssuerProfile{
				Enabled: to.Ptr(true),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectedCluster = armhybridkubernetes.ConnectedCluster{
	// 	Name: to.Ptr("connectedCluster1"),
	// 	Type: to.Ptr("Microsoft.Kubernetes/connectedClusters"),
	// 	ID: to.Ptr("/subscriptions/1bfbb5d0-917e-4346-9026-1d3b344417f5/resourceGroups/akkeshar/providers/Microsoft.Kubernetes/connectedClusters/connectedCluster1"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Identity: &armhybridkubernetes.ConnectedClusterIdentity{
	// 		Type: to.Ptr(armhybridkubernetes.ResourceIdentityTypeSystemAssigned),
	// 	},
	// 	Kind: to.Ptr(armhybridkubernetes.ConnectedClusterKindProvisionedCluster),
	// 	Properties: &armhybridkubernetes.ConnectedClusterProperties{
	// 		AADProfile: &armhybridkubernetes.AADProfile{
	// 			AdminGroupObjectIDs: []*string{
	// 				to.Ptr("56f988bf-86f1-41af-91ab-2d7cd011db47")},
	// 				EnableAzureRBAC: to.Ptr(true),
	// 				TenantID: to.Ptr("82f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 			},
	// 			AgentPublicKeyCertificate: to.Ptr(""),
	// 			AgentVersion: to.Ptr("0.1.0"),
	// 			ArcAgentProfile: &armhybridkubernetes.ArcAgentProfile{
	// 				AgentAutoUpgrade: to.Ptr(armhybridkubernetes.AutoUpgradeOptionsEnabled),
	// 				AgentState: to.Ptr("Succeeded"),
	// 				DesiredAgentVersion: to.Ptr("0.1.0"),
	// 				SystemComponents: []*armhybridkubernetes.SystemComponent{
	// 					{
	// 						Type: to.Ptr("Strato"),
	// 						CurrentVersion: to.Ptr("0.1.0"),
	// 						MajorVersion: to.Ptr[int32](0),
	// 						UserSpecifiedVersion: to.Ptr("0.1.1"),
	// 				}},
	// 			},
	// 			AzureHybridBenefit: to.Ptr(armhybridkubernetes.AzureHybridBenefitNotApplicable),
	// 			Distribution: to.Ptr("AKS"),
	// 			DistributionVersion: to.Ptr("1.0"),
	// 			KubernetesVersion: to.Ptr("1.17.0"),
	// 			OidcIssuerProfile: &armhybridkubernetes.OidcIssuerProfile{
	// 				Enabled: to.Ptr(true),
	// 				IssuerURL: to.Ptr("https://oidcdiscovery-northamerica-endpoint-gbcge4adgqebgxev.z01.azurefd.net/885hc665-0g4a-4a4b-732b-e4950new3bed/"),
	// 			},
	// 			ProvisioningState: to.Ptr(armhybridkubernetes.ProvisioningStateSucceeded),
	// 			TotalNodeCount: to.Ptr[int32](2),
	// 		},
	// 		SystemData: &armhybridkubernetes.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-17T07:06:33.917Z"); return t}()),
	// 			CreatedBy: to.Ptr("sikasire@microsoft.com"),
	// 			CreatedByType: to.Ptr(armhybridkubernetes.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-17T07:06:33.917Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("sikasire@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armhybridkubernetes.LastModifiedByTypeUser),
	// 		},
	// 	}
}
