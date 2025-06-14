using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventHubs.Models;
using Azure.ResourceManager.EventHubs;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/NameSpaces/VirtualNetworkRule/EHNetworkRuleSetCreate.json
// this example is just showing the usage of "Namespaces_CreateOrUpdateNetworkRuleSet" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventHubsNetworkRuleSetResource created on azure
// for more information of creating EventHubsNetworkRuleSetResource, please refer to the document of EventHubsNetworkRuleSetResource
string subscriptionId = "Subscription";
string resourceGroupName = "ResourceGroup";
string namespaceName = "sdk-Namespace-6019";
ResourceIdentifier eventHubsNetworkRuleSetResourceId = EventHubsNetworkRuleSetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName);
EventHubsNetworkRuleSetResource eventHubsNetworkRuleSet = client.GetEventHubsNetworkRuleSetResource(eventHubsNetworkRuleSetResourceId);

// invoke the operation
EventHubsNetworkRuleSetData data = new EventHubsNetworkRuleSetData
{
    DefaultAction = EventHubsNetworkRuleSetDefaultAction.Deny,
    VirtualNetworkRules = {new EventHubsNetworkRuleSetVirtualNetworkRules
    {
    SubnetId = new ResourceIdentifier("/subscriptions/subscriptionid/resourcegroups/resourcegroupid/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet2"),
    IgnoreMissingVnetServiceEndpoint = true,
    }, new EventHubsNetworkRuleSetVirtualNetworkRules
    {
    SubnetId = new ResourceIdentifier("/subscriptions/subscriptionid/resourcegroups/resourcegroupid/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet3"),
    IgnoreMissingVnetServiceEndpoint = false,
    }, new EventHubsNetworkRuleSetVirtualNetworkRules
    {
    SubnetId = new ResourceIdentifier("/subscriptions/subscriptionid/resourcegroups/resourcegroupid/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet6"),
    IgnoreMissingVnetServiceEndpoint = false,
    }},
    IPRules = {new EventHubsNetworkRuleSetIPRules
    {
    IPMask = "1.1.1.1",
    Action = EventHubsNetworkRuleIPAction.Allow,
    }, new EventHubsNetworkRuleSetIPRules
    {
    IPMask = "1.1.1.2",
    Action = EventHubsNetworkRuleIPAction.Allow,
    }, new EventHubsNetworkRuleSetIPRules
    {
    IPMask = "1.1.1.3",
    Action = EventHubsNetworkRuleIPAction.Allow,
    }, new EventHubsNetworkRuleSetIPRules
    {
    IPMask = "1.1.1.4",
    Action = EventHubsNetworkRuleIPAction.Allow,
    }, new EventHubsNetworkRuleSetIPRules
    {
    IPMask = "1.1.1.5",
    Action = EventHubsNetworkRuleIPAction.Allow,
    }},
};
ArmOperation<EventHubsNetworkRuleSetResource> lro = await eventHubsNetworkRuleSet.CreateOrUpdateAsync(WaitUntil.Completed, data);
EventHubsNetworkRuleSetResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EventHubsNetworkRuleSetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
