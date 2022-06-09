```go
package armredhatopenshift_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/stable/2022-04-01/examples/OpenShiftClusters_CreateOrUpdate.json
func ExampleOpenShiftClustersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armredhatopenshift.NewOpenShiftClustersClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"resourceGroup",
		"resourceName",
		armredhatopenshift.OpenShiftCluster{
			Location: to.Ptr("location"),
			Tags: map[string]*string{
				"key": to.Ptr("value"),
			},
			Properties: &armredhatopenshift.OpenShiftClusterProperties{
				ApiserverProfile: &armredhatopenshift.APIServerProfile{
					Visibility: to.Ptr(armredhatopenshift.VisibilityPublic),
				},
				ClusterProfile: &armredhatopenshift.ClusterProfile{
					Domain:               to.Ptr("cluster.location.aroapp.io"),
					FipsValidatedModules: to.Ptr(armredhatopenshift.FipsValidatedModulesEnabled),
					PullSecret:           to.Ptr("{\"auths\":{\"registry.connect.redhat.com\":{\"auth\":\"\"},\"registry.redhat.io\":{\"auth\":\"\"}}}"),
					ResourceGroupID:      to.Ptr("/subscriptions/subscriptionId/resourceGroups/clusterResourceGroup"),
				},
				ConsoleProfile: &armredhatopenshift.ConsoleProfile{},
				IngressProfiles: []*armredhatopenshift.IngressProfile{
					{
						Name:       to.Ptr("default"),
						Visibility: to.Ptr(armredhatopenshift.VisibilityPublic),
					}},
				MasterProfile: &armredhatopenshift.MasterProfile{
					EncryptionAtHost: to.Ptr(armredhatopenshift.EncryptionAtHostEnabled),
					SubnetID:         to.Ptr("/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/master"),
					VMSize:           to.Ptr("Standard_D8s_v3"),
				},
				NetworkProfile: &armredhatopenshift.NetworkProfile{
					PodCidr:     to.Ptr("10.128.0.0/14"),
					ServiceCidr: to.Ptr("172.30.0.0/16"),
				},
				ServicePrincipalProfile: &armredhatopenshift.ServicePrincipalProfile{
					ClientID:     to.Ptr("clientId"),
					ClientSecret: to.Ptr("clientSecret"),
				},
				WorkerProfiles: []*armredhatopenshift.WorkerProfile{
					{
						Name:       to.Ptr("worker"),
						Count:      to.Ptr[int32](3),
						DiskSizeGB: to.Ptr[int32](128),
						SubnetID:   to.Ptr("/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/worker"),
						VMSize:     to.Ptr("Standard_D2s_v3"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fredhatopenshift%2Farmredhatopenshift%2Fv1.0.0/sdk/resourcemanager/redhatopenshift/armredhatopenshift/README.md) on how to add the SDK to your project and authenticate.
