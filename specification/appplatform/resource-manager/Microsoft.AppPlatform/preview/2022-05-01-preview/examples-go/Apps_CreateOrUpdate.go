package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-05-01-preview/examples/Apps_CreateOrUpdate.json
func ExampleAppsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappplatform.NewAppsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myservice",
		"myapp",
		armappplatform.AppResource{
			Identity: &armappplatform.ManagedIdentityProperties{
				Type: to.Ptr(armappplatform.ManagedIdentityTypeSystemAssignedUserAssigned),
				UserAssignedIdentities: map[string]*armappplatform.UserAssignedManagedIdentity{
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/samplegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {},
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/samplegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2": {},
				},
			},
			Location: to.Ptr("eastus"),
			Properties: &armappplatform.AppResourceProperties{
				AddonConfigs: map[string]map[string]interface{}{
					"ApplicationConfigurationService": {
						"resourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/configurationServices/myacs",
					},
					"ServiceRegistry": {
						"resourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/serviceRegistries/myServiceRegistry",
					},
				},
				CustomPersistentDisks: []*armappplatform.CustomPersistentDiskResource{
					{
						CustomPersistentDiskProperties: &armappplatform.AzureFileVolume{
							Type: to.Ptr(armappplatform.TypeAzureFileVolume),
							MountOptions: []*string{
								to.Ptr("uid=0"),
								to.Ptr("gid=0"),
								to.Ptr("dir_mode=0777"),
								to.Ptr("file_mode=0777")},
							MountPath: to.Ptr("/mypath1/mypath2"),
							ShareName: to.Ptr("myFileShare"),
						},
						StorageID: to.Ptr("myASCStorageID"),
					}},
				EnableEndToEndTLS: to.Ptr(false),
				Fqdn:              to.Ptr("myapp.mydomain.com"),
				HTTPSOnly:         to.Ptr(false),
				LoadedCertificates: []*armappplatform.LoadedCertificate{
					{
						LoadTrustStore: to.Ptr(false),
						ResourceID:     to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/certificates/mycert1"),
					},
					{
						LoadTrustStore: to.Ptr(true),
						ResourceID:     to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/certificates/mycert2"),
					}},
				PersistentDisk: &armappplatform.PersistentDisk{
					MountPath: to.Ptr("/mypersistentdisk"),
					SizeInGB:  to.Ptr[int32](2),
				},
				Public: to.Ptr(true),
				TemporaryDisk: &armappplatform.TemporaryDisk{
					MountPath: to.Ptr("/mytemporarydisk"),
					SizeInGB:  to.Ptr[int32](2),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
