using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingSubscriptionsListByCustomerAtBillingAccount.json
// this example is just showing the usage of "BillingSubscriptions_ListByCustomerAtBillingAccount" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingCustomerResource created on azure
// for more information of creating BillingCustomerResource, please refer to the document of BillingCustomerResource
string billingAccountName = "a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31";
string customerName = "Q7GV-UUVA-PJA-TGB";
ResourceIdentifier billingCustomerResourceId = BillingCustomerResource.CreateResourceIdentifier(billingAccountName, customerName);
BillingCustomerResource billingCustomer = client.GetBillingCustomerResource(billingCustomerResourceId);

// invoke the operation and iterate over the result
BillingCustomerResourceGetBillingSubscriptionsByCustomerAtBillingAccountOptions options = new BillingCustomerResourceGetBillingSubscriptionsByCustomerAtBillingAccountOptions();
await foreach (BillingSubscriptionData item in billingCustomer.GetBillingSubscriptionsByCustomerAtBillingAccountAsync(options))
{
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {item.Id}");
}

Console.WriteLine("Succeeded");
