using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NetworkSecurityGroupRuleCreate.json
// this example is just showing the usage of "SecurityRules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityRuleResource created on azure
// for more information of creating SecurityRuleResource, please refer to the document of SecurityRuleResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string networkSecurityGroupName = "testnsg";
string securityRuleName = "rule1";
ResourceIdentifier securityRuleResourceId = SecurityRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkSecurityGroupName, securityRuleName);
SecurityRuleResource securityRule = client.GetSecurityRuleResource(securityRuleResourceId);

// invoke the operation
SecurityRuleData data = new SecurityRuleData
{
    Protocol = SecurityRuleProtocol.Asterisk,
    SourcePortRange = "*",
    DestinationPortRange = "8080",
    SourceAddressPrefix = "10.0.0.0/8",
    DestinationAddressPrefix = "11.0.0.0/8",
    Access = SecurityRuleAccess.Deny,
    Priority = 100,
    Direction = SecurityRuleDirection.Outbound,
};
ArmOperation<SecurityRuleResource> lro = await securityRule.UpdateAsync(WaitUntil.Completed, data);
SecurityRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
