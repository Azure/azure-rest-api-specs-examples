using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleAssignmentCreateByBillingAccount.json
// this example is just showing the usage of "BillingRoleAssignments_CreateByBillingAccount" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingAccountResource created on azure
// for more information of creating BillingAccountResource, please refer to the document of BillingAccountResource
string billingAccountName = "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30";
ResourceIdentifier billingAccountResourceId = BillingAccountResource.CreateResourceIdentifier(billingAccountName);
BillingAccountResource billingAccount = client.GetBillingAccountResource(billingAccountResourceId);

// invoke the operation
BillingRoleAssignmentProperties billingRoleAssignmentProperties = new BillingRoleAssignmentProperties(new ResourceIdentifier("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30/billingRoleDefinitions/10000000-aaaa-bbbb-cccc-100000000000"))
{
    PrincipalId = "00000000-0000-0000-0000-000000000000",
    PrincipalTenantId = Guid.Parse("076915e7-de10-4323-bb34-a58c904068bb"),
    UserEmailAddress = "john@contoso.com",
};
ArmOperation<BillingRoleAssignmentData> lro = await billingAccount.CreateByBillingAccountBillingRoleAssignmentAsync(WaitUntil.Completed, billingRoleAssignmentProperties);
BillingRoleAssignmentData result = lro.Value;

// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {result.Id}");
