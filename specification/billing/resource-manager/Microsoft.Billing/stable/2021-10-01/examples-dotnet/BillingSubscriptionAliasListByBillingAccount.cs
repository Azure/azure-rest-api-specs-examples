using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2021-10-01/examples/BillingSubscriptionAliasListByBillingAccount.json
// this example is just showing the usage of "BillingSubscriptionsAliases_ListByBillingAccount" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this BillingSubscriptionAliasResource
string billingAccountName = "{billingAccountName}";
BillingSubscriptionAliasCollection collection = tenantResource.GetBillingSubscriptionAliases(billingAccountName);

// invoke the operation and iterate over the result
await foreach (BillingSubscriptionAliasResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    BillingSubscriptionAliasData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
