Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.6.0/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/ApplicationWhitelistings/PutAdaptiveApplicationControls_example.json
func ExampleAdaptiveApplicationControlsClient_Put() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsecurity.NewAdaptiveApplicationControlsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Put(ctx,
		"<asc-location>",
		"<group-name>",
		armsecurity.AdaptiveApplicationControlGroup{
			Properties: &armsecurity.AdaptiveApplicationControlGroupData{
				EnforcementMode: to.Ptr(armsecurity.EnforcementModeAudit),
				PathRecommendations: []*armsecurity.PathRecommendation{
					{
						Type:                to.Ptr(armsecurity.RecommendationType("PublisherSignature")),
						Path:                to.Ptr("<path>"),
						Action:              to.Ptr(armsecurity.RecommendationActionRecommended),
						Common:              to.Ptr(true),
						ConfigurationStatus: to.Ptr(armsecurity.ConfigurationStatusConfigured),
						FileType:            to.Ptr(armsecurity.FileTypeExe),
						PublisherInfo: &armsecurity.PublisherInfo{
							BinaryName:    to.Ptr("<binary-name>"),
							ProductName:   to.Ptr("<product-name>"),
							PublisherName: to.Ptr("<publisher-name>"),
							Version:       to.Ptr("<version>"),
						},
						UserSids: []*string{
							to.Ptr("S-1-1-0")},
						Usernames: []*armsecurity.UserRecommendation{
							{
								RecommendationAction: to.Ptr(armsecurity.RecommendationActionRecommended),
								Username:             to.Ptr("<username>"),
							}},
					},
					{
						Type:                to.Ptr(armsecurity.RecommendationType("ProductSignature")),
						Path:                to.Ptr("<path>"),
						Action:              to.Ptr(armsecurity.RecommendationActionRecommended),
						Common:              to.Ptr(true),
						ConfigurationStatus: to.Ptr(armsecurity.ConfigurationStatusConfigured),
						FileType:            to.Ptr(armsecurity.FileTypeExe),
						PublisherInfo: &armsecurity.PublisherInfo{
							BinaryName:    to.Ptr("<binary-name>"),
							ProductName:   to.Ptr("<product-name>"),
							PublisherName: to.Ptr("<publisher-name>"),
							Version:       to.Ptr("<version>"),
						},
						UserSids: []*string{
							to.Ptr("S-1-1-0")},
						Usernames: []*armsecurity.UserRecommendation{
							{
								RecommendationAction: to.Ptr(armsecurity.RecommendationActionRecommended),
								Username:             to.Ptr("<username>"),
							}},
					},
					{
						Type:                to.Ptr(armsecurity.RecommendationType("PublisherSignature")),
						Path:                to.Ptr("<path>"),
						Action:              to.Ptr(armsecurity.RecommendationActionRecommended),
						Common:              to.Ptr(true),
						ConfigurationStatus: to.Ptr(armsecurity.ConfigurationStatusConfigured),
						FileType:            to.Ptr(armsecurity.FileTypeExe),
						PublisherInfo: &armsecurity.PublisherInfo{
							BinaryName:    to.Ptr("<binary-name>"),
							ProductName:   to.Ptr("<product-name>"),
							PublisherName: to.Ptr("<publisher-name>"),
							Version:       to.Ptr("<version>"),
						},
						UserSids: []*string{
							to.Ptr("S-1-1-0")},
						Usernames: []*armsecurity.UserRecommendation{
							{
								RecommendationAction: to.Ptr(armsecurity.RecommendationActionRecommended),
								Username:             to.Ptr("<username>"),
							}},
					},
					{
						Type:   to.Ptr(armsecurity.RecommendationType("File")),
						Path:   to.Ptr("<path>"),
						Action: to.Ptr(armsecurity.RecommendationActionAdd),
						Common: to.Ptr(true),
					}},
				ProtectionMode: &armsecurity.ProtectionMode{
					Exe:    to.Ptr(armsecurity.EnforcementModeAudit),
					Msi:    to.Ptr(armsecurity.EnforcementModeNone),
					Script: to.Ptr(armsecurity.EnforcementModeNone),
				},
				VMRecommendations: []*armsecurity.VMRecommendation{
					{
						ConfigurationStatus:  to.Ptr(armsecurity.ConfigurationStatusConfigured),
						EnforcementSupport:   to.Ptr(armsecurity.EnforcementSupportSupported),
						RecommendationAction: to.Ptr(armsecurity.RecommendationActionRecommended),
						ResourceID:           to.Ptr("<resource-id>"),
					},
					{
						ConfigurationStatus:  to.Ptr(armsecurity.ConfigurationStatusConfigured),
						EnforcementSupport:   to.Ptr(armsecurity.EnforcementSupportSupported),
						RecommendationAction: to.Ptr(armsecurity.RecommendationActionRecommended),
						ResourceID:           to.Ptr("<resource-id>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
