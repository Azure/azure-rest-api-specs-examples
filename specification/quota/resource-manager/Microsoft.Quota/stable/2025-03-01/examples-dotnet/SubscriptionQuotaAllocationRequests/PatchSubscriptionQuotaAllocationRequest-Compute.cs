using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Quota.Models;
using Azure.ResourceManager.Quota;

// Generated from example definition: specification/quota/resource-manager/Microsoft.Quota/stable/2025-03-01/examples/SubscriptionQuotaAllocationRequests/PatchSubscriptionQuotaAllocationRequest-Compute.json
// this example is just showing the usage of "GroupQuotaSubscriptionAllocationRequest_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionQuotaAllocationsListResource created on azure
// for more information of creating SubscriptionQuotaAllocationsListResource, please refer to the document of SubscriptionQuotaAllocationsListResource
string managementGroupId = "E7EC67B3-7657-4966-BFFC-41EFD36BAA09";
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string groupQuotaName = "groupquota1";
string resourceProviderName = "Microsoft.Compute";
AzureLocation location = new AzureLocation("westus");
ResourceIdentifier subscriptionQuotaAllocationsListResourceId = SubscriptionQuotaAllocationsListResource.CreateResourceIdentifier(managementGroupId, subscriptionId, groupQuotaName, resourceProviderName, location);
SubscriptionQuotaAllocationsListResource subscriptionQuotaAllocationsList = client.GetSubscriptionQuotaAllocationsListResource(subscriptionQuotaAllocationsListResourceId);

// invoke the operation
SubscriptionQuotaAllocationsListData data = new SubscriptionQuotaAllocationsListData
{
    Properties = new SubscriptionQuotaAllocationsListProperties
    {
        Value = {new SubscriptionQuotaAllocations
        {
        Properties = new SubscriptionQuotaAllocationsProperties
        {
        ResourceName = "standardddv4family",
        Limit = 110L,
        },
        }, new SubscriptionQuotaAllocations
        {
        Properties = new SubscriptionQuotaAllocationsProperties
        {
        ResourceName = "standardav2family",
        Limit = 110L,
        },
        }},
    },
};
ArmOperation<SubscriptionQuotaAllocationsListResource> lro = await subscriptionQuotaAllocationsList.UpdateAsync(WaitUntil.Completed, data);
SubscriptionQuotaAllocationsListResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SubscriptionQuotaAllocationsListData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
