using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleDefinitionGetByBillingProfile.json
// this example is just showing the usage of "BillingRoleDefinition_GetByBillingProfile" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingProfileRoleDefinitionResource created on azure
// for more information of creating BillingProfileRoleDefinitionResource, please refer to the document of BillingProfileRoleDefinitionResource
string billingAccountName = "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
string billingProfileName = "xxxx-xxxx-xxx-xxx";
string roleDefinitionName = "40000000-aaaa-bbbb-cccc-100000000000";
ResourceIdentifier billingProfileRoleDefinitionResourceId = BillingProfileRoleDefinitionResource.CreateResourceIdentifier(billingAccountName, billingProfileName, roleDefinitionName);
BillingProfileRoleDefinitionResource billingProfileRoleDefinition = client.GetBillingProfileRoleDefinitionResource(billingProfileRoleDefinitionResourceId);

// invoke the operation
BillingProfileRoleDefinitionResource result = await billingProfileRoleDefinition.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BillingRoleDefinitionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
