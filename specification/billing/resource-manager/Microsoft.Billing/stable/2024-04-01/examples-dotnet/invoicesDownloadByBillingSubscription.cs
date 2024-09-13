using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/invoicesDownloadByBillingSubscription.json
// this example is just showing the usage of "Invoices_DownloadByBillingSubscription" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionBillingInvoiceResource created on azure
// for more information of creating SubscriptionBillingInvoiceResource, please refer to the document of SubscriptionBillingInvoiceResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string invoiceName = "E123456789";
ResourceIdentifier subscriptionBillingInvoiceResourceId = SubscriptionBillingInvoiceResource.CreateResourceIdentifier(subscriptionId, invoiceName);
SubscriptionBillingInvoiceResource subscriptionBillingInvoice = client.GetSubscriptionBillingInvoiceResource(subscriptionBillingInvoiceResourceId);

// invoke the operation
string documentName = "12345678";
ArmOperation<BillingDocumentDownloadResult> lro = await subscriptionBillingInvoice.DownloadByBillingSubscriptionAsync(WaitUntil.Completed, documentName: documentName);
BillingDocumentDownloadResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
