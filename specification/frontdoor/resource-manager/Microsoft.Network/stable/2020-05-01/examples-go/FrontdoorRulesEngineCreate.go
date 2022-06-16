package armfrontdoor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorRulesEngineCreate.json
func ExampleRulesEnginesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armfrontdoor.NewRulesEnginesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"frontDoor1",
		"rulesEngine1",
		armfrontdoor.RulesEngine{
			Properties: &armfrontdoor.RulesEngineProperties{
				Rules: []*armfrontdoor.RulesEngineRule{
					{
						Name: to.Ptr("Rule1"),
						Action: &armfrontdoor.RulesEngineAction{
							RouteConfigurationOverride: &armfrontdoor.RedirectConfiguration{
								ODataType:         to.Ptr("#Microsoft.Azure.FrontDoor.Models.FrontdoorRedirectConfiguration"),
								CustomFragment:    to.Ptr("fragment"),
								CustomHost:        to.Ptr("www.bing.com"),
								CustomPath:        to.Ptr("/api"),
								CustomQueryString: to.Ptr("a=b"),
								RedirectProtocol:  to.Ptr(armfrontdoor.FrontDoorRedirectProtocolHTTPSOnly),
								RedirectType:      to.Ptr(armfrontdoor.FrontDoorRedirectTypeMoved),
							},
						},
						MatchConditions: []*armfrontdoor.RulesEngineMatchCondition{
							{
								RulesEngineMatchValue: []*string{
									to.Ptr("CH")},
								RulesEngineMatchVariable: to.Ptr(armfrontdoor.RulesEngineMatchVariableRemoteAddr),
								RulesEngineOperator:      to.Ptr(armfrontdoor.RulesEngineOperatorGeoMatch),
							}},
						MatchProcessingBehavior: to.Ptr(armfrontdoor.MatchProcessingBehaviorStop),
						Priority:                to.Ptr[int32](1),
					},
					{
						Name: to.Ptr("Rule2"),
						Action: &armfrontdoor.RulesEngineAction{
							ResponseHeaderActions: []*armfrontdoor.HeaderAction{
								{
									HeaderActionType: to.Ptr(armfrontdoor.HeaderActionTypeOverwrite),
									HeaderName:       to.Ptr("Cache-Control"),
									Value:            to.Ptr("public, max-age=31536000"),
								}},
						},
						MatchConditions: []*armfrontdoor.RulesEngineMatchCondition{
							{
								RulesEngineMatchValue: []*string{
									to.Ptr("jpg")},
								RulesEngineMatchVariable: to.Ptr(armfrontdoor.RulesEngineMatchVariableRequestFilenameExtension),
								RulesEngineOperator:      to.Ptr(armfrontdoor.RulesEngineOperatorEqual),
								Transforms: []*armfrontdoor.Transform{
									to.Ptr(armfrontdoor.TransformLowercase)},
							}},
						Priority: to.Ptr[int32](2),
					},
					{
						Name: to.Ptr("Rule3"),
						Action: &armfrontdoor.RulesEngineAction{
							RouteConfigurationOverride: &armfrontdoor.ForwardingConfiguration{
								ODataType: to.Ptr("#Microsoft.Azure.FrontDoor.Models.FrontdoorForwardingConfiguration"),
								BackendPool: &armfrontdoor.SubResource{
									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/backendPools/backendPool1"),
								},
								CacheConfiguration: &armfrontdoor.CacheConfiguration{
									CacheDuration:                to.Ptr("P1DT12H20M30S"),
									DynamicCompression:           to.Ptr(armfrontdoor.DynamicCompressionEnabledDisabled),
									QueryParameterStripDirective: to.Ptr(armfrontdoor.FrontDoorQueryStripOnly),
									QueryParameters:              to.Ptr("a=b,p=q"),
								},
								ForwardingProtocol: to.Ptr(armfrontdoor.FrontDoorForwardingProtocolHTTPSOnly),
							},
						},
						MatchConditions: []*armfrontdoor.RulesEngineMatchCondition{
							{
								NegateCondition: to.Ptr(false),
								RulesEngineMatchValue: []*string{
									to.Ptr("allowoverride")},
								RulesEngineMatchVariable: to.Ptr(armfrontdoor.RulesEngineMatchVariableRequestHeader),
								RulesEngineOperator:      to.Ptr(armfrontdoor.RulesEngineOperatorEqual),
								Selector:                 to.Ptr("Rules-Engine-Route-Forward"),
								Transforms: []*armfrontdoor.Transform{
									to.Ptr(armfrontdoor.TransformLowercase)},
							}},
						Priority: to.Ptr[int32](3),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
