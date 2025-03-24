package armhybridkubernetes_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridkubernetes/armhybridkubernetes/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ecee919199a39cc0d864410f540aa105bf7cdb64/specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/preview/2024-12-01-preview/examples/CreateClusterAgentless_KindAWSExample.json
func ExampleConnectedClusterClient_BeginCreateOrReplace_createClusterAgentlessKindAwsExample() {
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
			Type: to.Ptr(armhybridkubernetes.ResourceIdentityTypeNone),
		},
		Kind: to.Ptr(armhybridkubernetes.ConnectedClusterKindAWS),
		Properties: &armhybridkubernetes.ConnectedClusterProperties{
			AgentPublicKeyCertificate: to.Ptr(""),
			Distribution:              to.Ptr("eks"),
			Infrastructure:            to.Ptr("aws"),
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
	// 		Type: to.Ptr(armhybridkubernetes.ResourceIdentityTypeNone),
	// 	},
	// 	Kind: to.Ptr(armhybridkubernetes.ConnectedClusterKindAWS),
	// 	Properties: &armhybridkubernetes.ConnectedClusterProperties{
	// 		AADProfile: &armhybridkubernetes.AADProfile{
	// 		},
	// 		AgentPublicKeyCertificate: to.Ptr(""),
	// 		ArcAgentProfile: &armhybridkubernetes.ArcAgentProfile{
	// 		},
	// 		AzureHybridBenefit: to.Ptr(armhybridkubernetes.AzureHybridBenefitNotApplicable),
	// 		ConnectivityStatus: to.Ptr(armhybridkubernetes.ConnectivityStatusAgentNotInstalled),
	// 		Distribution: to.Ptr("eks"),
	// 		Infrastructure: to.Ptr("aws"),
	// 		PrivateLinkState: to.Ptr(armhybridkubernetes.PrivateLinkStateDisabled),
	// 		ProvisioningState: to.Ptr(armhybridkubernetes.ProvisioningStateSucceeded),
	// 	},
	// 	SystemData: &armhybridkubernetes.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-17T07:06:33.917Z"); return t}()),
	// 		CreatedBy: to.Ptr("sikasire@microsoft.com"),
	// 		CreatedByType: to.Ptr(armhybridkubernetes.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-17T07:06:33.917Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sikasire@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armhybridkubernetes.LastModifiedByTypeUser),
	// 	},
	// }
}
