using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NetworkManagerSecurityUserRulePut.json
// this example is just showing the usage of "SecurityUserRules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkManagerSecurityUserRulesResource created on azure
// for more information of creating NetworkManagerSecurityUserRulesResource, please refer to the document of NetworkManagerSecurityUserRulesResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string networkManagerName = "testNetworkManager";
string configurationName = "myTestSecurityConfig";
string ruleCollectionName = "testRuleCollection";
ResourceIdentifier networkManagerSecurityUserRulesResourceId = NetworkManagerSecurityUserRulesResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkManagerName, configurationName, ruleCollectionName);
NetworkManagerSecurityUserRulesResource networkManagerSecurityUserRules = client.GetNetworkManagerSecurityUserRulesResource(networkManagerSecurityUserRulesResourceId);

// get the collection of this NetworkManagerSecurityUserRuleResource
NetworkManagerSecurityUserRuleCollection collection = networkManagerSecurityUserRules.GetNetworkManagerSecurityUserRules();

// invoke the operation
string ruleName = "SampleUserRule";
NetworkManagerSecurityUserRuleData data = new NetworkManagerSecurityUserRuleData
{
    Description = "Sample User Rule",
    Protocol = SecurityConfigurationRuleProtocol.Tcp,
    Sources = {new AddressPrefixItem
    {
    AddressPrefix = "*",
    AddressPrefixType = AddressPrefixType.IPPrefix,
    }},
    Destinations = {new AddressPrefixItem
    {
    AddressPrefix = "*",
    AddressPrefixType = AddressPrefixType.IPPrefix,
    }},
    SourcePortRanges = { "0-65535" },
    DestinationPortRanges = { "22" },
    Direction = SecurityConfigurationRuleDirection.Inbound,
};
ArmOperation<NetworkManagerSecurityUserRuleResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, ruleName, data);
NetworkManagerSecurityUserRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkManagerSecurityUserRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
