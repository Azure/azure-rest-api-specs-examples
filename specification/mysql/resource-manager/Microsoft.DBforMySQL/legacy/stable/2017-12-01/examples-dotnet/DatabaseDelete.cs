using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MySql;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2017-12-01/examples/DatabaseDelete.json
// this example is just showing the usage of "Databases_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MySqlDatabaseResource created on azure
// for more information of creating MySqlDatabaseResource, please refer to the document of MySqlDatabaseResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestGroup";
string serverName = "testserver";
string databaseName = "db1";
ResourceIdentifier mySqlDatabaseResourceId = MySqlDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName);
MySqlDatabaseResource mySqlDatabase = client.GetMySqlDatabaseResource(mySqlDatabaseResourceId);

// invoke the operation
await mySqlDatabase.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
