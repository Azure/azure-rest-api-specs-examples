using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleAssignmentGetByEnrollmentAccount.json
// this example is just showing the usage of "BillingRoleAssignments_GetByEnrollmentAccount" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingEnrollmentAccountRoleAssignmentResource created on azure
// for more information of creating BillingEnrollmentAccountRoleAssignmentResource, please refer to the document of BillingEnrollmentAccountRoleAssignmentResource
string billingAccountName = "7898901";
string enrollmentAccountName = "225314";
string billingRoleAssignmentName = "9dfd08c2-62a3-4d47-85bd-1cdba1408402";
ResourceIdentifier billingEnrollmentAccountRoleAssignmentResourceId = BillingEnrollmentAccountRoleAssignmentResource.CreateResourceIdentifier(billingAccountName, enrollmentAccountName, billingRoleAssignmentName);
BillingEnrollmentAccountRoleAssignmentResource billingEnrollmentAccountRoleAssignment = client.GetBillingEnrollmentAccountRoleAssignmentResource(billingEnrollmentAccountRoleAssignmentResourceId);

// invoke the operation
BillingEnrollmentAccountRoleAssignmentResource result = await billingEnrollmentAccountRoleAssignment.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BillingRoleAssignmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
