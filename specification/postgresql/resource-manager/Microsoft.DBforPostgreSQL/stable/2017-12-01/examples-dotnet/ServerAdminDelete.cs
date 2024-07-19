using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PostgreSql.Models;
using Azure.ResourceManager.PostgreSql;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerAdminDelete.json
// this example is just showing the usage of "ServerAdministrators_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlServerAdministratorResource created on azure
// for more information of creating PostgreSqlServerAdministratorResource, please refer to the document of PostgreSqlServerAdministratorResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "testrg";
string serverName = "pgtestsvc4";
ResourceIdentifier postgreSqlServerAdministratorResourceId = PostgreSqlServerAdministratorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
PostgreSqlServerAdministratorResource postgreSqlServerAdministrator = client.GetPostgreSqlServerAdministratorResource(postgreSqlServerAdministratorResourceId);

// invoke the operation
await postgreSqlServerAdministrator.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
