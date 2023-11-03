using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MySql.FlexibleServers;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/Databases/preview/2023-06-01-preview/examples/DatabaseDelete.json
// this example is just showing the usage of "Databases_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MySqlFlexibleServerDatabaseResource created on azure
// for more information of creating MySqlFlexibleServerDatabaseResource, please refer to the document of MySqlFlexibleServerDatabaseResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestGroup";
string serverName = "testserver";
string databaseName = "db1";
ResourceIdentifier mySqlFlexibleServerDatabaseResourceId = MySqlFlexibleServerDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName);
MySqlFlexibleServerDatabaseResource mySqlFlexibleServerDatabase = client.GetMySqlFlexibleServerDatabaseResource(mySqlFlexibleServerDatabaseResourceId);

// invoke the operation
await mySqlFlexibleServerDatabase.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
