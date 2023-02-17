using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2021-10-01/examples/BillingSubscription.json
// this example is just showing the usage of "BillingSubscriptions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this BillingSubscriptionResource
string billingAccountName = "{billingAccountName}";
BillingSubscriptionCollection collection = tenantResource.GetBillingSubscriptions(billingAccountName);

// invoke the operation
string billingSubscriptionName = "418b0e9c-5dc3-4260-918f-30b90619fe07";
bool result = await collection.ExistsAsync(billingSubscriptionName);

Console.WriteLine($"Succeeded: {result}");
