Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.7.0/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/CreateIoTSecuritySolution.json
func ExampleIotSecuritySolutionClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewIotSecuritySolutionClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"MyGroup",
		"default",
		armsecurity.IoTSecuritySolutionModel{
			Tags:     map[string]*string{},
			Location: to.Ptr("East Us"),
			Properties: &armsecurity.IoTSecuritySolutionProperties{
				DisabledDataSources: []*armsecurity.DataSource{},
				DisplayName:         to.Ptr("Solution Default"),
				Export:              []*armsecurity.ExportData{},
				IotHubs: []*string{
					to.Ptr("/subscriptions/075423e9-7d33-4166-8bdf-3920b04e3735/resourceGroups/myRg/providers/Microsoft.Devices/IotHubs/FirstIotHub")},
				RecommendationsConfiguration: []*armsecurity.RecommendationConfigurationProperties{
					{
						RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTOpenPorts),
						Status:             to.Ptr(armsecurity.RecommendationConfigStatusDisabled),
					},
					{
						RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTSharedCredentials),
						Status:             to.Ptr(armsecurity.RecommendationConfigStatusDisabled),
					}},
				Status:                  to.Ptr(armsecurity.SecuritySolutionStatusEnabled),
				UnmaskedIPLoggingStatus: to.Ptr(armsecurity.UnmaskedIPLoggingStatusEnabled),
				UserDefinedResources: &armsecurity.UserDefinedResourcesProperties{
					Query: to.Ptr("where type != \"microsoft.devices/iothubs\" | where name contains \"iot\""),
					QuerySubscriptions: []*string{
						to.Ptr("075423e9-7d33-4166-8bdf-3920b04e3735")},
				},
				Workspace: to.Ptr("/subscriptions/c4930e90-cd72-4aa5-93e9-2d081d129569/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace1"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
