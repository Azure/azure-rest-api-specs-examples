package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/685aad3f33d355c1d9c89d493ee9398865367bd8/specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/Apps_Get_VNetInjection.json
func ExampleAppsClient_Get_appsGetVNetInjection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAppsClient().Get(ctx, "myResourceGroup", "myservice", "myapp", &armappplatform.AppsClientGetOptions{SyncStatus: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AppResource = armappplatform.AppResource{
	// 	Name: to.Ptr("myapp"),
	// 	Type: to.Ptr("Microsoft.AppPlatform/Spring/apps"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/apps/myapp"),
	// 	SystemData: &armappplatform.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
	// 		CreatedBy: to.Ptr("sample-user"),
	// 		CreatedByType: to.Ptr(armappplatform.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sample-user"),
	// 		LastModifiedByType: to.Ptr(armappplatform.LastModifiedByTypeUser),
	// 	},
	// 	Identity: &armappplatform.ManagedIdentityProperties{
	// 		Type: to.Ptr(armappplatform.ManagedIdentityTypeSystemAssignedUserAssigned),
	// 		PrincipalID: to.Ptr("principalid"),
	// 		TenantID: to.Ptr("tenantid"),
	// 		UserAssignedIdentities: map[string]*armappplatform.UserAssignedManagedIdentity{
	// 			"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/samplegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": &armappplatform.UserAssignedManagedIdentity{
	// 				ClientID: to.Ptr("clientId1"),
	// 				PrincipalID: to.Ptr("principalId1"),
	// 			},
	// 			"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/samplegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2": &armappplatform.UserAssignedManagedIdentity{
	// 				ClientID: to.Ptr("clientId2"),
	// 				PrincipalID: to.Ptr("principalId2"),
	// 			},
	// 		},
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armappplatform.AppResourceProperties{
	// 		CustomPersistentDisks: []*armappplatform.CustomPersistentDiskResource{
	// 			{
	// 				CustomPersistentDiskProperties: &armappplatform.AzureFileVolume{
	// 					Type: to.Ptr(armappplatform.TypeAzureFileVolume),
	// 					MountOptions: []*string{
	// 						to.Ptr("uid=0"),
	// 						to.Ptr("gid=0"),
	// 						to.Ptr("dir_mode=0777"),
	// 						to.Ptr("file_mode=0777")},
	// 						MountPath: to.Ptr("/mypath1/mypath2"),
	// 						ShareName: to.Ptr("myFileShare"),
	// 					},
	// 					StorageID: to.Ptr("myASCStorageID"),
	// 			}},
	// 			EnableEndToEndTLS: to.Ptr(false),
	// 			Fqdn: to.Ptr("myapp.private.mydomain.com"),
	// 			HTTPSOnly: to.Ptr(false),
	// 			LoadedCertificates: []*armappplatform.LoadedCertificate{
	// 				{
	// 					LoadTrustStore: to.Ptr(false),
	// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/certificates/mycert1"),
	// 				},
	// 				{
	// 					LoadTrustStore: to.Ptr(true),
	// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/certificates/mycert2"),
	// 			}},
	// 			PersistentDisk: &armappplatform.PersistentDisk{
	// 				MountPath: to.Ptr("/mypersistentdisk"),
	// 				SizeInGB: to.Ptr[int32](2),
	// 				UsedInGB: to.Ptr[int32](1),
	// 			},
	// 			ProvisioningState: to.Ptr(armappplatform.AppResourceProvisioningStateSucceeded),
	// 			Public: to.Ptr(true),
	// 			TemporaryDisk: &armappplatform.TemporaryDisk{
	// 				MountPath: to.Ptr("/mytemporarydisk"),
	// 				SizeInGB: to.Ptr[int32](2),
	// 			},
	// 			URL: to.Ptr("myapp.myservice.private.azuremicroservices.io"),
	// 			VnetAddons: &armappplatform.AppVNetAddons{
	// 				PublicEndpoint: to.Ptr(true),
	// 				PublicEndpointURL: to.Ptr("myapp.myservice.azuremicroservices.io"),
	// 			},
	// 		},
	// 	}
}
