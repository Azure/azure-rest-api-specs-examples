using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MySql.FlexibleServers.Models;
using Azure.ResourceManager.MySql.FlexibleServers;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/AAD/stable/2023-12-30/examples/AzureADAdministratorCreate.json
// this example is just showing the usage of "AzureADAdministrators_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MySqlFlexibleServerAadAdministratorResource created on azure
// for more information of creating MySqlFlexibleServerAadAdministratorResource, please refer to the document of MySqlFlexibleServerAadAdministratorResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "testrg";
string serverName = "mysqltestsvc4";
MySqlFlexibleServerAdministratorName administratorName = MySqlFlexibleServerAdministratorName.ActiveDirectory;
ResourceIdentifier mySqlFlexibleServerAadAdministratorResourceId = MySqlFlexibleServerAadAdministratorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, administratorName);
MySqlFlexibleServerAadAdministratorResource mySqlFlexibleServerAadAdministrator = client.GetMySqlFlexibleServerAadAdministratorResource(mySqlFlexibleServerAadAdministratorResourceId);

// invoke the operation
MySqlFlexibleServerAadAdministratorData data = new MySqlFlexibleServerAadAdministratorData
{
    AdministratorType = MySqlFlexibleServerAdministratorType.ActiveDirectory,
    Login = "bob@contoso.com",
    Sid = "c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c",
    TenantId = Guid.Parse("c12b7025-bfe2-46c1-b463-993b5e4cd467"),
    IdentityResourceId = new ResourceIdentifier("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/test-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-umi"),
};
ArmOperation<MySqlFlexibleServerAadAdministratorResource> lro = await mySqlFlexibleServerAadAdministrator.UpdateAsync(WaitUntil.Completed, data);
MySqlFlexibleServerAadAdministratorResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MySqlFlexibleServerAadAdministratorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
