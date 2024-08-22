package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c7b98b36e4023331545051284d8500adf98f02fe/specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/CloudServiceRoleInstance_Get_InstanceView.json
func ExampleCloudServiceRoleInstancesClient_GetInstanceView() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCloudServiceRoleInstancesClient().GetInstanceView(ctx, "{roleInstance-name}", "ConstosoRG", "{cs-name}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleInstanceView = armcompute.RoleInstanceView{
	// 	PlatformFaultDomain: to.Ptr[int32](0),
	// 	PlatformUpdateDomain: to.Ptr[int32](0),
	// 	PrivateID: to.Ptr("3491bc0c-1f6c-444f-b1d0-ec0751a74e3e"),
	// 	Statuses: []*armcompute.ResourceInstanceViewStatus{
	// 		{
	// 			Code: to.Ptr("RoleState/RoleStateStarted"),
	// 			DisplayStatus: to.Ptr("RoleStateStarted"),
	// 			Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 			Message: to.Ptr(""),
	// 	}},
	// }
}
