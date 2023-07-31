using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/NetworkManagerAdminRuleCollectionPut.json
// this example is just showing the usage of "AdminRuleCollections_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityAdminConfigurationResource created on azure
// for more information of creating SecurityAdminConfigurationResource, please refer to the document of SecurityAdminConfigurationResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string networkManagerName = "testNetworkManager";
string configurationName = "myTestSecurityConfig";
ResourceIdentifier securityAdminConfigurationResourceId = SecurityAdminConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkManagerName, configurationName);
SecurityAdminConfigurationResource securityAdminConfiguration = client.GetSecurityAdminConfigurationResource(securityAdminConfigurationResourceId);

// get the collection of this AdminRuleGroupResource
AdminRuleGroupCollection collection = securityAdminConfiguration.GetAdminRuleGroups();

// invoke the operation
string ruleCollectionName = "testRuleCollection";
AdminRuleGroupData data = new AdminRuleGroupData()
{
    Description = "A sample policy",
    AppliesToGroups =
    {
    new NetworkManagerSecurityGroupItem(new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/networkGroups/testGroup"))
    },
};
ArmOperation<AdminRuleGroupResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, ruleCollectionName, data);
AdminRuleGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AdminRuleGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
