using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MySql.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.MySql;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2017-12-01/examples/ServerUpdate.json
// this example is just showing the usage of "Servers_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MySqlServerResource created on azure
// for more information of creating MySqlServerResource, please refer to the document of MySqlServerResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "testrg";
string serverName = "mysqltestsvc4";
ResourceIdentifier mySqlServerResourceId = MySqlServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
MySqlServerResource mySqlServer = client.GetMySqlServerResource(mySqlServerResourceId);

// invoke the operation
MySqlServerPatch patch = new MySqlServerPatch()
{
    AdministratorLoginPassword = "<administratorLoginPassword>",
    SslEnforcement = MySqlSslEnforcementEnum.Disabled,
};
ArmOperation<MySqlServerResource> lro = await mySqlServer.UpdateAsync(WaitUntil.Completed, patch);
MySqlServerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MySqlServerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
