using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/enrollmentAccountByDepartment.json
// this example is just showing the usage of "EnrollmentAccounts_GetByDepartment" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingDepartmentResource created on azure
// for more information of creating BillingDepartmentResource, please refer to the document of BillingDepartmentResource
string billingAccountName = "6564892";
string departmentName = "164821";
ResourceIdentifier billingDepartmentResourceId = BillingDepartmentResource.CreateResourceIdentifier(billingAccountName, departmentName);
BillingDepartmentResource billingDepartment = client.GetBillingDepartmentResource(billingDepartmentResourceId);

// get the collection of this BillingDepartmentEnrollmentAccountResource
BillingDepartmentEnrollmentAccountCollection collection = billingDepartment.GetBillingDepartmentEnrollmentAccounts();

// invoke the operation
string enrollmentAccountName = "257698";
NullableResponse<BillingDepartmentEnrollmentAccountResource> response = await collection.GetIfExistsAsync(enrollmentAccountName);
BillingDepartmentEnrollmentAccountResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    BillingEnrollmentAccountData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
