package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: 2019-08-01/IoTSecuritySolutionsAnalytics/GetIoTSecuritySolutionsSecurityRecommendation.json
func ExampleIotSecuritySolutionsAnalyticsRecommendationClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("075423e9-7d33-4166-8bdf-3920b04e3735", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIotSecuritySolutionsAnalyticsRecommendationClient().Get(ctx, "IoTEdgeResources", "default", "OpenPortsOnDevice", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurity.IotSecuritySolutionsAnalyticsRecommendationClientGetResponse{
	// 	IoTSecurityAggregatedRecommendation: armsecurity.IoTSecurityAggregatedRecommendation{
	// 		Name: to.Ptr("OpenPortsOnDevice"),
	// 		Type: to.Ptr("Microsoft.Security/iotSecuritySolutions/analyticsModels/aggregatedRecommendations"),
	// 		ID: to.Ptr("/subscriptions/075423e9-7d33-4166-8bdf-3920b04e3735/resourceGroups/IoTEdgeResources/providers/Microsoft.Security/iotSecuritySolutions/default/analyticsModels/default/aggregatedRecommendations/OpenPortsOnDevice"),
	// 		Properties: &armsecurity.IoTSecurityAggregatedRecommendationProperties{
	// 			Description: to.Ptr("An allowed firewall policy was found in main firewall Chains (INPUT/OUTPUT). The policy should Deny all traffic by default define rules to allow necessary communication to/from the device"),
	// 			DetectedBy: to.Ptr("Microsoft"),
	// 			HealthyDevices: to.Ptr[int64](10000),
	// 			LogAnalyticsQuery: to.Ptr("SecurityRecommendation | where tolower(AssessedResourceId) == tolower('/subscriptions/075423e9-7d33-4166-8bdf-3920b04e3735/resourceGroups/IoTEdgeResources/providers/Microsoft.Devices/IotHubs/t-ofdadu-hub') and tolower(RecommendationName) == tolower('OpenPortsOnDevice')"),
	// 			RecommendationDisplayName: to.Ptr("Permissive firewall policy in one of the chains was found"),
	// 			RecommendationName: to.Ptr("OpenPortsOnDevice"),
	// 			RecommendationTypeID: to.Ptr("{20ff7fc3-e762-44dd-bd96-b71116dcdc23}"),
	// 			RemediationSteps: to.Ptr(""),
	// 			ReportedSeverity: to.Ptr(armsecurity.ReportedSeverityLow),
	// 			UnhealthyDeviceCount: to.Ptr[int64](200),
	// 		},
	// 	},
	// }
}
