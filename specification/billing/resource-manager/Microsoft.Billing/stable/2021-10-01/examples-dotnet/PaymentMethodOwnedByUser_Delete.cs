using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2021-10-01/examples/PaymentMethodOwnedByUser_Delete.json
// this example is just showing the usage of "PaymentMethods_DeleteByUser" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingPaymentMethodResource created on azure
// for more information of creating BillingPaymentMethodResource, please refer to the document of BillingPaymentMethodResource
string paymentMethodName = "ABCDABCDABC0";
ResourceIdentifier billingPaymentMethodResourceId = BillingPaymentMethodResource.CreateResourceIdentifier(paymentMethodName);
BillingPaymentMethodResource billingPaymentMethod = client.GetBillingPaymentMethodResource(billingPaymentMethodResourceId);

// invoke the operation
await billingPaymentMethod.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
