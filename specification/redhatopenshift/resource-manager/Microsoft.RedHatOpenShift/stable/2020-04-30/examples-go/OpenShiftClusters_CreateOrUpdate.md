Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fredhatopenshift%2Farmredhatopenshift%2Fv0.2.0/sdk/resourcemanager/redhatopenshift/armredhatopenshift/README.md) on how to add the SDK to your project and authenticate.

```go
package armredhatopenshift_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift"
)

// x-ms-original-file: specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/stable/2020-04-30/examples/OpenShiftClusters_CreateOrUpdate.json
func ExampleOpenShiftClustersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armredhatopenshift.NewOpenShiftClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armredhatopenshift.OpenShiftCluster{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"key": to.StringPtr("value"),
			},
			Properties: &armredhatopenshift.OpenShiftClusterProperties{
				ApiserverProfile: &armredhatopenshift.APIServerProfile{
					Visibility: armredhatopenshift.Visibility("Public").ToPtr(),
				},
				ClusterProfile: &armredhatopenshift.ClusterProfile{
					Domain:          to.StringPtr("<domain>"),
					PullSecret:      to.StringPtr("<pull-secret>"),
					ResourceGroupID: to.StringPtr("<resource-group-id>"),
				},
				ConsoleProfile: &armredhatopenshift.ConsoleProfile{},
				IngressProfiles: []*armredhatopenshift.IngressProfile{
					{
						Name:       to.StringPtr("<name>"),
						Visibility: armredhatopenshift.Visibility("Public").ToPtr(),
					}},
				MasterProfile: &armredhatopenshift.MasterProfile{
					SubnetID: to.StringPtr("<subnet-id>"),
					VMSize:   armredhatopenshift.VMSize("Standard_D8s_v3").ToPtr(),
				},
				NetworkProfile: &armredhatopenshift.NetworkProfile{
					PodCidr:     to.StringPtr("<pod-cidr>"),
					ServiceCidr: to.StringPtr("<service-cidr>"),
				},
				ServicePrincipalProfile: &armredhatopenshift.ServicePrincipalProfile{
					ClientID:     to.StringPtr("<client-id>"),
					ClientSecret: to.StringPtr("<client-secret>"),
				},
				WorkerProfiles: []*armredhatopenshift.WorkerProfile{
					{
						Name:       to.StringPtr("<name>"),
						Count:      to.Int32Ptr(3),
						DiskSizeGB: to.Int32Ptr(128),
						SubnetID:   to.StringPtr("<subnet-id>"),
						VMSize:     armredhatopenshift.VMSize("Standard_D2s_v3").ToPtr(),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.OpenShiftClustersClientCreateOrUpdateResult)
}
```
