package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/GetApplicationInProgress.json
func ExampleApplicationsClient_Get_getApplicationOnHdInsightClusterCreationInProgress() {
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
	// 	Etag: to.Ptr("2C128F8E-BB26-4637-99E4-18EBC39FD51F"),
	// 	Properties: &armhdinsight.ApplicationProperties{
	// 		ApplicationState: to.Ptr("AzureVMConfiguration"),
	// 		ApplicationType: to.Ptr("CustomApplication"),
	// 		ComputeProfile: &armhdinsight.ComputeProfile{
	// 			Roles: []*armhdinsight.Role{
	// 				{
	// 					Name: to.Ptr("edgenode"),
	// 					HardwareProfile: &armhdinsight.HardwareProfile{
	// 						VMSize: to.Ptr("Standard_D3"),
	// 					},
	// 					TargetInstanceCount: to.Ptr[int32](1),
	// 			}},
	// 		},
	// 		CreatedDate: to.Ptr("2017-03-28T02:01:25.107"),
	// 		HTTPSEndpoints: []*armhdinsight.ApplicationGetHTTPSEndpoint{
	// 			{
	// 				AccessModes: []*string{
	// 					to.Ptr("WebPage")},
	// 					DestinationPort: to.Ptr[int32](18630),
	// 					Location: to.Ptr("location"),
	// 					PublicPort: to.Ptr[int32](443),
	// 			}},
	// 			InstallScriptActions: []*armhdinsight.RuntimeScriptAction{
	// 				{
	// 					Name: to.Ptr("app-Install"),
	// 					Roles: []*string{
	// 						to.Ptr("edgenode")},
	// 						URI: to.Ptr("https://app.com/azure/app_install.sh"),
	// 				}},
	// 				MarketplaceIdentifier: to.Ptr("id"),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 				SSHEndpoints: []*armhdinsight.ApplicationGetEndpoint{
	// 				},
	// 				UninstallScriptActions: []*armhdinsight.RuntimeScriptAction{
	// 				},
	// 			},
	// 			Tags: map[string]*string{
	// 				"key1": to.Ptr("val1"),
	// 			},
	// 		}
}
