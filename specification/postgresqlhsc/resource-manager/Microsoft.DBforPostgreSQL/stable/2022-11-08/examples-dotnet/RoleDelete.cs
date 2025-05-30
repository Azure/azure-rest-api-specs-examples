using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDBForPostgreSql;

// Generated from example definition: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/RoleDelete.json
// this example is just showing the usage of "Roles_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBForPostgreSqlRoleResource created on azure
// for more information of creating CosmosDBForPostgreSqlRoleResource, please refer to the document of CosmosDBForPostgreSqlRoleResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestGroup";
string clusterName = "pgtestsvc4";
string roleName = "role1";
ResourceIdentifier cosmosDBForPostgreSqlRoleResourceId = CosmosDBForPostgreSqlRoleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, roleName);
CosmosDBForPostgreSqlRoleResource cosmosDBForPostgreSqlRole = client.GetCosmosDBForPostgreSqlRoleResource(cosmosDBForPostgreSqlRoleResourceId);

// invoke the operation
await cosmosDBForPostgreSqlRole.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
