using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Hci.Vm.Models;
using Azure.ResourceManager.Hci.Vm;

// Generated from example definition: 2025-06-01-preview/SecurityRules_Get.json
// this example is just showing the usage of "SecurityRule_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HciVmSecurityRuleResource created on azure
// for more information of creating HciVmSecurityRuleResource, please refer to the document of HciVmSecurityRuleResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "testrg";
string networkSecurityGroupName = "testnsg";
string securityRuleName = "rule1";
ResourceIdentifier hciVmSecurityRuleResourceId = HciVmSecurityRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkSecurityGroupName, securityRuleName);
HciVmSecurityRuleResource hciVmSecurityRule = client.GetHciVmSecurityRuleResource(hciVmSecurityRuleResourceId);

// invoke the operation
HciVmSecurityRuleResource result = await hciVmSecurityRule.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HciVmSecurityRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
