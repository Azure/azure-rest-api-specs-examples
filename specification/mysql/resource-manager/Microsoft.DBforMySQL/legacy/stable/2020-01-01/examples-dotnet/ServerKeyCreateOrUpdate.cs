using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MySql;
using Azure.ResourceManager.MySql.Models;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2020-01-01/examples/ServerKeyCreateOrUpdate.json
// this example is just showing the usage of "ServerKeys_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MySqlServerResource created on azure
// for more information of creating MySqlServerResource, please refer to the document of MySqlServerResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "testrg";
string serverName = "testserver";
ResourceIdentifier mySqlServerResourceId = MySqlServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
MySqlServerResource mySqlServer = client.GetMySqlServerResource(mySqlServerResourceId);

// get the collection of this MySqlServerKeyResource
MySqlServerKeyCollection collection = mySqlServer.GetMySqlServerKeys();

// invoke the operation
string keyName = "someVault_someKey_01234567890123456789012345678901";
MySqlServerKeyData data = new MySqlServerKeyData()
{
    ServerKeyType = MySqlServerKeyType.AzureKeyVault,
    Uri = new Uri("https://someVault.vault.azure.net/keys/someKey/01234567890123456789012345678901"),
};
ArmOperation<MySqlServerKeyResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, keyName, data);
MySqlServerKeyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MySqlServerKeyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
