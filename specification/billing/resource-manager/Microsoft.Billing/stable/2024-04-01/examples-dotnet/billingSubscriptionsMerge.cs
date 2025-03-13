using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingSubscriptionsMerge.json
// this example is just showing the usage of "BillingSubscriptions_Merge" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingSubscriptionResource created on azure
// for more information of creating BillingSubscriptionResource, please refer to the document of BillingSubscriptionResource
string billingAccountName = "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
string billingSubscriptionName = "11111111-1111-1111-1111-111111111111";
ResourceIdentifier billingSubscriptionResourceId = BillingSubscriptionResource.CreateResourceIdentifier(billingAccountName, billingSubscriptionName);
BillingSubscriptionResource billingSubscription = client.GetBillingSubscriptionResource(billingSubscriptionResourceId);

// invoke the operation
BillingSubscriptionMergeContent content = new BillingSubscriptionMergeContent
{
    TargetBillingSubscriptionName = "22222222-2222-2222-2222-222222222222",
    Quantity = 1,
};
ArmOperation<BillingSubscriptionResource> lro = await billingSubscription.MergeAsync(WaitUntil.Completed, content);
BillingSubscriptionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BillingSubscriptionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
