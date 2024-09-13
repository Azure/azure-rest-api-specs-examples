using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/invoicesDownloadSummaryByBillingAccount.json
// this example is just showing the usage of "Invoices_DownloadSummaryByBillingAccount" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingInvoiceResource created on azure
// for more information of creating BillingInvoiceResource, please refer to the document of BillingInvoiceResource
string billingAccountName = "123456789";
string invoiceName = "G123456789";
ResourceIdentifier billingInvoiceResourceId = BillingInvoiceResource.CreateResourceIdentifier(billingAccountName, invoiceName);
BillingInvoiceResource billingInvoice = client.GetBillingInvoiceResource(billingInvoiceResourceId);

// invoke the operation
ArmOperation<BillingDocumentDownloadResult> lro = await billingInvoice.DownloadSummaryByBillingAccountAsync(WaitUntil.Completed);
BillingDocumentDownloadResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
