using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/policiesPutByBillingProfile.json
// this example is just showing the usage of "Policies_CreateOrUpdateByBillingProfile" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingProfilePolicyResource created on azure
// for more information of creating BillingProfilePolicyResource, please refer to the document of BillingProfilePolicyResource
string billingAccountName = "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
string billingProfileName = "xxxx-xxxx-xxx-xxx";
ResourceIdentifier billingProfilePolicyResourceId = BillingProfilePolicyResource.CreateResourceIdentifier(billingAccountName, billingProfileName);
BillingProfilePolicyResource billingProfilePolicy = client.GetBillingProfilePolicyResource(billingProfilePolicyResourceId);

// invoke the operation
BillingProfilePolicyData data = new BillingProfilePolicyData
{
    Properties = new BillingProfilePolicyProperties
    {
        InvoiceSectionLabelManagement = InvoiceSectionLabelManagementPolicy.Allowed,
        MarketplacePurchases = MarketplacePurchasesPolicy.AllAllowed,
        ReservationPurchases = ReservationPurchasesPolicy.Allowed,
        SavingsPlanPurchases = SavingsPlanPurchasesPolicy.Allowed,
    },
};
ArmOperation<BillingProfilePolicyResource> lro = await billingProfilePolicy.CreateOrUpdateAsync(WaitUntil.Completed, data);
BillingProfilePolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BillingProfilePolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
