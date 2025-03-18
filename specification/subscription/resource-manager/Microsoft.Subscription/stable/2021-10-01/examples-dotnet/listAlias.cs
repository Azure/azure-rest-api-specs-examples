using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Subscription.Models;
using Azure.ResourceManager.Subscription;

// Generated from example definition: specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/listAlias.json
// this example is just showing the usage of "Alias_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this SubscriptionAliasResource
SubscriptionAliasCollection collection = tenantResource.GetSubscriptionAliases();

// invoke the operation and iterate over the result
await foreach (SubscriptionAliasResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SubscriptionAliasData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
