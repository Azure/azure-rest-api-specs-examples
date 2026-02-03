using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkManagerAdminRulePut_NetworkGroupSource.json
// this example is just showing the usage of "AdminRules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BaseAdminRuleResource created on azure
// for more information of creating BaseAdminRuleResource, please refer to the document of BaseAdminRuleResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string networkManagerName = "testNetworkManager";
string configurationName = "myTestSecurityConfig";
string ruleCollectionName = "testRuleCollection";
string ruleName = "SampleAdminRule";
ResourceIdentifier baseAdminRuleResourceId = BaseAdminRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkManagerName, configurationName, ruleCollectionName, ruleName);
BaseAdminRuleResource baseAdminRule = client.GetBaseAdminRuleResource(baseAdminRuleResourceId);

// invoke the operation
BaseAdminRuleData data = new NetworkAdminRule
{
    Description = "This is Sample Admin Rule",
    Protocol = SecurityConfigurationRuleProtocol.Tcp,
    Sources = {new AddressPrefixItem
    {
    AddressPrefix = "Internet",
    AddressPrefixType = AddressPrefixType.ServiceTag,
    }},
    Destinations = {new AddressPrefixItem
    {
    AddressPrefix = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/ng1",
    AddressPrefixType = AddressPrefixType.NetworkGroup,
    }},
    SourcePortRanges = { "0-65535" },
    DestinationPortRanges = { "22" },
    Access = SecurityConfigurationRuleAccess.Deny,
    Priority = 1,
    Direction = SecurityConfigurationRuleDirection.Inbound,
};
ArmOperation<BaseAdminRuleResource> lro = await baseAdminRule.UpdateAsync(WaitUntil.Completed, data);
BaseAdminRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BaseAdminRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
