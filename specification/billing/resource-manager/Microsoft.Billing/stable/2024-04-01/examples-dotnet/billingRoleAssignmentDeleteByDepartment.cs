using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleAssignmentDeleteByDepartment.json
// this example is just showing the usage of "BillingRoleAssignments_DeleteByDepartment" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingDepartmentRoleAssignmentResource created on azure
// for more information of creating BillingDepartmentRoleAssignmentResource, please refer to the document of BillingDepartmentRoleAssignmentResource
string billingAccountName = "8608480";
string departmentName = "123456";
string billingRoleAssignmentName = "9dfd08c2-62a3-4d47-85bd-1cdba1408402";
ResourceIdentifier billingDepartmentRoleAssignmentResourceId = BillingDepartmentRoleAssignmentResource.CreateResourceIdentifier(billingAccountName, departmentName, billingRoleAssignmentName);
BillingDepartmentRoleAssignmentResource billingDepartmentRoleAssignment = client.GetBillingDepartmentRoleAssignmentResource(billingDepartmentRoleAssignmentResourceId);

// invoke the operation
await billingDepartmentRoleAssignment.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
