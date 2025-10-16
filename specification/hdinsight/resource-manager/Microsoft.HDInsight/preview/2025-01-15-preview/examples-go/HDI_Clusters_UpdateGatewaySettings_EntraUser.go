package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a9dbb28e788355a47dc5bad3ea5f8da212b4bf6/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2025-01-15-preview/examples/HDI_Clusters_UpdateGatewaySettings_EntraUser.json
func ExampleClustersClient_BeginUpdateGatewaySettings_updateEntraUserInHdInsight() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginUpdateGatewaySettings(ctx, "rg1", "cluster1", armhdinsight.UpdateGatewaySettingsParameters{
		IsCredentialEnabled: to.Ptr(false),
		RestAuthEntraUsers: []*armhdinsight.EntraUserInfo{
			{
				DisplayName: to.Ptr("displayName"),
				ObjectID:    to.Ptr("00000000-0000-0000-0000-000000000000"),
				Upn:         to.Ptr("user@microsoft.com"),
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
