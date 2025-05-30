using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleAssignmentCreateOrUpdateByDepartment.json
// this example is just showing the usage of "BillingRoleAssignments_CreateOrUpdateByDepartment" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingDepartmentResource created on azure
// for more information of creating BillingDepartmentResource, please refer to the document of BillingDepartmentResource
string billingAccountName = "7898901";
string departmentName = "12345";
ResourceIdentifier billingDepartmentResourceId = BillingDepartmentResource.CreateResourceIdentifier(billingAccountName, departmentName);
BillingDepartmentResource billingDepartment = client.GetBillingDepartmentResource(billingDepartmentResourceId);

// get the collection of this BillingDepartmentRoleAssignmentResource
BillingDepartmentRoleAssignmentCollection collection = billingDepartment.GetBillingDepartmentRoleAssignments();

// invoke the operation
string billingRoleAssignmentName = "9dfd08c2-62a3-4d47-85bd-1cdba1408402";
BillingRoleAssignmentData data = new BillingRoleAssignmentData
{
    Properties = new BillingRoleAssignmentProperties(new ResourceIdentifier("/providers/Microsoft.Billing/billingAccounts/7898901/departments/12345/billingRoleDefinitions/9f1983cb-2574-400c-87e9-34cf8e2280db"))
    {
        PrincipalId = "00000000-0000-0000-0000-000000000000",
        PrincipalTenantId = Guid.Parse("076915e7-de10-4323-bb34-a58c904068bb"),
        UserEmailAddress = "john@contoso.com",
    },
};
ArmOperation<BillingDepartmentRoleAssignmentResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, billingRoleAssignmentName, data);
BillingDepartmentRoleAssignmentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BillingRoleAssignmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
