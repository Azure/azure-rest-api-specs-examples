using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagementGroups;
using Azure.ResourceManager.Quota;

// Generated from example definition: specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/SubscriptionQuotaAllocation/SubscriptionQuotaAllocation_List-Compute.json
// this example is just showing the usage of "GroupQuotaSubscriptionAllocation_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagementGroupResource created on azure
// for more information of creating ManagementGroupResource, please refer to the document of ManagementGroupResource
string managementGroupId = "E7EC67B3-7657-4966-BFFC-41EFD36BAA09";
ResourceIdentifier managementGroupResourceId = ManagementGroupResource.CreateResourceIdentifier(managementGroupId);
ManagementGroupResource managementGroupResource = client.GetManagementGroupResource(managementGroupResourceId);

// get the collection of this SubscriptionQuotaAllocationResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string groupQuotaName = "groupquota1";
SubscriptionQuotaAllocationCollection collection = managementGroupResource.GetSubscriptionQuotaAllocations(subscriptionId, groupQuotaName);

// invoke the operation and iterate over the result
string filter = "provider eq Microsoft.Compute & location eq westus";
await foreach (SubscriptionQuotaAllocationResource item in collection.GetAllAsync(filter))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SubscriptionQuotaAllocationData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
