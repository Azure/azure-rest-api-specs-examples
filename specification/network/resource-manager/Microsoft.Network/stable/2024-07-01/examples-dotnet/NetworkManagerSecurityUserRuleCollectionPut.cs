using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NetworkManagerSecurityUserRuleCollectionPut.json
// this example is just showing the usage of "SecurityUserRuleCollections_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
NetworkManagerSecurityUserRulesData data = new NetworkManagerSecurityUserRulesData
{
    Description = "A sample policy",
    AppliesToGroups = { new SecurityUserGroupItem("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/testGroup") },
};
ArmOperation<NetworkManagerSecurityUserRulesResource> lro = await networkManagerSecurityUserRules.UpdateAsync(WaitUntil.Completed, data);
NetworkManagerSecurityUserRulesResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkManagerSecurityUserRulesData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
