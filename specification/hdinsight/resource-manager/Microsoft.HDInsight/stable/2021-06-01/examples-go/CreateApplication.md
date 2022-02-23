Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhdinsight%2Farmhdinsight%2Fv0.2.1/sdk/resourcemanager/hdinsight/armhdinsight/README.md) on how to add the SDK to your project and authenticate.

```go
package armhdinsight_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/CreateApplication.json
func ExampleApplicationsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhdinsight.NewApplicationsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<application-name>",
		armhdinsight.Application{
			Properties: &armhdinsight.ApplicationProperties{
				ApplicationType: to.StringPtr("<application-type>"),
				ComputeProfile: &armhdinsight.ComputeProfile{
					Roles: []*armhdinsight.Role{
						{
							Name: to.StringPtr("<name>"),
							HardwareProfile: &armhdinsight.HardwareProfile{
								VMSize: to.StringPtr("<vmsize>"),
							},
							TargetInstanceCount: to.Int32Ptr(1),
						}},
				},
				Errors: []*armhdinsight.Errors{},
				HTTPSEndpoints: []*armhdinsight.ApplicationGetHTTPSEndpoint{
					{
						AccessModes: []*string{
							to.StringPtr("WebPage")},
						DestinationPort: to.Int32Ptr(20000),
						SubDomainSuffix: to.StringPtr("<sub-domain-suffix>"),
					}},
				InstallScriptActions: []*armhdinsight.RuntimeScriptAction{
					{
						Name:       to.StringPtr("<name>"),
						Parameters: to.StringPtr("<parameters>"),
						Roles: []*string{
							to.StringPtr("edgenode")},
						URI: to.StringPtr("<uri>"),
					}},
				UninstallScriptActions: []*armhdinsight.RuntimeScriptAction{},
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
	log.Printf("Response result: %#v\n", res.ApplicationsClientCreateResult)
}
```
