using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRequestsCreateOrUpdate.json
// this example is just showing the usage of "BillingRequests_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this BillingRequestResource
BillingRequestCollection collection = tenantResource.GetBillingRequests();

// invoke the operation
string billingRequestName = "00000000-0000-0000-0000-000000000000";
BillingRequestData data = new BillingRequestData
{
    Properties = new BillingRequestProperties
    {
        AdditionalInformation =
        {
        ["RoleId"] = "40000000-aaaa-bbbb-cccc-200000000006"
        },
        DecisionReason = "New team member",
        RequestScope = "/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx",
        Status = BillingRequestStatus.Pending,
        RequestType = BillingRequestType.RoleAssignment,
    },
};
ArmOperation<BillingRequestResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, billingRequestName, data);
BillingRequestResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BillingRequestData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
