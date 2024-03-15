using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.MySql.FlexibleServers;
using Azure.ResourceManager.MySql.FlexibleServers.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/preview/2023-12-01-preview/examples/ServerUpdate.json
// this example is just showing the usage of "Servers_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MySqlFlexibleServerResource created on azure
// for more information of creating MySqlFlexibleServerResource, please refer to the document of MySqlFlexibleServerResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "testrg";
string serverName = "mysqltestserver";
ResourceIdentifier mySqlFlexibleServerResourceId = MySqlFlexibleServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
MySqlFlexibleServerResource mySqlFlexibleServer = client.GetMySqlFlexibleServerResource(mySqlFlexibleServerResourceId);

// invoke the operation
MySqlFlexibleServerPatch patch = new MySqlFlexibleServerPatch()
{
    Storage = new MySqlFlexibleServerStorage()
    {
        StorageSizeInGB = 30,
        Iops = 200,
        AutoGrow = MySqlFlexibleServerEnableStatusEnum.Disabled,
        AutoIoScaling = MySqlFlexibleServerEnableStatusEnum.Disabled,
    },
    Network = new MySqlFlexibleServerNetwork()
    {
        PublicNetworkAccess = MySqlFlexibleServerEnableStatusEnum.Disabled,
    },
};
ArmOperation<MySqlFlexibleServerResource> lro = await mySqlFlexibleServer.UpdateAsync(WaitUntil.Completed, patch);
MySqlFlexibleServerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MySqlFlexibleServerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
