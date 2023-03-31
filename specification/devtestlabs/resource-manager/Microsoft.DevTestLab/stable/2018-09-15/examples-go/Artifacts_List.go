package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Artifacts_List.json
func ExampleArtifactsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewArtifactsClient().NewListPager("resourceGroupName", "{labName}", "{artifactSourceName}", &armdevtestlabs.ArtifactsClientListOptions{Expand: nil,
		Filter:  nil,
		Top:     nil,
		Orderby: nil,
	})
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
		// page.ArtifactList = armdevtestlabs.ArtifactList{
		// 	Value: []*armdevtestlabs.Artifact{
		// 		{
		// 			Name: to.Ptr("{artifactName}"),
		// 			Type: to.Ptr("Microsoft.DevTestLab/labs/artifactSources/artifacts"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.DevTestLab/labs/{labName}/artifactSources/{artifactSourceName}/artifacts/{artifactName}"),
		// 			Location: to.Ptr("{location}"),
		// 			Tags: map[string]*string{
		// 				"MyTag": to.Ptr("MyValue"),
		// 			},
		// 			Properties: &armdevtestlabs.ArtifactProperties{
		// 				Description: to.Ptr("Sample artifact description."),
		// 				FilePath: to.Ptr("{artifactsPath}/{artifactName}"),
		// 				Parameters: map[string]any{
		// 					"uri":map[string]any{
		// 						"type": "string",
		// 						"description": "Sample parameter 1 description.",
		// 						"defaultValue": "https://{labStorageAccount}.blob.core.windows.net/{artifactName}/...",
		// 						"displayName": "Sample Parameter 1",
		// 					},
		// 				},
		// 				Publisher: to.Ptr("Microsoft"),
		// 				TargetOsType: to.Ptr("Windows"),
		// 				Title: to.Ptr("Sample Artifact Title"),
		// 			},
		// 	}},
		// }
	}
}
