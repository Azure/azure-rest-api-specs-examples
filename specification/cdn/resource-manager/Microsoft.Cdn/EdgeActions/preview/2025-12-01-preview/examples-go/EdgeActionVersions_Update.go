package armedgeactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeactions/armedgeactions"
)

// Generated from example definition: 2025-12-01-preview/EdgeActionVersions_Update.json
func ExampleEdgeActionVersionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armedgeactions.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEdgeActionVersionsClient().BeginUpdate(ctx, "testrg", "edgeAction1", "version1", armedgeactions.EdgeActionVersionUpdate{
		Properties: &armedgeactions.EdgeActionVersionUpdateProperties{
			DeploymentType: to.Ptr(armedgeactions.EdgeActionVersionDeploymentTypeOthers),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armedgeactions.EdgeActionVersionsClientUpdateResponse{
	// 	EdgeActionVersion: &armedgeactions.EdgeActionVersion{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Cdn/edgeActions/edgeAction1/versions/version1"),
	// 		Name: to.Ptr("version1"),
	// 		Type: to.Ptr("Microsoft.Cdn/edgeActions/versions"),
	// 		Location: to.Ptr("global"),
	// 		Properties: &armedgeactions.EdgeActionVersionProperties{
	// 			ProvisioningState: to.Ptr(armedgeactions.ProvisioningStateSucceeded),
	// 			DeploymentType: to.Ptr(armedgeactions.EdgeActionVersionDeploymentTypeOthers),
	// 			ValidationStatus: to.Ptr(armedgeactions.EdgeActionVersionValidationStatusSucceeded),
	// 			IsDefaultVersion: to.Ptr(armedgeactions.EdgeActionIsDefaultVersionTrue),
	// 			LastPackageUpdateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-05-25T16:19:23Z"); return t}()),
	// 		},
	// 	},
	// }
}
