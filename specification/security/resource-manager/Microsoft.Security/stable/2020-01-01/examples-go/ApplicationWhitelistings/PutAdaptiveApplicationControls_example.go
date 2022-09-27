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
	}
	ctx := context.Background()
	client, err := armsecurity.NewAdaptiveApplicationControlsClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Put(ctx, "centralus", "ERELGROUP1", armsecurity.AdaptiveApplicationControlGroup{
		Properties: &armsecurity.AdaptiveApplicationControlGroupData{
			EnforcementMode: to.Ptr(armsecurity.EnforcementModeAudit),
			PathRecommendations: []*armsecurity.PathRecommendation{
				{
					Type:                to.Ptr(armsecurity.RecommendationType("PublisherSignature")),
					Path:                to.Ptr("[Exe] O=MICROSOFT CORPORATION, L=REDMOND, S=WASHINGTON, C=US\\*\\*\\0.0.0.0"),
					Action:              to.Ptr(armsecurity.RecommendationActionRecommended),
					Common:              to.Ptr(true),
					ConfigurationStatus: to.Ptr(armsecurity.ConfigurationStatusConfigured),
					FileType:            to.Ptr(armsecurity.FileTypeExe),
					PublisherInfo: &armsecurity.PublisherInfo{
						BinaryName:    to.Ptr("*"),
						ProductName:   to.Ptr("*"),
						PublisherName: to.Ptr("O=MICROSOFT CORPORATION, L=REDMOND, S=WASHINGTON, C=US"),
						Version:       to.Ptr("0.0.0.0"),
					},
					UserSids: []*string{
						to.Ptr("S-1-1-0")},
					Usernames: []*armsecurity.UserRecommendation{
						{
							RecommendationAction: to.Ptr(armsecurity.RecommendationActionRecommended),
							Username:             to.Ptr("Everyone"),
						}},
				},
				{
					Type:                to.Ptr(armsecurity.RecommendationType("ProductSignature")),
					Path:                to.Ptr("%OSDRIVE%\\WINDOWSAZURE\\SECAGENT\\WASECAGENTPROV.EXE"),
					Action:              to.Ptr(armsecurity.RecommendationActionRecommended),
					Common:              to.Ptr(true),
					ConfigurationStatus: to.Ptr(armsecurity.ConfigurationStatusConfigured),
					FileType:            to.Ptr(armsecurity.FileTypeExe),
					PublisherInfo: &armsecurity.PublisherInfo{
						BinaryName:    to.Ptr("*"),
						ProductName:   to.Ptr("MICROSOFT® COREXT"),
						PublisherName: to.Ptr("CN=MICROSOFT AZURE DEPENDENCY CODE SIGN"),
						Version:       to.Ptr("0.0.0.0"),
					},
					UserSids: []*string{
						to.Ptr("S-1-1-0")},
					Usernames: []*armsecurity.UserRecommendation{
						{
							RecommendationAction: to.Ptr(armsecurity.RecommendationActionRecommended),
							Username:             to.Ptr("NT AUTHORITY\\SYSTEM"),
						}},
				},
				{
					Type:                to.Ptr(armsecurity.RecommendationType("PublisherSignature")),
					Path:                to.Ptr("%OSDRIVE%\\WINDOWSAZURE\\PACKAGES_201973_7415\\COLLECTGUESTLOGS.EXE"),
					Action:              to.Ptr(armsecurity.RecommendationActionRecommended),
					Common:              to.Ptr(true),
					ConfigurationStatus: to.Ptr(armsecurity.ConfigurationStatusConfigured),
					FileType:            to.Ptr(armsecurity.FileTypeExe),
					PublisherInfo: &armsecurity.PublisherInfo{
						BinaryName:    to.Ptr("*"),
						ProductName:   to.Ptr("*"),
						PublisherName: to.Ptr("CN=MICROSOFT AZURE DEPENDENCY CODE SIGN"),
						Version:       to.Ptr("0.0.0.0"),
					},
					UserSids: []*string{
						to.Ptr("S-1-1-0")},
					Usernames: []*armsecurity.UserRecommendation{
						{
							RecommendationAction: to.Ptr(armsecurity.RecommendationActionRecommended),
							Username:             to.Ptr("NT AUTHORITY\\SYSTEM"),
						}},
				},
				{
					Type:   to.Ptr(armsecurity.RecommendationType("File")),
					Path:   to.Ptr("C:\\directory\\file.exe"),
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
					ResourceID:           to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/erelh-stable/providers/microsoft.compute/virtualmachines/erelh-16090"),
				},
				{
					ConfigurationStatus:  to.Ptr(armsecurity.ConfigurationStatusConfigured),
					EnforcementSupport:   to.Ptr(armsecurity.EnforcementSupportSupported),
					RecommendationAction: to.Ptr(armsecurity.RecommendationActionRecommended),
					ResourceID:           to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/matanvs/providers/microsoft.compute/virtualmachines/matanvs19"),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
