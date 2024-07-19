using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MySql.Models;
using Azure.ResourceManager.MySql;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2020-01-01/examples/ServerKeyDelete.json
// this example is just showing the usage of "ServerKeys_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MySqlServerKeyResource created on azure
// for more information of creating MySqlServerKeyResource, please refer to the document of MySqlServerKeyResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "testrg";
string serverName = "testserver";
string keyName = "someVault_someKey_01234567890123456789012345678901";
ResourceIdentifier mySqlServerKeyResourceId = MySqlServerKeyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, keyName);
MySqlServerKeyResource mySqlServerKey = client.GetMySqlServerKeyResource(mySqlServerKeyResourceId);

// invoke the operation
await mySqlServerKey.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
