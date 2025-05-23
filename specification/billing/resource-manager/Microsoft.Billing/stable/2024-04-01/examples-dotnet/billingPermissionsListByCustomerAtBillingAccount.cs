using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingPermissionsListByCustomerAtBillingAccount.json
// this example is just showing the usage of "BillingPermissions_ListByCustomerAtBillingAccount" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingCustomerResource created on azure
// for more information of creating BillingCustomerResource, please refer to the document of BillingCustomerResource
string billingAccountName = "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
string customerName = "11111111-1111-1111-1111-111111111111";
ResourceIdentifier billingCustomerResourceId = BillingCustomerResource.CreateResourceIdentifier(billingAccountName, customerName);
BillingCustomerResource billingCustomer = client.GetBillingCustomerResource(billingCustomerResourceId);

// invoke the operation and iterate over the result
await foreach (BillingPermission item in billingCustomer.GetBillingPermissionsByCustomerAtBillingAccountAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
