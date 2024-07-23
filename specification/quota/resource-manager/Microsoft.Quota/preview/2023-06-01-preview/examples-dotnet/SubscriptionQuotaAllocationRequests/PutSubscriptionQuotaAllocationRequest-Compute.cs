using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagementGroups;
using Azure.ResourceManager.Quota.Models;
using Azure.ResourceManager.Quota;

// Generated from example definition: specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/SubscriptionQuotaAllocationRequests/PutSubscriptionQuotaAllocationRequest-Compute.json
// this example is just showing the usage of "GroupQuotaSubscriptionAllocationRequest_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagementGroupResource created on azure
// for more information of creating ManagementGroupResource, please refer to the document of ManagementGroupResource
string managementGroupId = "E7EC67B3-7657-4966-BFFC-41EFD36BAA09";
ResourceIdentifier managementGroupResourceId = ManagementGroupResource.CreateResourceIdentifier(managementGroupId);
ManagementGroupResource managementGroupResource = client.GetManagementGroupResource(managementGroupResourceId);

// invoke the operation
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string groupQuotaName = "groupquota1";
string resourceProviderName = "Microsoft.Compute";
string resourceName = "standardav2family";
QuotaAllocationRequestStatusData data = new QuotaAllocationRequestStatusData()
{
    RequestedResource = new QuotaAllocationRequestBase()
    {
        Limit = 10,
        Region = "westus",
    },
};
ArmOperation<QuotaAllocationRequestStatusResource> lro = await managementGroupResource.CreateOrUpdateGroupQuotaSubscriptionAllocationRequestAsync(WaitUntil.Completed, subscriptionId, groupQuotaName, resourceProviderName, resourceName, data);
QuotaAllocationRequestStatusResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
QuotaAllocationRequestStatusData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
