package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aec83a5f0ed56da4fd16fa027b9fa27edfa8988b/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/GetApplicationCreated.json
func ExampleApplicationsClient_Get_getApplicationOnHdInsightClusterSuccessfullyCreated() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationsClient().Get(ctx, "rg1", "cluster1", "app", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Application = armhdinsight.Application{
	// 	Name: to.Ptr("app"),
	// 	Type: to.Ptr("Microsoft.HDInsight/clusters/applications"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.HDInsight/clusters/cluster1/applications/app"),
	// 	Etag: to.Ptr("CF938302-6B4D-44A0-A6D2-C0D67E847AEC"),
	// 	Properties: &armhdinsight.ApplicationProperties{
	// 		ApplicationState: to.Ptr("Running"),
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
	// 		CreatedDate: to.Ptr("2017-03-22T21:34:39.293"),
	// 		HTTPSEndpoints: []*armhdinsight.ApplicationGetHTTPSEndpoint{
	// 			{
	// 				AccessModes: []*string{
	// 					to.Ptr("WebPage")},
	// 					DestinationPort: to.Ptr[int32](20000),
	// 					Location: to.Ptr("https://cluster1.apps.azurehdinsight.net:443"),
	// 					PublicPort: to.Ptr[int32](443),
	// 			}},
	// 			InstallScriptActions: []*armhdinsight.RuntimeScriptAction{
	// 				{
	// 					Name: to.Ptr("app-install"),
	// 					Roles: []*string{
	// 						to.Ptr("edgenode")},
	// 						URI: to.Ptr("http://app.com/public/hdi-app/20170301/app-install.sh"),
	// 				}},
	// 				MarketplaceIdentifier: to.Ptr("appMarketId"),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 				SSHEndpoints: []*armhdinsight.ApplicationGetEndpoint{
	// 					{
	// 						DestinationPort: to.Ptr[int32](22),
	// 						Location: to.Ptr("cluster1-ssh.azurehdinsight.net:22"),
	// 						PublicPort: to.Ptr[int32](22),
	// 				}},
	// 				UninstallScriptActions: []*armhdinsight.RuntimeScriptAction{
	// 				},
	// 			},
	// 			Tags: map[string]*string{
	// 				"key1": to.Ptr("val1"),
	// 			},
	// 		}
}
