using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleAssignmentCreateOrUpdateByEnrollmentAccount.json
// this example is just showing the usage of "BillingRoleAssignments_CreateOrUpdateByEnrollmentAccount" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingEnrollmentAccountRoleAssignmentResource created on azure
// for more information of creating BillingEnrollmentAccountRoleAssignmentResource, please refer to the document of BillingEnrollmentAccountRoleAssignmentResource
string billingAccountName = "7898901";
string enrollmentAccountName = "123456";
string billingRoleAssignmentName = "9dfd08c2-62a3-4d47-85bd-1cdba1408402";
ResourceIdentifier billingEnrollmentAccountRoleAssignmentResourceId = BillingEnrollmentAccountRoleAssignmentResource.CreateResourceIdentifier(billingAccountName, enrollmentAccountName, billingRoleAssignmentName);
BillingEnrollmentAccountRoleAssignmentResource billingEnrollmentAccountRoleAssignment = client.GetBillingEnrollmentAccountRoleAssignmentResource(billingEnrollmentAccountRoleAssignmentResourceId);

// invoke the operation
BillingRoleAssignmentData data = new BillingRoleAssignmentData
{
    Properties = new BillingRoleAssignmentProperties(new ResourceIdentifier("/providers/Microsoft.Billing/billingAccounts/7898901/enrollmentAccounts/123456/billingRoleDefinitions/9f1983cb-2574-400c-87e9-34cf8e2280db"))
    {
        PrincipalId = "00000000-0000-0000-0000-000000000000",
        PrincipalTenantId = Guid.Parse("076915e7-de10-4323-bb34-a58c904068bb"),
        UserEmailAddress = "john@contoso.com",
    },
};
ArmOperation<BillingEnrollmentAccountRoleAssignmentResource> lro = await billingEnrollmentAccountRoleAssignment.UpdateAsync(WaitUntil.Completed, data);
BillingEnrollmentAccountRoleAssignmentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BillingRoleAssignmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
