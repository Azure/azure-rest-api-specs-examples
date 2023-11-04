using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MySql;
using Azure.ResourceManager.MySql.Models;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2017-12-01/examples/ServerAdminCreateUpdate.json
// this example is just showing the usage of "ServerAdministrators_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MySqlServerAdministratorResource created on azure
// for more information of creating MySqlServerAdministratorResource, please refer to the document of MySqlServerAdministratorResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "testrg";
string serverName = "mysqltestsvc4";
ResourceIdentifier mySqlServerAdministratorResourceId = MySqlServerAdministratorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
MySqlServerAdministratorResource mySqlServerAdministrator = client.GetMySqlServerAdministratorResource(mySqlServerAdministratorResourceId);

// invoke the operation
MySqlServerAdministratorData data = new MySqlServerAdministratorData()
{
    AdministratorType = MySqlAdministratorType.ActiveDirectory,
    LoginAccountName = "bob@contoso.com",
    SecureId = Guid.Parse("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
    TenantId = Guid.Parse("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
};
ArmOperation<MySqlServerAdministratorResource> lro = await mySqlServerAdministrator.CreateOrUpdateAsync(WaitUntil.Completed, data);
MySqlServerAdministratorResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MySqlServerAdministratorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
