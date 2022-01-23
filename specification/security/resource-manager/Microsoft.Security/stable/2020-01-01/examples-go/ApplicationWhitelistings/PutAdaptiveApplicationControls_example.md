Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.3.0/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/ApplicationWhitelistings/PutAdaptiveApplicationControls_example.json
func ExampleAdaptiveApplicationControlsClient_Put() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewAdaptiveApplicationControlsClient("<subscription-id>",
		"<asc-location>", cred, nil)
	res, err := client.Put(ctx,
		"<group-name>",
		armsecurity.AdaptiveApplicationControlGroup{
			Properties: &armsecurity.AdaptiveApplicationControlGroupData{
				EnforcementMode: armsecurity.EnforcementMode("Audit").ToPtr(),
				PathRecommendations: []*armsecurity.PathRecommendation{
					{
						Type:                armsecurity.RecommendationType("PublisherSignature").ToPtr(),
						Path:                to.StringPtr("<path>"),
						Action:              armsecurity.RecommendationAction("Recommended").ToPtr(),
						Common:              to.BoolPtr(true),
						ConfigurationStatus: armsecurity.ConfigurationStatus("Configured").ToPtr(),
						FileType:            armsecurity.FileType("Exe").ToPtr(),
						PublisherInfo: &armsecurity.PublisherInfo{
							BinaryName:    to.StringPtr("<binary-name>"),
							ProductName:   to.StringPtr("<product-name>"),
							PublisherName: to.StringPtr("<publisher-name>"),
							Version:       to.StringPtr("<version>"),
						},
						UserSids: []*string{
							to.StringPtr("S-1-1-0")},
						Usernames: []*armsecurity.UserRecommendation{
							{
								RecommendationAction: armsecurity.RecommendationAction("Recommended").ToPtr(),
								Username:             to.StringPtr("<username>"),
							}},
					},
					{
						Type:                armsecurity.RecommendationType("ProductSignature").ToPtr(),
						Path:                to.StringPtr("<path>"),
						Action:              armsecurity.RecommendationAction("Recommended").ToPtr(),
						Common:              to.BoolPtr(true),
						ConfigurationStatus: armsecurity.ConfigurationStatus("Configured").ToPtr(),
						FileType:            armsecurity.FileType("Exe").ToPtr(),
						PublisherInfo: &armsecurity.PublisherInfo{
							BinaryName:    to.StringPtr("<binary-name>"),
							ProductName:   to.StringPtr("<product-name>"),
							PublisherName: to.StringPtr("<publisher-name>"),
							Version:       to.StringPtr("<version>"),
						},
						UserSids: []*string{
							to.StringPtr("S-1-1-0")},
						Usernames: []*armsecurity.UserRecommendation{
							{
								RecommendationAction: armsecurity.RecommendationAction("Recommended").ToPtr(),
								Username:             to.StringPtr("<username>"),
							}},
					},
					{
						Type:                armsecurity.RecommendationType("PublisherSignature").ToPtr(),
						Path:                to.StringPtr("<path>"),
						Action:              armsecurity.RecommendationAction("Recommended").ToPtr(),
						Common:              to.BoolPtr(true),
						ConfigurationStatus: armsecurity.ConfigurationStatus("Configured").ToPtr(),
						FileType:            armsecurity.FileType("Exe").ToPtr(),
						PublisherInfo: &armsecurity.PublisherInfo{
							BinaryName:    to.StringPtr("<binary-name>"),
							ProductName:   to.StringPtr("<product-name>"),
							PublisherName: to.StringPtr("<publisher-name>"),
							Version:       to.StringPtr("<version>"),
						},
						UserSids: []*string{
							to.StringPtr("S-1-1-0")},
						Usernames: []*armsecurity.UserRecommendation{
							{
								RecommendationAction: armsecurity.RecommendationAction("Recommended").ToPtr(),
								Username:             to.StringPtr("<username>"),
							}},
					},
					{
						Type:   armsecurity.RecommendationType("File").ToPtr(),
						Path:   to.StringPtr("<path>"),
						Action: armsecurity.RecommendationAction("Add").ToPtr(),
						Common: to.BoolPtr(true),
					}},
				ProtectionMode: &armsecurity.ProtectionMode{
					Exe:    armsecurity.EnforcementMode("Audit").ToPtr(),
					Msi:    armsecurity.EnforcementMode("None").ToPtr(),
					Script: armsecurity.EnforcementMode("None").ToPtr(),
				},
				VMRecommendations: []*armsecurity.VMRecommendation{
					{
						ConfigurationStatus:  armsecurity.ConfigurationStatus("Configured").ToPtr(),
						EnforcementSupport:   armsecurity.EnforcementSupport("Supported").ToPtr(),
						RecommendationAction: armsecurity.RecommendationAction("Recommended").ToPtr(),
						ResourceID:           to.StringPtr("<resource-id>"),
					},
					{
						ConfigurationStatus:  armsecurity.ConfigurationStatus("Configured").ToPtr(),
						EnforcementSupport:   armsecurity.EnforcementSupport("Supported").ToPtr(),
						RecommendationAction: armsecurity.RecommendationAction("Recommended").ToPtr(),
						ResourceID:           to.StringPtr("<resource-id>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AdaptiveApplicationControlsClientPutResult)
}
```
