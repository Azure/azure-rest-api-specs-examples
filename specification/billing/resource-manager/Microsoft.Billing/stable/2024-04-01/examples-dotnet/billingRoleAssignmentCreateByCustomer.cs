using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleAssignmentCreateByCustomer.json
// this example is just showing the usage of "BillingRoleAssignments_CreateByCustomer" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingProfileCustomerResource created on azure
// for more information of creating BillingProfileCustomerResource, please refer to the document of BillingProfileCustomerResource
string billingAccountName = "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30";
string billingProfileName = "BKM6-54VH-BG7-PGB";
string customerName = "703ab484-dda2-4402-827b-a74513b61e2d";
ResourceIdentifier billingProfileCustomerResourceId = BillingProfileCustomerResource.CreateResourceIdentifier(billingAccountName, billingProfileName, customerName);
BillingProfileCustomerResource billingProfileCustomer = client.GetBillingProfileCustomerResource(billingProfileCustomerResourceId);

// invoke the operation
BillingRoleAssignmentProperties billingRoleAssignmentProperties = new BillingRoleAssignmentProperties(new ResourceIdentifier("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30/billingProfileName/BKM6-54VH-BG7-PGB/customers/703ab484-dda2-4402-827b-a74513b61e2d/billingRoleDefinitions/30000000-aaaa-bbbb-cccc-100000000000"))
{
    PrincipalId = "00000000-0000-0000-0000-000000000000",
    PrincipalTenantId = Guid.Parse("076915e7-de10-4323-bb34-a58c904068bb"),
    UserEmailAddress = "john@contoso.com",
};
ArmOperation<BillingRoleAssignmentData> lro = await billingProfileCustomer.CreateByCustomerBillingRoleAssignmentAsync(WaitUntil.Completed, billingRoleAssignmentProperties);
BillingRoleAssignmentData result = lro.Value;

// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {result.Id}");
