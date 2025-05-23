using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.ManagedNetworkFabric;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkTapRules_Update_MaximumSet_Gen.json
// this example is just showing the usage of "NetworkTapRules_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkTapRuleResource created on azure
// for more information of creating NetworkTapRuleResource, please refer to the document of NetworkTapRuleResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string resourceGroupName = "example-rg";
string networkTapRuleName = "example-tapRule";
ResourceIdentifier networkTapRuleResourceId = NetworkTapRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkTapRuleName);
NetworkTapRuleResource networkTapRule = client.GetNetworkTapRuleResource(networkTapRuleResourceId);

// invoke the operation
NetworkTapRulePatch patch = new NetworkTapRulePatch
{
    Annotation = "annotation",
    ConfigurationType = NetworkFabricConfigurationType.File,
    TapRulesUri = new Uri("https://microsoft.com/amdsdx"),
    MatchConfigurations = {new NetworkTapRuleMatchConfiguration
    {
    MatchConfigurationName = "config1",
    SequenceNumber = 10L,
    IPAddressType = NetworkFabricIPAddressType.IPv4,
    MatchConditions = {new NetworkTapRuleMatchCondition
    {
    EncapsulationType = NetworkTapEncapsulationType.None,
    PortCondition = new NetworkFabricPortCondition(Layer4Protocol.Tcp)
    {
    PortType = NetworkFabricPortType.SourcePort,
    Ports = {"100"},
    PortGroupNames = {"example-portGroup1"},
    },
    ProtocolTypes = {"TCP"},
    VlanMatchCondition = new VlanMatchCondition
    {
    Vlans = {"10"},
    InnerVlans = {"11-20"},
    VlanGroupNames = {"exmaple-vlanGroup"},
    },
    IPCondition = new IPMatchCondition
    {
    SourceDestinationType = SourceDestinationType.SourceIP,
    PrefixType = IPMatchConditionPrefixType.Prefix,
    IPPrefixValues = {"10.10.10.10/20"},
    IPGroupNames = {"example-ipGroup"},
    },
    }},
    Actions = {new NetworkTapRuleAction
    {
    TapRuleActionType = TapRuleActionType.Goto,
    Truncate = "100",
    IsTimestampEnabled = NetworkFabricBooleanValue.True,
    DestinationId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/neighborGroups/example-neighborGroup"),
    MatchConfigurationName = "match1",
    }},
    }},
    DynamicMatchConfigurations = {new CommonDynamicMatchConfiguration
    {
    IPGroups = {new MatchConfigurationIPGroupProperties
    {
    Name = "example-ipGroup1",
    IPAddressType = NetworkFabricIPAddressType.IPv4,
    IPPrefixes = {"10.10.10.10/30"},
    }},
    VlanGroups = {new VlanGroupProperties
    {
    Name = "exmaple-vlanGroup",
    Vlans = {"10", "100-200"},
    }},
    PortGroups = {new PortGroupProperties
    {
    Name = "example-portGroup1",
    Ports = {"100-200"},
    }},
    }},
    Tags =
    {
    ["keyID"] = "keyValue"
    },
};
ArmOperation<NetworkTapRuleResource> lro = await networkTapRule.UpdateAsync(WaitUntil.Completed, patch);
NetworkTapRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkTapRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
