package armredhatopenshifthcp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshifthcp/armredhatopenshifthcp"
)

// Generated from example definition: 2026-09-01-preview/HcpOpenShiftClusters_RequestAdminCredential_MaximumSet_Gen.json
func ExampleHcpOpenShiftClustersClient_BeginRequestAdminCredential() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredhatopenshifthcp.NewClientFactory("FDEA43EA-0230-4A7D-BDEE-F3AFF2183B1D", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewHcpOpenShiftClustersClient().BeginRequestAdminCredential(ctx, "rgopenapi", "hcpCluster-name", armredhatopenshifthcp.HcpOpenShiftClusterAdminCredentialRequest{
		CertificateSigningRequest: to.Ptr("-----BEGIN CERTIFICATE REQUEST-----\nMIIBhTCB7wIBADBFMQswCQYDVQQGEwJVUzELMAkGA1UECAwCQ0ExDjAMBgNVBAoM\nBVRlc3QxGTAXBgNVBAMMEHRlc3QuZXhhbXBsZS5jb20wdjAQBgcqhkjOPQIBBgUr\ngQQAIgNiAARIm+7hphQ7m8kzCB5keJ3lPVQvsEH6ABXz0kIvxkNF7+OBFCdPJIBT\nksaGJnJFfPUROYGJIo7FMOO/vEqE9gHqRCVao0RPDaZLtceCYqbeI0vFhW7qTmYL\nNp/RTer7C0+gITAfBgkqhkiG9w0BCQ4xEjAQMA4GA1UdEQQHMAWCA2FiYzAKBggq\nhkjOPQQDAgNoADBlAjBLQDR3K8k1XPFH3Y0oEFYrBi3L4FOX0kz0aK/JuFJN/kBP\nA2ViVNHl+5iVxvpJE5sCMQCF+nPr18qRaib09BHSBKl+ZVpXC1K3PN/VGjYv+Zjl\nK8eCiPwwRBpRMbqMSXxlS3Q=\n-----END CERTIFICATE REQUEST-----\n"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armredhatopenshifthcp.HcpOpenShiftClustersClientRequestAdminCredentialResponse{
	// 	HcpOpenShiftClusterAdminCredential: armredhatopenshifthcp.HcpOpenShiftClusterAdminCredential{
	// 		Kubeconfig: to.Ptr("apiVersion: v1\nclusters:\n- cluster:\n    server: https://api.example.com:6443\n  name: cluster\ncontexts:\n- context:\n    cluster: cluster\n    user: admin\n  name: admin\ncurrent-context: admin\nkind: Config\nusers:\n- name: admin\n  user:\n    client-certificate-data: LS0tLS1C...\n"),
	// 		ExpirationTimestamp: to.Ptr(time.Date(2025, time.April, 23, 5, 55, 13, 791000000, time.UTC)),
	// 	},
	// }
}
