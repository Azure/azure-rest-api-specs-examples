using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Subscription;
using Azure.ResourceManager.Subscription.Models;

// Generated from example definition: specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/getAlias.json
// this example is just showing the usage of "Alias_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this SubscriptionAliasResource
SubscriptionAliasCollection collection = tenantResource.GetSubscriptionAliases();

// invoke the operation
string aliasName = "aliasForNewSub";
bool result = await collection.ExistsAsync(aliasName);

Console.WriteLine($"Succeeded: {result}");
