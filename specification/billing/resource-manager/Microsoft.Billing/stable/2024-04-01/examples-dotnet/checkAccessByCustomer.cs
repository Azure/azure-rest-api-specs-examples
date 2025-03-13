using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/checkAccessByCustomer.json
// this example is just showing the usage of "BillingPermissions_CheckAccessByCustomer" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingProfileCustomerResource created on azure
// for more information of creating BillingProfileCustomerResource, please refer to the document of BillingProfileCustomerResource
string billingAccountName = "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
string billingProfileName = "xxxx-xxxx-xxx-xxx";
string customerName = "703ab484-dda2-4402-827b-a74513b61e2d";
ResourceIdentifier billingProfileCustomerResourceId = BillingProfileCustomerResource.CreateResourceIdentifier(billingAccountName, billingProfileName, customerName);
BillingProfileCustomerResource billingProfileCustomer = client.GetBillingProfileCustomerResource(billingProfileCustomerResourceId);

// invoke the operation and iterate over the result
BillingCheckAccessContent content = new BillingCheckAccessContent
{
    Actions = { "Microsoft.Billing/billingAccounts/read", "Microsoft.Subscription/subscriptions/write" },
};
await foreach (BillingCheckAccessResult item in billingProfileCustomer.CheckAccessBillingPermissionsAsync(content))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
