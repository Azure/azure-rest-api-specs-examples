using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MySql;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2018-06-01/examples/WaitStatisticsGet.json
// this example is just showing the usage of "WaitStatistics_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MySqlWaitStatisticResource created on azure
// for more information of creating MySqlWaitStatisticResource, please refer to the document of MySqlWaitStatisticResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "testResourceGroupName";
string serverName = "testServerName";
string waitStatisticsId = "636927606000000000-636927615000000000-send-wait/io/socket/sql/client_connection-2--0";
ResourceIdentifier mySqlWaitStatisticResourceId = MySqlWaitStatisticResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, waitStatisticsId);
MySqlWaitStatisticResource mySqlWaitStatistic = client.GetMySqlWaitStatisticResource(mySqlWaitStatisticResourceId);

// invoke the operation
MySqlWaitStatisticResource result = await mySqlWaitStatistic.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MySqlWaitStatisticData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
