package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/ComponentGet.json
func ExampleComponentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewComponentsClient().Get(ctx, "rg", "testNf", "testComponent", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Component = armhybridnetwork.Component{
	// 	Name: to.Ptr("testComponent"),
	// 	Type: to.Ptr("Microsoft.HybridNetwork/networkFunctions/components"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/networkFunctions/testNf/components/testComponent"),
	// 	Properties: &armhybridnetwork.ComponentProperties{
	// 		DeploymentProfile: to.Ptr("{\"chart\":{\"name\":\"testChartName\",\"version\":\"v1.0.0\"},\"releaseName\":\"testReleaseName\",\"targetNamespace\":\"testTargetNameSpace\",\"values\":{\".repoBase\":\"testrepo.azurecr.io/\"}}"),
	// 		DeploymentStatus: &armhybridnetwork.DeploymentStatusProperties{
	// 			NextExpectedUpdateAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-10T02:06:01.116Z"); return t}()),
	// 			Resources: &armhybridnetwork.Resources{
	// 				Deployments: []*armhybridnetwork.Deployment{
	// 					{
	// 						Name: to.Ptr("nginix"),
	// 						Available: to.Ptr[int32](3),
	// 						CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-10T02:05:01.116Z"); return t}()),
	// 						Desired: to.Ptr[int32](3),
	// 						Namespace: to.Ptr("core"),
	// 						Ready: to.Ptr[int32](3),
	// 						UpToDate: to.Ptr[int32](3),
	// 				}},
	// 				Pods: []*armhybridnetwork.Pod{
	// 					{
	// 						Name: to.Ptr("nginix"),
	// 						CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-10T02:05:01.116Z"); return t}()),
	// 						Desired: to.Ptr[int32](3),
	// 						Events: []*armhybridnetwork.PodEvent{
	// 							{
	// 								Type: to.Ptr(armhybridnetwork.PodEventTypeNormal),
	// 								LastSeenTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-10T02:06:01.116Z"); return t}()),
	// 								Message: to.Ptr("Pulling image nginx"),
	// 								Reason: to.Ptr("Pulling"),
	// 						}},
	// 						Namespace: to.Ptr("core"),
	// 						Ready: to.Ptr[int32](3),
	// 						Status: to.Ptr(armhybridnetwork.PodStatusSucceeded),
	// 				}},
	// 			},
	// 			Status: to.Ptr(armhybridnetwork.StatusInstalling),
	// 		},
	// 		ProvisioningState: to.Ptr(armhybridnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}
