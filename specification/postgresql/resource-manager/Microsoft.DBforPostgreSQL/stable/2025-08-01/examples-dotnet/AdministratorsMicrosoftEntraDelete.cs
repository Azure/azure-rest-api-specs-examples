using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PostgreSql.FlexibleServers.Models;
using Azure.ResourceManager.PostgreSql.FlexibleServers;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/AdministratorsMicrosoftEntraDelete.json
// this example is just showing the usage of "AdministratorsMicrosoftEntra_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlFlexibleServerMicrosoftEntraAdministratorResource created on azure
// for more information of creating PostgreSqlFlexibleServerMicrosoftEntraAdministratorResource, please refer to the document of PostgreSqlFlexibleServerMicrosoftEntraAdministratorResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "exampleresourcegroup";
string serverName = "exampleserver";
string objectId = "oooooooo-oooo-oooo-oooo-oooooooooooo";
ResourceIdentifier postgreSqlFlexibleServerMicrosoftEntraAdministratorResourceId = PostgreSqlFlexibleServerMicrosoftEntraAdministratorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, objectId);
PostgreSqlFlexibleServerMicrosoftEntraAdministratorResource postgreSqlFlexibleServerMicrosoftEntraAdministrator = client.GetPostgreSqlFlexibleServerMicrosoftEntraAdministratorResource(postgreSqlFlexibleServerMicrosoftEntraAdministratorResourceId);

// invoke the operation
await postgreSqlFlexibleServerMicrosoftEntraAdministrator.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
