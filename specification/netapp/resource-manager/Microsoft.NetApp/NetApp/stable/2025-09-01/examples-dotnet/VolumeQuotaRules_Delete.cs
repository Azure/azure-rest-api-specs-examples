using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/NetApp/stable/2025-09-01/examples/VolumeQuotaRules_Delete.json
// this example is just showing the usage of "VolumeQuotaRules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppVolumeQuotaRuleResource created on azure
// for more information of creating NetAppVolumeQuotaRuleResource, please refer to the document of NetAppVolumeQuotaRuleResource
string subscriptionId = "5275316f-a498-48d6-b324-2cbfdc4311b9";
string resourceGroupName = "myRG";
string accountName = "account-9957";
string poolName = "pool-5210";
string volumeName = "volume-6387";
string volumeQuotaRuleName = "rule-0004";
ResourceIdentifier netAppVolumeQuotaRuleResourceId = NetAppVolumeQuotaRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, poolName, volumeName, volumeQuotaRuleName);
NetAppVolumeQuotaRuleResource netAppVolumeQuotaRule = client.GetNetAppVolumeQuotaRuleResource(netAppVolumeQuotaRuleResourceId);

// invoke the operation
await netAppVolumeQuotaRule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
