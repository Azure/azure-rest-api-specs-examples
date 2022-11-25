package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-11-01-preview/examples/CustomizedAccelerators_CreateOrUpdate.json
func ExampleCustomizedAcceleratorsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappplatform.NewCustomizedAcceleratorsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "myResourceGroup", "myservice", "default", "acc-name", armappplatform.CustomizedAcceleratorResource{
		Properties: &armappplatform.CustomizedAcceleratorProperties{
			Description: to.Ptr("acc-desc"),
			AcceleratorTags: []*string{
				to.Ptr("tag-a"),
				to.Ptr("tag-b")},
			DisplayName: to.Ptr("acc-name"),
			GitRepository: &armappplatform.AcceleratorGitRepository{
				AuthSetting: &armappplatform.AcceleratorSSHSetting{
					AuthType:         to.Ptr("SSH"),
					HostKey:          to.Ptr("git-auth-hostkey"),
					HostKeyAlgorithm: to.Ptr("git-auth-algorithm"),
					PrivateKey:       to.Ptr("git-auth-privatekey"),
				},
				Branch:            to.Ptr("git-branch"),
				Commit:            to.Ptr("12345"),
				GitTag:            to.Ptr("git-tag"),
				IntervalInSeconds: to.Ptr[int32](70),
				URL:               to.Ptr("git-url"),
			},
			IconURL: to.Ptr("acc-icon"),
		},
		SKU: &armappplatform.SKU{
			Name:     to.Ptr("E0"),
			Capacity: to.Ptr[int32](2),
			Tier:     to.Ptr("Enterprise"),
		},
	}, nil)
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
