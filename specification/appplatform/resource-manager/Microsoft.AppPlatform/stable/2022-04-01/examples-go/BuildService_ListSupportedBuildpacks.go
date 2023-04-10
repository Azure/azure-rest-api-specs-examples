package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/BuildService_ListSupportedBuildpacks.json
func ExampleBuildServiceClient_ListSupportedBuildpacks() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBuildServiceClient().ListSupportedBuildpacks(ctx, "myResourceGroup", "myservice", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SupportedBuildpacksCollection = armappplatform.SupportedBuildpacksCollection{
	// 	Value: []*armappplatform.SupportedBuildpackResource{
	// 		{
	// 			Name: to.Ptr("tanzu-buildpacks-java-azure"),
	// 			Type: to.Ptr("Microsoft.AppPlatform/Spring/buildServices/supportedBuildpacks"),
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/buildServices/default/supportedBuildpacks/tanzu-buildpacks-java-azure"),
	// 			SystemData: &armappplatform.SystemData{
	// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
	// 				CreatedBy: to.Ptr("sample-user"),
	// 				CreatedByType: to.Ptr(armappplatform.CreatedByTypeUser),
	// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
	// 				LastModifiedBy: to.Ptr("sample-user"),
	// 				LastModifiedByType: to.Ptr(armappplatform.LastModifiedByTypeUser),
	// 			},
	// 			Properties: &armappplatform.SupportedBuildpackResourceProperties{
	// 				BuildpackID: to.Ptr("tanzu-buildpacks/java-azure"),
	// 			},
	// 	}},
	// }
}
