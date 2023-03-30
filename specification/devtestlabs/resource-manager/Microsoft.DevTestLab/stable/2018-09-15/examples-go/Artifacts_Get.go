package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Artifacts_Get.json
func ExampleArtifactsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewArtifactsClient().Get(ctx, "resourceGroupName", "{labName}", "{artifactSourceName}", "{artifactName}", &armdevtestlabs.ArtifactsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Artifact = armdevtestlabs.Artifact{
	// 	Name: to.Ptr("{artifactName}"),
	// 	Type: to.Ptr("Microsoft.DevTestLab/labs/artifactSources/artifacts"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.DevTestLab/labs/{labName}/artifactSources/{artifactSourceName}/artifacts/{artifactName}"),
	// 	Location: to.Ptr("{location}"),
	// 	Tags: map[string]*string{
	// 		"MyTag": to.Ptr("MyValue"),
	// 	},
	// 	Properties: &armdevtestlabs.ArtifactProperties{
	// 		Description: to.Ptr("Sample artifact description."),
	// 		FilePath: to.Ptr("{artifactsPath}/{artifactName}"),
	// 		Parameters: map[string]any{
	// 			"uri":map[string]any{
	// 				"type": "string",
	// 				"description": "Sample parameter 1 description.",
	// 				"defaultValue": "https://{labStorageAccount}.blob.core.windows.net/{artifactName}/...",
	// 				"displayName": "Sample Parameter 1",
	// 			},
	// 		},
	// 		Publisher: to.Ptr("Microsoft"),
	// 		TargetOsType: to.Ptr("Windows"),
	// 		Title: to.Ptr("Sample Artifact Title"),
	// 	},
	// }
}
