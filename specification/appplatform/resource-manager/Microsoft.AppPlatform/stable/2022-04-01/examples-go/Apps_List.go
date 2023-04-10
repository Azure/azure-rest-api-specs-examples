package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Apps_List.json
func ExampleAppsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAppsClient().NewListPager("myResourceGroup", "myservice", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AppResourceCollection = armappplatform.AppResourceCollection{
		// 	Value: []*armappplatform.AppResource{
		// 		{
		// 			Name: to.Ptr("myapp"),
		// 			Type: to.Ptr("Microsoft.AppPlatform/Spring/apps"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/apps/myapp"),
		// 			SystemData: &armappplatform.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
		// 				CreatedBy: to.Ptr("sample-user"),
		// 				CreatedByType: to.Ptr(armappplatform.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("sample-user"),
		// 				LastModifiedByType: to.Ptr(armappplatform.LastModifiedByTypeUser),
		// 			},
		// 			Identity: &armappplatform.ManagedIdentityProperties{
		// 				Type: to.Ptr(armappplatform.ManagedIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("principalid"),
		// 				TenantID: to.Ptr("tenantid"),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armappplatform.AppResourceProperties{
		// 				EnableEndToEndTLS: to.Ptr(false),
		// 				Fqdn: to.Ptr("myapp.mydomain.com"),
		// 				HTTPSOnly: to.Ptr(false),
		// 				LoadedCertificates: []*armappplatform.LoadedCertificate{
		// 					{
		// 						LoadTrustStore: to.Ptr(false),
		// 						ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/certificates/mycert1"),
		// 					},
		// 					{
		// 						LoadTrustStore: to.Ptr(true),
		// 						ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/certificates/mycert2"),
		// 				}},
		// 				PersistentDisk: &armappplatform.PersistentDisk{
		// 					MountPath: to.Ptr("/mypersistentdisk"),
		// 					SizeInGB: to.Ptr[int32](2),
		// 					UsedInGB: to.Ptr[int32](1),
		// 				},
		// 				ProvisioningState: to.Ptr(armappplatform.AppResourceProvisioningStateSucceeded),
		// 				Public: to.Ptr(true),
		// 				TemporaryDisk: &armappplatform.TemporaryDisk{
		// 					MountPath: to.Ptr("/mytemporarydisk"),
		// 					SizeInGB: to.Ptr[int32](2),
		// 				},
		// 				URL: to.Ptr("myapp.myservice.azuremicroservices.io"),
		// 			},
		// 	}},
		// }
	}
}
