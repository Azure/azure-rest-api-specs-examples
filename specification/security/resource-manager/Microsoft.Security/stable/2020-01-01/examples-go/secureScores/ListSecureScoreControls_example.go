package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/secureScores/ListSecureScoreControls_example.json
func ExampleSecureScoreControlsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSecureScoreControlsClient().NewListPager(&armsecurity.SecureScoreControlsClientListOptions{Expand: nil})
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
		// page.SecureScoreControlList = armsecurity.SecureScoreControlList{
		// 	Value: []*armsecurity.SecureScoreControlDetails{
		// 		{
		// 			Name: to.Ptr("a9909064-42b4-4d34-8143-275477afe18b"),
		// 			Type: to.Ptr("Microsoft.Security/secureScores/secureScoreControls"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/secureScores/ascScore/secureScoreControls/a9909064-42b4-4d34-8143-275477afe18b"),
		// 			Properties: &armsecurity.SecureScoreControlScoreDetails{
		// 				DisplayName: to.Ptr("Protect applications against DDoS attacks"),
		// 				HealthyResourceCount: to.Ptr[int32](0),
		// 				NotApplicableResourceCount: to.Ptr[int32](1),
		// 				Score: &armsecurity.ScoreDetails{
		// 					Current: to.Ptr[float64](0),
		// 					Max: to.Ptr[int32](0),
		// 					Percentage: to.Ptr[float64](0),
		// 				},
		// 				UnhealthyResourceCount: to.Ptr[int32](0),
		// 				Weight: to.Ptr[int64](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("34a42fb3-e6db-409c-b56b-7b1db6b8aee0"),
		// 			Type: to.Ptr("Microsoft.Security/secureScores/secureScoreControls"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/secureScores/ascScore/secureScoreControls/34a42fb3-e6db-409c-b56b-7b1db6b8aee0"),
		// 			Properties: &armsecurity.SecureScoreControlScoreDetails{
		// 				DisplayName: to.Ptr("Enable MFA"),
		// 				HealthyResourceCount: to.Ptr[int32](1),
		// 				NotApplicableResourceCount: to.Ptr[int32](0),
		// 				Score: &armsecurity.ScoreDetails{
		// 					Current: to.Ptr[float64](10),
		// 					Max: to.Ptr[int32](10),
		// 					Percentage: to.Ptr[float64](1),
		// 				},
		// 				UnhealthyResourceCount: to.Ptr[int32](0),
		// 				Weight: to.Ptr[int64](1),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("76763537-9feb-42d3-b4f4-78c01117be80"),
		// 			Type: to.Ptr("Microsoft.Security/secureScores/secureScoreControls"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/secureScores/ascScore/secureScoreControls/76763537-9feb-42d3-b4f4-78c01117be80"),
		// 			Properties: &armsecurity.SecureScoreControlScoreDetails{
		// 				DisplayName: to.Ptr("Implement security best practices"),
		// 				HealthyResourceCount: to.Ptr[int32](7),
		// 				NotApplicableResourceCount: to.Ptr[int32](1),
		// 				Score: &armsecurity.ScoreDetails{
		// 					Current: to.Ptr[float64](0),
		// 					Max: to.Ptr[int32](0),
		// 					Percentage: to.Ptr[float64](0),
		// 				},
		// 				UnhealthyResourceCount: to.Ptr[int32](0),
		// 				Weight: to.Ptr[int64](7),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("4db8d6cf-075b-4149-a813-da09ca2ae120"),
		// 			Type: to.Ptr("Microsoft.Security/secureScores/secureScoreControls"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/secureScores/ascScore/secureScoreControls/4db8d6cf-075b-4149-a813-da09ca2ae120"),
		// 			Properties: &armsecurity.SecureScoreControlScoreDetails{
		// 				DisplayName: to.Ptr("Enable Advanced Threat Protection"),
		// 				HealthyResourceCount: to.Ptr[int32](0),
		// 				NotApplicableResourceCount: to.Ptr[int32](0),
		// 				Score: &armsecurity.ScoreDetails{
		// 					Current: to.Ptr[float64](0),
		// 					Max: to.Ptr[int32](0),
		// 					Percentage: to.Ptr[float64](0),
		// 				},
		// 				UnhealthyResourceCount: to.Ptr[int32](1),
		// 				Weight: to.Ptr[int64](1),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("27b24961-75ba-4fe4-8909-97286d5dd5ee"),
		// 			Type: to.Ptr("Microsoft.Security/secureScores/secureScoreControls"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/secureScores/ascScore/secureScoreControls/27b24961-75ba-4fe4-8909-97286d5dd5ee"),
		// 			Properties: &armsecurity.SecureScoreControlScoreDetails{
		// 				DisplayName: to.Ptr("Enable auditing and logging"),
		// 				HealthyResourceCount: to.Ptr[int32](1),
		// 				NotApplicableResourceCount: to.Ptr[int32](0),
		// 				Score: &armsecurity.ScoreDetails{
		// 					Current: to.Ptr[float64](0.2),
		// 					Max: to.Ptr[int32](1),
		// 					Percentage: to.Ptr[float64](0.2),
		// 				},
		// 				UnhealthyResourceCount: to.Ptr[int32](4),
		// 				Weight: to.Ptr[int64](5),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("2d85f639-0bea-4a4a-b6c6-608952a1414a"),
		// 			Type: to.Ptr("Microsoft.Security/secureScores/secureScoreControls"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/secureScores/ascScore/secureScoreControls/2d85f639-0bea-4a4a-b6c6-608952a1414a"),
		// 			Properties: &armsecurity.SecureScoreControlScoreDetails{
		// 				DisplayName: to.Ptr("Manage access and permissions"),
		// 				HealthyResourceCount: to.Ptr[int32](1),
		// 				NotApplicableResourceCount: to.Ptr[int32](0),
		// 				Score: &armsecurity.ScoreDetails{
		// 					Current: to.Ptr[float64](2),
		// 					Max: to.Ptr[int32](4),
		// 					Percentage: to.Ptr[float64](0.5),
		// 				},
		// 				UnhealthyResourceCount: to.Ptr[int32](1),
		// 				Weight: to.Ptr[int64](2),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("99fc8df2-e0f5-40f8-9415-a7f7ca948b5a"),
		// 			Type: to.Ptr("Microsoft.Security/secureScores/secureScoreControls"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/secureScores/ascScore/secureScoreControls/99fc8df2-e0f5-40f8-9415-a7f7ca948b5a"),
		// 			Properties: &armsecurity.SecureScoreControlScoreDetails{
		// 				DisplayName: to.Ptr("Remediate security configurations"),
		// 				HealthyResourceCount: to.Ptr[int32](0),
		// 				NotApplicableResourceCount: to.Ptr[int32](0),
		// 				Score: &armsecurity.ScoreDetails{
		// 					Current: to.Ptr[float64](0),
		// 					Max: to.Ptr[int32](4),
		// 					Percentage: to.Ptr[float64](0),
		// 				},
		// 				UnhealthyResourceCount: to.Ptr[int32](1),
		// 				Weight: to.Ptr[int64](1),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("12136bd9-dc24-44f2-9587-7be3af6aac14"),
		// 			Type: to.Ptr("Microsoft.Security/secureScores/secureScoreControls"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/secureScores/ascScore/secureScoreControls/12136bd9-dc24-44f2-9587-7be3af6aac14"),
		// 			Properties: &armsecurity.SecureScoreControlScoreDetails{
		// 				DisplayName: to.Ptr("Enable endpoint protection"),
		// 				HealthyResourceCount: to.Ptr[int32](0),
		// 				NotApplicableResourceCount: to.Ptr[int32](0),
		// 				Score: &armsecurity.ScoreDetails{
		// 					Current: to.Ptr[float64](0),
		// 					Max: to.Ptr[int32](2),
		// 					Percentage: to.Ptr[float64](0),
		// 				},
		// 				UnhealthyResourceCount: to.Ptr[int32](1),
		// 				Weight: to.Ptr[int64](1),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("61702b76-1fab-41f2-bcbc-50b7870dcf38"),
		// 			Type: to.Ptr("Microsoft.Security/secureScores/secureScoreControls"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/secureScores/ascScore/secureScoreControls/61702b76-1fab-41f2-bcbc-50b7870dcf38"),
		// 			Properties: &armsecurity.SecureScoreControlScoreDetails{
		// 				DisplayName: to.Ptr("Apply system updates"),
		// 				HealthyResourceCount: to.Ptr[int32](0),
		// 				NotApplicableResourceCount: to.Ptr[int32](0),
		// 				Score: &armsecurity.ScoreDetails{
		// 					Current: to.Ptr[float64](0),
		// 					Max: to.Ptr[int32](6),
		// 					Percentage: to.Ptr[float64](0),
		// 				},
		// 				UnhealthyResourceCount: to.Ptr[int32](1),
		// 				Weight: to.Ptr[int64](1),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("0e55495e-034f-4abc-8293-767229250176"),
		// 			Type: to.Ptr("Microsoft.Security/secureScores/secureScoreControls"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/secureScores/ascScore/secureScoreControls/0e55495e-034f-4abc-8293-767229250176"),
		// 			Properties: &armsecurity.SecureScoreControlScoreDetails{
		// 				DisplayName: to.Ptr("Encrypt data in transit"),
		// 				HealthyResourceCount: to.Ptr[int32](5),
		// 				NotApplicableResourceCount: to.Ptr[int32](0),
		// 				Score: &armsecurity.ScoreDetails{
		// 					Current: to.Ptr[float64](3.33),
		// 					Max: to.Ptr[int32](4),
		// 					Percentage: to.Ptr[float64](0.8325),
		// 				},
		// 				UnhealthyResourceCount: to.Ptr[int32](1),
		// 				Weight: to.Ptr[int64](6),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("f9d5432b-8f7b-45e9-b90c-e214a30f6a02"),
		// 			Type: to.Ptr("Microsoft.Security/secureScores/secureScoreControls"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/secureScores/ascScore/secureScoreControls/f9d5432b-8f7b-45e9-b90c-e214a30f6a02"),
		// 			Properties: &armsecurity.SecureScoreControlScoreDetails{
		// 				DisplayName: to.Ptr("Restrict unauthorized network access"),
		// 				HealthyResourceCount: to.Ptr[int32](1),
		// 				NotApplicableResourceCount: to.Ptr[int32](0),
		// 				Score: &armsecurity.ScoreDetails{
		// 					Current: to.Ptr[float64](4),
		// 					Max: to.Ptr[int32](4),
		// 					Percentage: to.Ptr[float64](1),
		// 				},
		// 				UnhealthyResourceCount: to.Ptr[int32](0),
		// 				Weight: to.Ptr[int64](1),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("8fd574ec-43cf-426e-a439-a67cbaf2d564"),
		// 			Type: to.Ptr("Microsoft.Security/secureScores/secureScoreControls"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/secureScores/ascScore/secureScoreControls/8fd574ec-43cf-426e-a439-a67cbaf2d564"),
		// 			Properties: &armsecurity.SecureScoreControlScoreDetails{
		// 				DisplayName: to.Ptr("Enable encryption at rest"),
		// 				HealthyResourceCount: to.Ptr[int32](1),
		// 				NotApplicableResourceCount: to.Ptr[int32](0),
		// 				Score: &armsecurity.ScoreDetails{
		// 					Current: to.Ptr[float64](4),
		// 					Max: to.Ptr[int32](4),
		// 					Percentage: to.Ptr[float64](1),
		// 				},
		// 				UnhealthyResourceCount: to.Ptr[int32](0),
		// 				Weight: to.Ptr[int64](1),
		// 			},
		// 	}},
		// }
	}
}
