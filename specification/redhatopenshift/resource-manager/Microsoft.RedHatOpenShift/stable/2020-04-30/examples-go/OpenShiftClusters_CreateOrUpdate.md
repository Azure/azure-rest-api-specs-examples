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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/stable/2020-04-30/examples/OpenShiftClusters_CreateOrUpdate.json
func ExampleOpenShiftClustersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armredhatopenshift.NewOpenShiftClustersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armredhatopenshift.OpenShiftCluster{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"key": to.Ptr("value"),
			},
			Properties: &armredhatopenshift.OpenShiftClusterProperties{
				ApiserverProfile: &armredhatopenshift.APIServerProfile{
					Visibility: to.Ptr(armredhatopenshift.VisibilityPublic),
				},
				ClusterProfile: &armredhatopenshift.ClusterProfile{
					Domain:          to.Ptr("<domain>"),
					PullSecret:      to.Ptr("<pull-secret>"),
					ResourceGroupID: to.Ptr("<resource-group-id>"),
				},
				ConsoleProfile: &armredhatopenshift.ConsoleProfile{},
				IngressProfiles: []*armredhatopenshift.IngressProfile{
					{
						Name:       to.Ptr("<name>"),
						Visibility: to.Ptr(armredhatopenshift.VisibilityPublic),
					}},
				MasterProfile: &armredhatopenshift.MasterProfile{
					SubnetID: to.Ptr("<subnet-id>"),
					VMSize:   to.Ptr(armredhatopenshift.VMSizeStandardD8SV3),
				},
				NetworkProfile: &armredhatopenshift.NetworkProfile{
					PodCidr:     to.Ptr("<pod-cidr>"),
					ServiceCidr: to.Ptr("<service-cidr>"),
				},
				ServicePrincipalProfile: &armredhatopenshift.ServicePrincipalProfile{
					ClientID:     to.Ptr("<client-id>"),
					ClientSecret: to.Ptr("<client-secret>"),
				},
				WorkerProfiles: []*armredhatopenshift.WorkerProfile{
					{
						Name:       to.Ptr("<name>"),
						Count:      to.Ptr[int32](3),
						DiskSizeGB: to.Ptr[int32](128),
						SubnetID:   to.Ptr("<subnet-id>"),
						VMSize:     to.Ptr(armredhatopenshift.VMSizeStandardD2SV3),
					}},
			},
		},
		&armredhatopenshift.OpenShiftClustersClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fredhatopenshift%2Farmredhatopenshift%2Fv0.4.0/sdk/resourcemanager/redhatopenshift/armredhatopenshift/README.md) on how to add the SDK to your project and authenticate.
