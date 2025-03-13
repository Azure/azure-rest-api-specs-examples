using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/policiesPutByCustomer.json
// this example is just showing the usage of "Policies_CreateOrUpdateByCustomer" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingProfileCustomerResource created on azure
// for more information of creating BillingProfileCustomerResource, please refer to the document of BillingProfileCustomerResource
string billingAccountName = "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
string billingProfileName = "xxxx-xxxx-xxx-xxx";
string customerName = "11111111-1111-1111-1111-111111111111";
ResourceIdentifier billingProfileCustomerResourceId = BillingProfileCustomerResource.CreateResourceIdentifier(billingAccountName, billingProfileName, customerName);
BillingProfileCustomerResource billingProfileCustomer = client.GetBillingProfileCustomerResource(billingProfileCustomerResourceId);

// invoke the operation
BillingCustomerPolicyData data = new BillingCustomerPolicyData
{
    Properties = new BillingCustomerPolicyProperties(ViewChargesPolicy.Allowed),
};
ArmOperation<BillingCustomerPolicyData> lro = await billingProfileCustomer.CreateOrUpdateByCustomerPolicyAsync(WaitUntil.Completed, data);
BillingCustomerPolicyData result = lro.Value;

// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {result.Id}");
