using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/invoicesDownloadDocumentsByBillingSubscription.json
// this example is just showing the usage of "Invoices_DownloadDocumentsByBillingSubscription" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
string subscriptionId = "00000000-0000-0000-0000-000000000000";
IEnumerable<BillingDocumentDownloadRequestContent> arrayOfDocumentDownloadRequest = new BillingDocumentDownloadRequestContent[]
{
new BillingDocumentDownloadRequestContent
{
DocumentName = "12345678",
InvoiceName = "E123456789",
},
new BillingDocumentDownloadRequestContent
{
DocumentName = "12345678",
InvoiceName = "E987654321",
}
};
ArmOperation<BillingDocumentDownloadResult> lro = await tenantResource.DownloadDocumentsByBillingSubscriptionInvoiceAsync(WaitUntil.Completed, subscriptionId, arrayOfDocumentDownloadRequest);
BillingDocumentDownloadResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
