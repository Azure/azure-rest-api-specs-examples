package armdeviceregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry/v2"
)

// Generated from example definition: 2025-10-01/Get_AssetEndpointProfile_With_SyncStatus.json
func ExampleAssetEndpointProfilesClient_Get_getAssetEndpointProfileWithSyncStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssetEndpointProfilesClient().Get(ctx, "myResourceGroup", "my-assetendpointprofile", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdeviceregistry.AssetEndpointProfilesClientGetResponse{
	// 	AssetEndpointProfile: &armdeviceregistry.AssetEndpointProfile{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/my-assetendpointprofile"),
	// 		Name: to.Ptr("my-assetendpointprofile"),
	// 		Type: to.Ptr("Microsoft.DeviceRegistry/assetEndpointProfiles"),
	// 		Location: to.Ptr("West Europe"),
	// 		ExtendedLocation: &armdeviceregistry.ExtendedLocation{
	// 			Type: to.Ptr("CustomLocation"),
	// 			Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"site": to.Ptr("building-1"),
	// 		},
	// 		SystemData: &armdeviceregistry.SystemData{
	// 			CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
	// 			CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.2516899Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
	// 			LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.0922793Z"); return t}()),
	// 		},
	// 		Properties: &armdeviceregistry.AssetEndpointProfileProperties{
	// 			UUID: to.Ptr("1824a74f-21e1-4458-ae07-604d3a241d2e"),
	// 			TargetAddress: to.Ptr("https://www.example.com/myTargetAddress"),
	// 			EndpointProfileType: to.Ptr("myEndpointProfileType"),
	// 			Authentication: &armdeviceregistry.Authentication{
	// 				Method: to.Ptr(armdeviceregistry.AuthenticationMethodUsernamePassword),
	// 				UsernamePasswordCredentials: &armdeviceregistry.UsernamePasswordCredentials{
	// 					UsernameSecretName: to.Ptr("usernameRef"),
	// 					PasswordSecretName: to.Ptr("passwordRef"),
	// 				},
	// 			},
	// 			Status: &armdeviceregistry.AssetEndpointProfileStatus{
	// 				Errors: []*armdeviceregistry.AssetEndpointProfileStatusError{
	// 					{
	// 						Code: to.Ptr[int32](500),
	// 						Message: to.Ptr("Internal Server Error"),
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
