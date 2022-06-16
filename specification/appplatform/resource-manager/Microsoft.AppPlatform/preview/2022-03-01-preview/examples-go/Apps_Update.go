package armappplatform_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-03-01-preview/examples/Apps_Update.json
func ExampleAppsClient_BeginUpdate() {
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
	poller, err := client.BeginUpdate(ctx,
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
				CustomPersistentDisks: []*armappplatform.CustomPersistentDiskResource{
					{
						CustomPersistentDiskProperties: &armappplatform.AzureFileVolume{
							Type:         to.Ptr(armappplatform.TypeAzureFileVolume),
							MountOptions: []*string{},
							MountPath:    to.Ptr("<mount-path>"),
							ShareName:    to.Ptr("<share-name>"),
						},
						StorageID: to.Ptr("<storage-id>"),
					}},
				EnableEndToEndTLS: to.Ptr(false),
				Fqdn:              to.Ptr("<fqdn>"),
				HTTPSOnly:         to.Ptr(false),
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
		&armappplatform.AppsClientBeginUpdateOptions{ResumeToken: ""})
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
