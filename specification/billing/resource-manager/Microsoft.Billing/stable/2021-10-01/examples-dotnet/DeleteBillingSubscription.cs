using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Billing;
using Azure.ResourceManager.Billing.Models;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2021-10-01/examples/DeleteBillingSubscription.json
// this example is just showing the usage of "BillingSubscriptions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingSubscriptionResource created on azure
// for more information of creating BillingSubscriptionResource, please refer to the document of BillingSubscriptionResource
string billingAccountName = "{billingAccountName}";
string billingSubscriptionName = "7f271b5e-a89e-45cc-8fa5-cd5c643f8b5e";
ResourceIdentifier billingSubscriptionResourceId = BillingSubscriptionResource.CreateResourceIdentifier(billingAccountName, billingSubscriptionName);
BillingSubscriptionResource billingSubscription = client.GetBillingSubscriptionResource(billingSubscriptionResourceId);

// invoke the operation
await billingSubscription.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
