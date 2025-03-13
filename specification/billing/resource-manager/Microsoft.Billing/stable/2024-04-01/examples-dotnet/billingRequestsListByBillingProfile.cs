using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRequestsListByBillingProfile.json
// this example is just showing the usage of "BillingRequests_ListByBillingProfile" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingProfileResource created on azure
// for more information of creating BillingProfileResource, please refer to the document of BillingProfileResource
string billingAccountName = "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
string billingProfileName = "xxxx-xxxx-xxx-xxx";
ResourceIdentifier billingProfileResourceId = BillingProfileResource.CreateResourceIdentifier(billingAccountName, billingProfileName);
BillingProfileResource billingProfile = client.GetBillingProfileResource(billingProfileResourceId);

// invoke the operation and iterate over the result
BillingProfileResourceGetBillingRequestsOptions options = new BillingProfileResourceGetBillingRequestsOptions();
await foreach (BillingRequestResource item in billingProfile.GetBillingRequestsAsync(options))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    BillingRequestData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
