using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/NetApp/stable/2025-09-01/examples/VolumeQuotaRules_Create.json
// this example is just showing the usage of "VolumeQuotaRules_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppVolumeResource created on azure
// for more information of creating NetAppVolumeResource, please refer to the document of NetAppVolumeResource
string subscriptionId = "5275316f-a498-48d6-b324-2cbfdc4311b9";
string resourceGroupName = "myRG";
string accountName = "account-9957";
string poolName = "pool-5210";
string volumeName = "volume-6387";
ResourceIdentifier netAppVolumeResourceId = NetAppVolumeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, poolName, volumeName);
NetAppVolumeResource netAppVolume = client.GetNetAppVolumeResource(netAppVolumeResourceId);

// get the collection of this NetAppVolumeQuotaRuleResource
NetAppVolumeQuotaRuleCollection collection = netAppVolume.GetNetAppVolumeQuotaRules();

// invoke the operation
string volumeQuotaRuleName = "rule-0004";
NetAppVolumeQuotaRuleData data = new NetAppVolumeQuotaRuleData(new AzureLocation("westus"))
{
    QuotaSizeInKiBs = 100005L,
    QuotaType = NetAppVolumeQuotaType.IndividualUserQuota,
    QuotaTarget = "1821",
};
ArmOperation<NetAppVolumeQuotaRuleResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, volumeQuotaRuleName, data);
NetAppVolumeQuotaRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetAppVolumeQuotaRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
