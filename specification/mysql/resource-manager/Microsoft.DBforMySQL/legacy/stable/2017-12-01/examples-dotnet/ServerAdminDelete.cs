using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MySql;
using Azure.ResourceManager.MySql.Models;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2017-12-01/examples/ServerAdminDelete.json
// this example is just showing the usage of "ServerAdministrators_Delete" operation, for the dependent resources, they will have to be created separately.

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
await mySqlServerAdministrator.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
