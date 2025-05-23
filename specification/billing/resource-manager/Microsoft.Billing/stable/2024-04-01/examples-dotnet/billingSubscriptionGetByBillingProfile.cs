using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingSubscriptionGetByBillingProfile.json
// this example is just showing the usage of "BillingSubscriptions_GetByBillingProfile" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingProfileSubscriptionResource created on azure
// for more information of creating BillingProfileSubscriptionResource, please refer to the document of BillingProfileSubscriptionResource
string billingAccountName = "pcn.94077792";
string billingProfileName = "6478903";
string billingSubscriptionName = "6b96d3f2-9008-4a9d-912f-f87744185aa3";
ResourceIdentifier billingProfileSubscriptionResourceId = BillingProfileSubscriptionResource.CreateResourceIdentifier(billingAccountName, billingProfileName, billingSubscriptionName);
BillingProfileSubscriptionResource billingProfileSubscription = client.GetBillingProfileSubscriptionResource(billingProfileSubscriptionResourceId);

// invoke the operation
BillingProfileSubscriptionResource result = await billingProfileSubscription.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BillingSubscriptionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
