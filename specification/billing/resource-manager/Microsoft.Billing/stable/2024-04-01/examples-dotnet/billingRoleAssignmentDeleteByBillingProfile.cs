using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleAssignmentDeleteByBillingProfile.json
// this example is just showing the usage of "BillingRoleAssignments_DeleteByBillingProfile" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingProfileRoleAssignmentResource created on azure
// for more information of creating BillingProfileRoleAssignmentResource, please refer to the document of BillingProfileRoleAssignmentResource
string billingAccountName = "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30";
string billingProfileName = "BKM6-54VH-BG7-PGB";
string billingRoleAssignmentName = "10000000-aaaa-bbbb-cccc-100000000000_6fd330f6-7d26-4aff-b9cf-7bd699f965b9";
ResourceIdentifier billingProfileRoleAssignmentResourceId = BillingProfileRoleAssignmentResource.CreateResourceIdentifier(billingAccountName, billingProfileName, billingRoleAssignmentName);
BillingProfileRoleAssignmentResource billingProfileRoleAssignment = client.GetBillingProfileRoleAssignmentResource(billingProfileRoleAssignmentResourceId);

// invoke the operation
await billingProfileRoleAssignment.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
