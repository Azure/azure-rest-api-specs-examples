package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aec83a5f0ed56da4fd16fa027b9fa27edfa8988b/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/CreateApplication.json
func ExampleApplicationsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationsClient().BeginCreate(ctx, "rg1", "cluster1", "hue", armhdinsight.Application{
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
	// res.Application = armhdinsight.Application{
	// 	Name: to.Ptr("hue"),
	// 	Type: to.Ptr("Microsoft.HDInsight/clusters/applications"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.HDInsight/clusters/cluster1/applications/hue"),
	// 	Etag: to.Ptr("etag"),
	// 	Properties: &armhdinsight.ApplicationProperties{
	// 		ApplicationState: to.Ptr("ApplicationConfiguration"),
	// 		ApplicationType: to.Ptr("CustomApplication"),
	// 		ComputeProfile: &armhdinsight.ComputeProfile{
	// 			Roles: []*armhdinsight.Role{
	// 				{
	// 					Name: to.Ptr("edgenode"),
	// 					HardwareProfile: &armhdinsight.HardwareProfile{
	// 						VMSize: to.Ptr("Standard_D12_v2"),
	// 					},
	// 					TargetInstanceCount: to.Ptr[int32](1),
	// 			}},
	// 		},
	// 		CreatedDate: to.Ptr("2017-02-28"),
	// 		Errors: []*armhdinsight.Errors{
	// 		},
	// 		HTTPSEndpoints: []*armhdinsight.ApplicationGetHTTPSEndpoint{
	// 			{
	// 				AccessModes: []*string{
	// 					to.Ptr("WebPage")},
	// 					DestinationPort: to.Ptr[int32](20000),
	// 			}},
	// 			InstallScriptActions: []*armhdinsight.RuntimeScriptAction{
	// 				{
	// 					Name: to.Ptr("app-install-app1"),
	// 					Roles: []*string{
	// 						to.Ptr("edgenode")},
	// 						URI: to.Ptr("https://.../install.sh"),
	// 				}},
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 				SSHEndpoints: []*armhdinsight.ApplicationGetEndpoint{
	// 				},
	// 				UninstallScriptActions: []*armhdinsight.RuntimeScriptAction{
	// 				},
	// 			},
	// 			Tags: map[string]*string{
	// 			},
	// 		}
}
