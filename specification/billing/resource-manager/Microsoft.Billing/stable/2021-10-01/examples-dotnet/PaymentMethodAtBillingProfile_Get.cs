using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2021-10-01/examples/PaymentMethodAtBillingProfile_Get.json
// this example is just showing the usage of "PaymentMethods_GetByBillingProfile" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this BillingPaymentMethodLinkResource
string billingAccountName = "00000000-0000-0000-0000-000000000032:00000000-0000-0000-0000-000000000099_2019-05-31";
string billingProfileName = "ABC1-A1CD-AB1-BP1";
BillingPaymentMethodLinkCollection collection = tenantResource.GetBillingPaymentMethodLinks(billingAccountName, billingProfileName);

// invoke the operation
string paymentMethodName = "ABCDABCDABC0";
bool result = await collection.ExistsAsync(paymentMethodName);

Console.WriteLine($"Succeeded: {result}");
