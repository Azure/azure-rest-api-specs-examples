Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ffrontdoor%2Farmfrontdoor%2Fv0.1.0/sdk/resourcemanager/frontdoor/armfrontdoor/README.md) on how to add the SDK to your project and authenticate.

```go
package armfrontdoor_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorRulesEngineCreate.json
func ExampleRulesEnginesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armfrontdoor.NewRulesEnginesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<front-door-name>",
		"<rules-engine-name>",
		armfrontdoor.RulesEngine{
			Properties: &armfrontdoor.RulesEngineProperties{
				RulesEngineUpdateParameters: armfrontdoor.RulesEngineUpdateParameters{
					Rules: []*armfrontdoor.RulesEngineRule{
						{
							Name: to.StringPtr("<name>"),
							Action: &armfrontdoor.RulesEngineAction{
								RouteConfigurationOverride: &armfrontdoor.RedirectConfiguration{
									RouteConfiguration: armfrontdoor.RouteConfiguration{
										ODataType: to.StringPtr("<odata-type>"),
									},
									CustomFragment:    to.StringPtr("<custom-fragment>"),
									CustomHost:        to.StringPtr("<custom-host>"),
									CustomPath:        to.StringPtr("<custom-path>"),
									CustomQueryString: to.StringPtr("<custom-query-string>"),
									RedirectProtocol:  armfrontdoor.FrontDoorRedirectProtocolHTTPSOnly.ToPtr(),
									RedirectType:      armfrontdoor.FrontDoorRedirectTypeMoved.ToPtr(),
								},
							},
							MatchConditions: []*armfrontdoor.RulesEngineMatchCondition{
								{
									RulesEngineMatchValue: []*string{
										to.StringPtr("CH")},
									RulesEngineMatchVariable: armfrontdoor.RulesEngineMatchVariableRemoteAddr.ToPtr(),
									RulesEngineOperator:      armfrontdoor.RulesEngineOperatorGeoMatch.ToPtr(),
								}},
							MatchProcessingBehavior: armfrontdoor.MatchProcessingBehaviorStop.ToPtr(),
							Priority:                to.Int32Ptr(1),
						},
						{
							Name: to.StringPtr("<name>"),
							Action: &armfrontdoor.RulesEngineAction{
								ResponseHeaderActions: []*armfrontdoor.HeaderAction{
									{
										HeaderActionType: armfrontdoor.HeaderActionTypeOverwrite.ToPtr(),
										HeaderName:       to.StringPtr("<header-name>"),
										Value:            to.StringPtr("<value>"),
									}},
							},
							MatchConditions: []*armfrontdoor.RulesEngineMatchCondition{
								{
									RulesEngineMatchValue: []*string{
										to.StringPtr("jpg")},
									RulesEngineMatchVariable: armfrontdoor.RulesEngineMatchVariableRequestFilenameExtension.ToPtr(),
									RulesEngineOperator:      armfrontdoor.RulesEngineOperatorEqual.ToPtr(),
									Transforms: []*armfrontdoor.Transform{
										armfrontdoor.TransformLowercase.ToPtr()},
								}},
							Priority: to.Int32Ptr(2),
						},
						{
							Name: to.StringPtr("<name>"),
							Action: &armfrontdoor.RulesEngineAction{
								RouteConfigurationOverride: &armfrontdoor.ForwardingConfiguration{
									RouteConfiguration: armfrontdoor.RouteConfiguration{
										ODataType: to.StringPtr("<odata-type>"),
									},
									BackendPool: &armfrontdoor.SubResource{
										ID: to.StringPtr("<id>"),
									},
									CacheConfiguration: &armfrontdoor.CacheConfiguration{
										CacheDuration:                to.StringPtr("<cache-duration>"),
										DynamicCompression:           armfrontdoor.DynamicCompressionEnabledDisabled.ToPtr(),
										QueryParameterStripDirective: armfrontdoor.FrontDoorQueryStripOnly.ToPtr(),
										QueryParameters:              to.StringPtr("<query-parameters>"),
									},
									CustomForwardingPath: to.StringPtr("<custom-forwarding-path>"),
									ForwardingProtocol:   armfrontdoor.FrontDoorForwardingProtocolHTTPSOnly.ToPtr(),
								},
							},
							MatchConditions: []*armfrontdoor.RulesEngineMatchCondition{
								{
									NegateCondition: to.BoolPtr(false),
									RulesEngineMatchValue: []*string{
										to.StringPtr("allowoverride")},
									RulesEngineMatchVariable: armfrontdoor.RulesEngineMatchVariableRequestHeader.ToPtr(),
									RulesEngineOperator:      armfrontdoor.RulesEngineOperatorEqual.ToPtr(),
									Selector:                 to.StringPtr("<selector>"),
									Transforms: []*armfrontdoor.Transform{
										armfrontdoor.TransformLowercase.ToPtr()},
								}},
							Priority: to.Int32Ptr(3),
						}},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("RulesEngine.ID: %s\n", *res.ID)
}
```
