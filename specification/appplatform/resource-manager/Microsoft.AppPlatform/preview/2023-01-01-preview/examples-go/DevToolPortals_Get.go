package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2023-01-01-preview/examples/DevToolPortals_Get.json
func ExampleDevToolPortalsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDevToolPortalsClient().Get(ctx, "myResourceGroup", "myservice", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DevToolPortalResource = armappplatform.DevToolPortalResource{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.AppPlatform/Spring/applicationLiveViews"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/applicationLiveViews/default"),
	// 	SystemData: &armappplatform.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
	// 		CreatedBy: to.Ptr("sample-user"),
	// 		CreatedByType: to.Ptr(armappplatform.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sample-user"),
	// 		LastModifiedByType: to.Ptr(armappplatform.LastModifiedByTypeUser),
	// 	},
	// 	Properties: &armappplatform.DevToolPortalProperties{
	// 		Features: &armappplatform.DevToolPortalFeatureSettings{
	// 			ApplicationAccelerator: &armappplatform.DevToolPortalFeatureDetail{
	// 				Route: to.Ptr("create"),
	// 				State: to.Ptr(armappplatform.DevToolPortalFeatureStateEnabled),
	// 			},
	// 			ApplicationLiveView: &armappplatform.DevToolPortalFeatureDetail{
	// 				Route: to.Ptr("appliveview"),
	// 				State: to.Ptr(armappplatform.DevToolPortalFeatureStateEnabled),
	// 			},
	// 		},
	// 		Instances: []*armappplatform.DevToolPortalInstance{
	// 			{
	// 				Name: to.Ptr("app-live-view-server-name"),
	// 				Status: to.Ptr("Running"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armappplatform.DevToolPortalProvisioningStateSucceeded),
	// 		Public: to.Ptr(true),
	// 		ResourceRequests: &armappplatform.DevToolPortalResourceRequests{
	// 			CPU: to.Ptr("1"),
	// 			InstanceCount: to.Ptr[int32](1),
	// 			Memory: to.Ptr("1Gi"),
	// 		},
	// 		SsoProperties: &armappplatform.DevToolPortalSsoProperties{
	// 			ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			MetadataURL: to.Ptr("https://login.microsoft.com/00000000-0000-0000-0000-000000000000/v2.0/.well-known/openid-configuration"),
	// 			Scopes: []*string{
	// 				to.Ptr("openid")},
	// 			},
	// 			URL: to.Ptr("aaa.com"),
	// 		},
	// 	}
}
