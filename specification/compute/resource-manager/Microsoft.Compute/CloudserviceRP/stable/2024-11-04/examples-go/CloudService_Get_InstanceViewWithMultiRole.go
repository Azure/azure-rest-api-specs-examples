package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2024-11-04/examples/CloudService_Get_InstanceViewWithMultiRole.json
func ExampleCloudServicesClient_GetInstanceView() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCloudServicesClient().GetInstanceView(ctx, "ConstosoRG", "{cs-name}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CloudServiceInstanceView = armcompute.CloudServiceInstanceView{
	// 	PrivateIDs: []*string{
	// 		to.Ptr("3491bc0c-1f6c-444f-b1d0-ec0751a74e3e")},
	// 		RoleInstance: &armcompute.InstanceViewStatusesSummary{
	// 			StatusesSummary: []*armcompute.StatusCodeCount{
	// 				{
	// 					Code: to.Ptr("ProvisioningState/succeeded"),
	// 					Count: to.Ptr[int32](4),
	// 				},
	// 				{
	// 					Code: to.Ptr("PowerState/started"),
	// 					Count: to.Ptr[int32](4),
	// 				},
	// 				{
	// 					Code: to.Ptr("RoleState/RoleStateStarted"),
	// 					Count: to.Ptr[int32](4),
	// 			}},
	// 		},
	// 		SdkVersion: to.Ptr("2.9.6496.3"),
	// 		Statuses: []*armcompute.ResourceInstanceViewStatus{
	// 			{
	// 				Code: to.Ptr("ProvisioningState/succeeded"),
	// 				DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 				Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-12T11:20:07.095Z"); return t}()),
	// 			},
	// 			{
	// 				Code: to.Ptr("PowerState/started"),
	// 				DisplayStatus: to.Ptr("Started"),
	// 				Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-12T11:20:07.095Z"); return t}()),
	// 			},
	// 			{
	// 				Code: to.Ptr("CurrentUpgradeDomain/-1"),
	// 				DisplayStatus: to.Ptr("Current Upgrade Domain of cloud service is -1."),
	// 				Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 			},
	// 			{
	// 				Code: to.Ptr("MaxUpgradeDomain/1"),
	// 				DisplayStatus: to.Ptr("Max Upgrade Domain of cloud service is 1."),
	// 				Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 		}},
	// 	}
}
