using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Billing;
using Azure.ResourceManager.Billing.Models;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2021-10-01/examples/BillingSubscription.json
// this example is just showing the usage of "BillingSubscriptions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingSubscriptionResource created on azure
// for more information of creating BillingSubscriptionResource, please refer to the document of BillingSubscriptionResource
string billingAccountName = "{billingAccountName}";
string billingSubscriptionName = "418b0e9c-5dc3-4260-918f-30b90619fe07";
ResourceIdentifier billingSubscriptionResourceId = BillingSubscriptionResource.CreateResourceIdentifier(billingAccountName, billingSubscriptionName);
BillingSubscriptionResource billingSubscription = client.GetBillingSubscriptionResource(billingSubscriptionResourceId);

// invoke the operation
BillingSubscriptionResource result = await billingSubscription.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BillingSubscriptionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
