package armartifactsigning_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/artifactsigning/armartifactsigning"
)

// Generated from example definition: 2026-05-15-preview/CertificateProfiles_RevokeCertificates.json
func ExampleCertificateProfilesClient_RevokeCertificates() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armartifactsigning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewCertificateProfilesClient().RevokeCertificates(ctx, "MyResourceGroup", "MyAccount", "profileA", armartifactsigning.RevokeCertificateList{
		RevokeCertificates: []*armartifactsigning.RevokeCertificate{
			{
				EffectiveAt:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-12T23:40:25+00:00"); return t }()),
				Reason:       to.Ptr("KeyCompromised"),
				Remarks:      to.Ptr("test"),
				SerialNumber: to.Ptr("xxxxxxxxxxxxxxxxxx"),
				Thumbprint:   to.Ptr("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			},
			{
				EffectiveAt:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-12T23:40:25+00:00"); return t }()),
				Reason:       to.Ptr("KeyCompromised"),
				Remarks:      to.Ptr("test"),
				SerialNumber: to.Ptr("yyyyyyyyyyyyyyyyyy"),
				Thumbprint:   to.Ptr("yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
