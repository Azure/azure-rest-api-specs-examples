Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhdinsight%2Farmhdinsight%2Fv1.0.0/sdk/resourcemanager/hdinsight/armhdinsight/README.md) on how to add the SDK to your project and authenticate.

```go
package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/CreateApplication.json
func ExampleApplicationsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhdinsight.NewApplicationsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"rg1",
		"cluster1",
		"hue",
		armhdinsight.Application{
			Properties: &armhdinsight.ApplicationProperties{
				ApplicationType: to.Ptr("CustomApplication"),
				ComputeProfile: &armhdinsight.ComputeProfile{
					Roles: []*armhdinsight.Role{
						{
							Name: to.Ptr("edgenode"),
							HardwareProfile: &armhdinsight.HardwareProfile{
								VMSize: to.Ptr("Standard_D12_v2"),
							},
							TargetInstanceCount: to.Ptr[int32](1),
						}},
				},
				Errors: []*armhdinsight.Errors{},
				HTTPSEndpoints: []*armhdinsight.ApplicationGetHTTPSEndpoint{
					{
						AccessModes: []*string{
							to.Ptr("WebPage")},
						DestinationPort: to.Ptr[int32](20000),
						SubDomainSuffix: to.Ptr("dss"),
					}},
				InstallScriptActions: []*armhdinsight.RuntimeScriptAction{
					{
						Name:       to.Ptr("app-install-app1"),
						Parameters: to.Ptr("-version latest -port 20000"),
						Roles: []*string{
							to.Ptr("edgenode")},
						URI: to.Ptr("https://.../install.sh"),
					}},
				UninstallScriptActions: []*armhdinsight.RuntimeScriptAction{},
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
