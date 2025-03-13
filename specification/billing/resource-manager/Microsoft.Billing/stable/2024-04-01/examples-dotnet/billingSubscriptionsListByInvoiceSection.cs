using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingSubscriptionsListByInvoiceSection.json
// this example is just showing the usage of "BillingSubscriptions_ListByInvoiceSection" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingInvoiceSectionResource created on azure
// for more information of creating BillingInvoiceSectionResource, please refer to the document of BillingInvoiceSectionResource
string billingAccountName = "a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31";
string billingProfileName = "ea36e548-1505-41db-bebc-46fff3d37998";
string invoiceSectionName = "Q7GV-UUVA-PJA-TGB";
ResourceIdentifier billingInvoiceSectionResourceId = BillingInvoiceSectionResource.CreateResourceIdentifier(billingAccountName, billingProfileName, invoiceSectionName);
BillingInvoiceSectionResource billingInvoiceSection = client.GetBillingInvoiceSectionResource(billingInvoiceSectionResourceId);

// invoke the operation and iterate over the result
BillingInvoiceSectionResourceGetBillingSubscriptionsOptions options = new BillingInvoiceSectionResourceGetBillingSubscriptionsOptions();
await foreach (BillingSubscriptionData item in billingInvoiceSection.GetBillingSubscriptionsAsync(options))
{
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {item.Id}");
}

Console.WriteLine("Succeeded");
