package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2025-09-01/ScriptPackages_Get.json
func ExampleScriptPackagesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScriptPackagesClient().Get(ctx, "group1", "cloud1", "Microsoft.AVS.Management@3.0.48", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armavs.ScriptPackagesClientGetResponse{
	// 	ScriptPackage: &armavs.ScriptPackage{
	// 		Name: to.Ptr("Microsoft.AVS.Management@3.0.48"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/scriptPackages/Microsoft.AVS.Management@3.0.48"),
	// 		Properties: &armavs.ScriptPackageProperties{
	// 			Description: to.Ptr("Various cmdlets for elevated access to Private Cloud administrative functions"),
	// 			Version: to.Ptr("3.0.48"),
	// 			Company: to.Ptr("Microsoft"),
	// 			URI: to.Ptr("https://microsoft.com"),
	// 		},
	// 		Type: to.Ptr("Microsoft.AVS/privateClouds/scriptPackages"),
	// 	},
	// }
}
