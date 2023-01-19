using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Billing;
using Azure.ResourceManager.Billing.Models;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2021-10-01/examples/MergeSubscription.json
// this example is just showing the usage of "BillingSubscriptions_Merge" operation, for the dependent resources, they will have to be created separately.

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
BillingSubscriptionMergeContent content = new BillingSubscriptionMergeContent()
{
    TargetBillingSubscriptionName = "9e030ae5-51ed-4e02-8e17-d51b3aa55980",
    Quantity = 4,
};
ArmOperation<BillingSubscriptionResource> lro = await billingSubscription.MergeAsync(WaitUntil.Completed, content);
BillingSubscriptionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BillingSubscriptionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
