using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleAssignmentGetByCustomer.json
// this example is just showing the usage of "BillingRoleAssignments_GetByCustomer" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingCustomerRoleAssignmentResource created on azure
// for more information of creating BillingCustomerRoleAssignmentResource, please refer to the document of BillingCustomerRoleAssignmentResource
string billingAccountName = "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30";
string billingProfileName = "xxxx-xxxx-xxx-xxx";
string customerName = "703ab484-dda2-4402-827b-a74513b61e2d";
string billingRoleAssignmentName = "30000000-aaaa-bbbb-cccc-100000000000_6fd330f6-7d26-4aff-b9cf-7bd699f965b9";
ResourceIdentifier billingCustomerRoleAssignmentResourceId = BillingCustomerRoleAssignmentResource.CreateResourceIdentifier(billingAccountName, billingProfileName, customerName, billingRoleAssignmentName);
BillingCustomerRoleAssignmentResource billingCustomerRoleAssignment = client.GetBillingCustomerRoleAssignmentResource(billingCustomerRoleAssignmentResourceId);

// invoke the operation
BillingCustomerRoleAssignmentResource result = await billingCustomerRoleAssignment.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BillingRoleAssignmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
