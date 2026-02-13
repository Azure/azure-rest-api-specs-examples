package armedgeactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeactions/armedgeactions"
)

// Generated from example definition: 2025-12-01-preview/EdgeActionVersions_ListByEdgeAction.json
func ExampleEdgeActionVersionsClient_NewListByEdgeActionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armedgeactions.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEdgeActionVersionsClient().NewListByEdgeActionPager("testrg", "edgeAction1", nil)
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
		// page = armedgeactions.EdgeActionVersionsClientListByEdgeActionResponse{
		// 	EdgeActionVersionListResult: armedgeactions.EdgeActionVersionListResult{
		// 		Value: []*armedgeactions.EdgeActionVersion{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Cdn/edgeActions/edgeAction1/versions/version1"),
		// 				Name: to.Ptr("version1"),
		// 				Type: to.Ptr("Microsoft.Cdn/edgeActions/versions"),
		// 				Location: to.Ptr("global"),
		// 				Properties: &armedgeactions.EdgeActionVersionProperties{
		// 					ProvisioningState: to.Ptr(armedgeactions.ProvisioningStateSucceeded),
		// 					DeploymentType: to.Ptr(armedgeactions.EdgeActionVersionDeploymentTypeZip),
		// 					ValidationStatus: to.Ptr(armedgeactions.EdgeActionVersionValidationStatusSucceeded),
		// 					IsDefaultVersion: to.Ptr(armedgeactions.EdgeActionIsDefaultVersionTrue),
		// 					LastPackageUpdateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-05-25T16:19:23Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Cdn/edgeActions/edgeAction1/versions/version2"),
		// 				Name: to.Ptr("version2"),
		// 				Type: to.Ptr("Microsoft.Cdn/edgeActions/versions"),
		// 				Location: to.Ptr("global"),
		// 				Properties: &armedgeactions.EdgeActionVersionProperties{
		// 					ProvisioningState: to.Ptr(armedgeactions.ProvisioningStateSucceeded),
		// 					DeploymentType: to.Ptr(armedgeactions.EdgeActionVersionDeploymentTypeZip),
		// 					ValidationStatus: to.Ptr(armedgeactions.EdgeActionVersionValidationStatusSucceeded),
		// 					IsDefaultVersion: to.Ptr(armedgeactions.EdgeActionIsDefaultVersionFalse),
		// 					LastPackageUpdateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-05-25T11:19:23Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
