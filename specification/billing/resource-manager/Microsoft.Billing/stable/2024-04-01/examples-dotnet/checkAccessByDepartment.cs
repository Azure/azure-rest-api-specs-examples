using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/checkAccessByDepartment.json
// this example is just showing the usage of "BillingPermissions_CheckAccessByDepartment" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingDepartmentResource created on azure
// for more information of creating BillingDepartmentResource, please refer to the document of BillingDepartmentResource
string billingAccountName = "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
string departmentName = "12345";
ResourceIdentifier billingDepartmentResourceId = BillingDepartmentResource.CreateResourceIdentifier(billingAccountName, departmentName);
BillingDepartmentResource billingDepartment = client.GetBillingDepartmentResource(billingDepartmentResourceId);

// invoke the operation and iterate over the result
BillingCheckAccessContent content = new BillingCheckAccessContent()
{
    Actions =
    {
    "Microsoft.Billing/billingAccounts/read","Microsoft.Subscription/subscriptions/write"
    },
};
await foreach (BillingCheckAccessResult item in billingDepartment.CheckAccessBillingPermissionsAsync(content))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
