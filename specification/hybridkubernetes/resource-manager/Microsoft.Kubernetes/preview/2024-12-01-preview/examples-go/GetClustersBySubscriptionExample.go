package armhybridkubernetes_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridkubernetes/armhybridkubernetes/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ecee919199a39cc0d864410f540aa105bf7cdb64/specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/preview/2024-12-01-preview/examples/GetClustersBySubscriptionExample.json
func ExampleConnectedClusterClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridkubernetes.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectedClusterClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ConnectedClusterList = armhybridkubernetes.ConnectedClusterList{
		// 	Value: []*armhybridkubernetes.ConnectedCluster{
		// 		{
		// 			Name: to.Ptr("connectedCluster1"),
		// 			Type: to.Ptr("Microsoft.Kubernetes/connectedClusters"),
		// 			ID: to.Ptr("/subscriptions/1bfbb5d0-917e-4346-9026-1d3b344417f5/resourceGroups/akkeshar/providers/Microsoft.Kubernetes/connectedClusters/connectedCluster1"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armhybridkubernetes.ConnectedClusterIdentity{
		// 				Type: to.Ptr(armhybridkubernetes.ResourceIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("d8cd1fd9-154f-4da7-b348-595f283c13a3"),
		// 				TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			},
		// 			Properties: &armhybridkubernetes.ConnectedClusterProperties{
		// 				AgentPublicKeyCertificate: to.Ptr("MIICYzCCAcygAwIBAgIBADANBgkqhkiG9w0BAQUFADAuMQswCQYDVQQGEwJVUzEMMAoGA1UEChMDSUJNMREwDwYDVQQLEwhMb2NhbCBDQTAeFw05OTEyMjIwNTAwMDBaFw0wMDEyMjMwNDU5NTlaMC4xCzAJBgNVBAYTAlVTMQwwCgYDVQQKEwNJQk0xETAPBgNVBAsTCExvY2FsIENBMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQD2bZEo7xGaX2/0GHkrNFZvlxBou9v1Jmt/PDiTMPve8r9FeJAQ0QdvFST/0JPQYD20rH0bimdDLgNdNynmyRoS2S/IInfpmf69iyc2G0TPyRvmHIiOZbdCd+YBHQi1adkj17NDcWj6S14tVurFX73zx0sNoMS79q3tuXKrDsxeuwIDAQABo4GQMIGNMEsGCVUdDwGG+EIBDQQ+EzxHZW5lcmF0ZWQgYnkgdGhlIFNlY3VyZVdheSBTZWN1cml0eSBTZXJ2ZXIgZm9yIE9TLzM5MCAoUkFDRikwDgYDVR0PAQH/BAQDAgAGMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFJ3+ocRyCTJw067dLSwr/nalx6YMMA0GCSqGSIb3DQEBBQUAA4GBAMaQzt+zaj1GU77yzlr8iiMBXgdQrwsZZWJo5exnAucJAEYQZmOfyLiM D6oYq+ZnfvM0n8G/Y79q8nhwvuxpYOnRSAXFp6xSkrIOeZtJMY1h00LKp/JX3Ng1svZ2agE126JHsQ0bhzN5TKsYfbwfTwfjdWAGy6Vf1nYi/rO+ryMO"),
		// 				AgentVersion: to.Ptr("0.1.0"),
		// 				AzureHybridBenefit: to.Ptr(armhybridkubernetes.AzureHybridBenefitNotApplicable),
		// 				Distribution: to.Ptr("AKS"),
		// 				DistributionVersion: to.Ptr("1.0"),
		// 				KubernetesVersion: to.Ptr("1.17.0"),
		// 				MiscellaneousProperties: map[string]*string{
		// 					"createCorrelationId": to.Ptr("8e42616e-74dc-4117-9aaa-71ad5d0180d8"),
		// 					"infrastructureVersion": to.Ptr("1.1"),
		// 				},
		// 				PrivateLinkState: to.Ptr(armhybridkubernetes.PrivateLinkStateDisabled),
		// 				ProvisioningState: to.Ptr(armhybridkubernetes.ProvisioningStateSucceeded),
		// 				TotalNodeCount: to.Ptr[int32](2),
		// 			},
		// 			SystemData: &armhybridkubernetes.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-17T07:06:33.917Z"); return t}()),
		// 				CreatedBy: to.Ptr("sikasire@microsoft.com"),
		// 				CreatedByType: to.Ptr(armhybridkubernetes.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-17T07:06:33.917Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("sikasire@microsoft.com"),
		// 				LastModifiedByType: to.Ptr(armhybridkubernetes.LastModifiedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("connectedCluster2"),
		// 			Type: to.Ptr("Microsoft.Kubernetes/connectedClusters"),
		// 			ID: to.Ptr("/subscriptions/1bfbb5d0-917e-4346-9026-1d3b344417f5/resourceGroups/akkeshar/providers/Microsoft.Kubernetes/connectedClusters/connectedCluster2"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armhybridkubernetes.ConnectedClusterIdentity{
		// 				Type: to.Ptr(armhybridkubernetes.ResourceIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("d8cd1fd9-154f-4da7-b348-595f283c13a3"),
		// 				TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			},
		// 			Properties: &armhybridkubernetes.ConnectedClusterProperties{
		// 				AgentPublicKeyCertificate: to.Ptr("XIICYzCCAcygAwIBAgIBADANBgkqhkiG9w0BAQUFADAuMQswCQYDVQQGEwJVUzEMMAoGA1UEChMDSUJNMREwDwYDVQQLEwhMb2NhbCBDQTAeFw05OTEyMjIwNTAwMDBaFw0wMDEyMjMwNDU5NTlaMC4xCzAJBgNVBAYTAlVTMQwwCgYDVQQKEwNJQk0xETAPBgNVBAsTCExvY2FsIENBMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQD2bZEo7xGaX2/0GHkrNFZvlxBou9v1Jmt/PDiTMPve8r9FeJAQ0QdvFST/0JPQYD20rH0bimdDLgNdNynmyRoS2S/IInfpmf69iyc2G0TPyRvmHIiOZbdCd+YBHQi1adkj17NDcWj6S14tVurFX73zx0sNoMS79q3tuXKrDsxeuwIDAQABo4GQMIGNMEsGCVUdDwGG+EIBDQQ+EzxHZW5lcmF0ZWQgYnkgdGhlIFNlY3VyZVdheSBTZWN1cml0eSBTZXJ2ZXIgZm9yIE9TLzM5MCAoUkFDRikwDgYDVR0PAQH/BAQDAgAGMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFJ3+ocRyCTJw067dLSwr/nalx6YMMA0GCSqGSIb3DQEBBQUAA4GBAMaQzt+zaj1GU77yzlr8iiMBXgdQrwsZZWJo5exnAucJAEYQZmOfyLiM D6oYq+ZnfvM0n8G/Y79q8nhwvuxpYOnRSAXFp6xSkrIOeZtJMY1h00LKp/JX3Ng1svZ2agE126JHsQ0bhzN5TKsYfbwfTwfjdWAGy6Vf1nYi/rO+ryMO"),
		// 				AgentVersion: to.Ptr("0.1.0"),
		// 				AzureHybridBenefit: to.Ptr(armhybridkubernetes.AzureHybridBenefitNotApplicable),
		// 				Distribution: to.Ptr("AKS"),
		// 				DistributionVersion: to.Ptr("1.0"),
		// 				KubernetesVersion: to.Ptr("1.16.3"),
		// 				MiscellaneousProperties: map[string]*string{
		// 					"createCorrelationId": to.Ptr("8e42616e-74dc-4117-9aaa-71ad5d0180d8"),
		// 					"infrastructureVersion": to.Ptr("1.1"),
		// 				},
		// 				PrivateLinkScopeResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName"),
		// 				PrivateLinkState: to.Ptr(armhybridkubernetes.PrivateLinkStateEnabled),
		// 				ProvisioningState: to.Ptr(armhybridkubernetes.ProvisioningStateSucceeded),
		// 				TotalNodeCount: to.Ptr[int32](4),
		// 			},
		// 			SystemData: &armhybridkubernetes.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-17T07:06:33.917Z"); return t}()),
		// 				CreatedBy: to.Ptr("sikasire@microsoft.com"),
		// 				CreatedByType: to.Ptr(armhybridkubernetes.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-17T07:06:33.917Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("sikasire@microsoft.com"),
		// 				LastModifiedByType: to.Ptr(armhybridkubernetes.LastModifiedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
