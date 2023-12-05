using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.PostgreSql.FlexibleServers;
using Azure.ResourceManager.PostgreSql.FlexibleServers.Models;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-01-preview/examples/AdministratorDelete.json
// this example is just showing the usage of "Administrators_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlFlexibleServerActiveDirectoryAdministratorResource created on azure
// for more information of creating PostgreSqlFlexibleServerActiveDirectoryAdministratorResource, please refer to the document of PostgreSqlFlexibleServerActiveDirectoryAdministratorResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "testrg";
string serverName = "testserver";
string objectId = "oooooooo-oooo-oooo-oooo-oooooooooooo";
ResourceIdentifier postgreSqlFlexibleServerActiveDirectoryAdministratorResourceId = PostgreSqlFlexibleServerActiveDirectoryAdministratorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, objectId);
PostgreSqlFlexibleServerActiveDirectoryAdministratorResource postgreSqlFlexibleServerActiveDirectoryAdministrator = client.GetPostgreSqlFlexibleServerActiveDirectoryAdministratorResource(postgreSqlFlexibleServerActiveDirectoryAdministratorResourceId);

// invoke the operation
await postgreSqlFlexibleServerActiveDirectoryAdministrator.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
