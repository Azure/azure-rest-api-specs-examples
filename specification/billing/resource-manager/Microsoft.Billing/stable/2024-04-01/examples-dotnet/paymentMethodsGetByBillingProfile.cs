using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/paymentMethodsGetByBillingProfile.json
// this example is just showing the usage of "PaymentMethods_GetByBillingProfile" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingProfileResource created on azure
// for more information of creating BillingProfileResource, please refer to the document of BillingProfileResource
string billingAccountName = "00000000-0000-0000-0000-000000000032:00000000-0000-0000-0000-000000000099_2019-05-31";
string billingProfileName = "ABC1-A1CD-AB1-BP1";
ResourceIdentifier billingProfileResourceId = BillingProfileResource.CreateResourceIdentifier(billingAccountName, billingProfileName);
BillingProfileResource billingProfile = client.GetBillingProfileResource(billingProfileResourceId);

// get the collection of this BillingPaymentMethodLinkResource
BillingPaymentMethodLinkCollection collection = billingProfile.GetBillingPaymentMethodLinks();

// invoke the operation
string paymentMethodName = "ABCDABCDABC0";
NullableResponse<BillingPaymentMethodLinkResource> response = await collection.GetIfExistsAsync(paymentMethodName);
BillingPaymentMethodLinkResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    BillingPaymentMethodLinkData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
