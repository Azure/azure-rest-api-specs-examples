package armappplatform_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-03-01-preview/examples/Apps_CreateOrUpdate.json
func ExampleAppsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armappplatform.NewAppsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<app-name>",
		armappplatform.AppResource{
			Identity: &armappplatform.ManagedIdentityProperties{
				Type: to.Ptr(armappplatform.ManagedIdentityTypeSystemAssignedUserAssigned),
				UserAssignedIdentities: map[string]*armappplatform.UserAssignedManagedIdentity{
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/samplegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {},
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/samplegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2": {},
				},
			},
			Location: to.Ptr("<location>"),
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
							MountPath: to.Ptr("<mount-path>"),
							ShareName: to.Ptr("<share-name>"),
						},
						StorageID: to.Ptr("<storage-id>"),
					}},
				EnableEndToEndTLS: to.Ptr(false),
				Fqdn:              to.Ptr("<fqdn>"),
				HTTPSOnly:         to.Ptr(false),
				LoadedCertificates: []*armappplatform.LoadedCertificate{
					{
						LoadTrustStore: to.Ptr(false),
						ResourceID:     to.Ptr("<resource-id>"),
					},
					{
						LoadTrustStore: to.Ptr(true),
						ResourceID:     to.Ptr("<resource-id>"),
					}},
				PersistentDisk: &armappplatform.PersistentDisk{
					MountPath: to.Ptr("<mount-path>"),
					SizeInGB:  to.Ptr[int32](2),
				},
				Public: to.Ptr(true),
				TemporaryDisk: &armappplatform.TemporaryDisk{
					MountPath: to.Ptr("<mount-path>"),
					SizeInGB:  to.Ptr[int32](2),
				},
			},
		},
		&armappplatform.AppsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
