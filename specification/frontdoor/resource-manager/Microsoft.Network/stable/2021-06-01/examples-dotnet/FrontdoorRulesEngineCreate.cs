using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.FrontDoor.Models;
using Azure.ResourceManager.FrontDoor;

// Generated from example definition: specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorRulesEngineCreate.json
// this example is just showing the usage of "RulesEngines_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FrontDoorRulesEngineResource created on azure
// for more information of creating FrontDoorRulesEngineResource, please refer to the document of FrontDoorRulesEngineResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string frontDoorName = "frontDoor1";
string rulesEngineName = "rulesEngine1";
ResourceIdentifier frontDoorRulesEngineResourceId = FrontDoorRulesEngineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, frontDoorName, rulesEngineName);
FrontDoorRulesEngineResource frontDoorRulesEngine = client.GetFrontDoorRulesEngineResource(frontDoorRulesEngineResourceId);

// invoke the operation
FrontDoorRulesEngineData data = new FrontDoorRulesEngineData
{
    Rules = {new RulesEngineRule("Rule1", 1, new RulesEngineAction
    {
    RouteConfigurationOverride = new RedirectConfiguration
    {
    RedirectType = FrontDoorRedirectType.Moved,
    RedirectProtocol = FrontDoorRedirectProtocol.HttpsOnly,
    CustomHost = "www.bing.com",
    CustomPath = "/api",
    CustomFragment = "fragment",
    CustomQueryString = "a=b",
    },
    })
    {
    MatchConditions = {new RulesEngineMatchCondition(RulesEngineMatchVariable.RemoteAddr, RulesEngineOperator.GeoMatch, new string[]{"CH"})},
    MatchProcessingBehavior = MatchProcessingBehavior.Stop,
    }, new RulesEngineRule("Rule2", 2, new RulesEngineAction
    {
    ResponseHeaderActions = {new RulesEngineHeaderAction(RulesEngineHeaderActionType.Overwrite, "Cache-Control")
    {
    Value = "public, max-age=31536000",
    }},
    })
    {
    MatchConditions = {new RulesEngineMatchCondition(RulesEngineMatchVariable.RequestFilenameExtension, RulesEngineOperator.Equal, new string[]{"jpg"})
    {
    Transforms = {RulesEngineMatchTransform.Lowercase},
    }},
    }, new RulesEngineRule("Rule3", 3, new RulesEngineAction
    {
    RouteConfigurationOverride = new ForwardingConfiguration
    {
    CustomForwardingPath = null,
    ForwardingProtocol = FrontDoorForwardingProtocol.HttpsOnly,
    CacheConfiguration = new FrontDoorCacheConfiguration
    {
    QueryParameterStripDirective = FrontDoorQuery.StripOnly,
    QueryParameters = "a=b,p=q",
    DynamicCompression = DynamicCompressionEnabled.Disabled,
    CacheDuration = XmlConvert.ToTimeSpan("P1DT12H20M30S"),
    },
    BackendPoolId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/backendPools/backendPool1"),
    },
    })
    {
    MatchConditions = {new RulesEngineMatchCondition(RulesEngineMatchVariable.RequestHeader, RulesEngineOperator.Equal, new string[]{"allowoverride"})
    {
    Selector = "Rules-Engine-Route-Forward",
    IsNegateCondition = false,
    Transforms = {RulesEngineMatchTransform.Lowercase},
    }},
    }},
};
ArmOperation<FrontDoorRulesEngineResource> lro = await frontDoorRulesEngine.UpdateAsync(WaitUntil.Completed, data);
FrontDoorRulesEngineResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FrontDoorRulesEngineData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
