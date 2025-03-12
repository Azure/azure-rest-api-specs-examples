using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleDefinitionGetByEnrollmentAccount.json
// this example is just showing the usage of "BillingRoleDefinition_GetByEnrollmentAccount" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingEnrollmentAccountResource created on azure
// for more information of creating BillingEnrollmentAccountResource, please refer to the document of BillingEnrollmentAccountResource
string billingAccountName = "123456";
string enrollmentAccountName = "4568789";
ResourceIdentifier billingEnrollmentAccountResourceId = BillingEnrollmentAccountResource.CreateResourceIdentifier(billingAccountName, enrollmentAccountName);
BillingEnrollmentAccountResource billingEnrollmentAccount = client.GetBillingEnrollmentAccountResource(billingEnrollmentAccountResourceId);

// get the collection of this BillingEnrollmentAccountRoleDefinitionResource
BillingEnrollmentAccountRoleDefinitionCollection collection = billingEnrollmentAccount.GetBillingEnrollmentAccountRoleDefinitions();

// invoke the operation
string roleDefinitionName = "50000000-aaaa-bbbb-cccc-100000000000";
NullableResponse<BillingEnrollmentAccountRoleDefinitionResource> response = await collection.GetIfExistsAsync(roleDefinitionName);
BillingEnrollmentAccountRoleDefinitionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    BillingRoleDefinitionData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
