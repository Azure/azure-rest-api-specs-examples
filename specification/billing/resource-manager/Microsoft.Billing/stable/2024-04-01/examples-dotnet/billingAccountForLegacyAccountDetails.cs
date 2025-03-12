using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingAccountForLegacyAccountDetails.json
// this example is just showing the usage of "BillingAccounts_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this BillingAccountResource
BillingAccountCollection collection = tenantResource.GetBillingAccounts();

// invoke the operation and iterate over the result
BillingAccountCollectionGetAllOptions options = new BillingAccountCollectionGetAllOptions();
await foreach (BillingAccountResource item in collection.GetAllAsync(options))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    BillingAccountData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
