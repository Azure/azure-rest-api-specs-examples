package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2023-01-01-preview/examples/CustomizedAccelerators_Validate.json
func ExampleCustomizedAcceleratorsClient_Validate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomizedAcceleratorsClient().Validate(ctx, "myResourceGroup", "myservice", "default", "acc-name", armappplatform.CustomizedAcceleratorProperties{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomizedAcceleratorValidateResult = armappplatform.CustomizedAcceleratorValidateResult{
	// 	ErrorMessage: to.Ptr(""),
	// 	State: to.Ptr(armappplatform.CustomizedAcceleratorValidateResultStateValid),
	// }
}
