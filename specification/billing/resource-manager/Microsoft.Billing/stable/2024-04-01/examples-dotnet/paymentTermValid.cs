using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/paymentTermValid.json
// this example is just showing the usage of "BillingAccounts_ValidatePaymentTerms" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingAccountResource created on azure
// for more information of creating BillingAccountResource, please refer to the document of BillingAccountResource
string billingAccountName = "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
ResourceIdentifier billingAccountResourceId = BillingAccountResource.CreateResourceIdentifier(billingAccountName);
BillingAccountResource billingAccount = client.GetBillingAccountResource(billingAccountResourceId);

// invoke the operation
IEnumerable<BillingPaymentTerm> arrayOfPaymentTerm = new BillingPaymentTerm[]
{
new BillingPaymentTerm
{
Term = "net10",
StartOn = DateTimeOffset.Parse("2023-01-05T22:39:34.2606750Z"),
EndOn = DateTimeOffset.Parse("2023-01-25T22:39:34.2606750Z"),
}
};
PaymentTermsEligibilityResult result = await billingAccount.ValidatePaymentTermsAsync(arrayOfPaymentTerm);

Console.WriteLine($"Succeeded: {result}");
