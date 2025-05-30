using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MySql;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2018-06-01/examples/TopQueryStatisticsGet.json
// this example is just showing the usage of "TopQueryStatistics_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MySqlQueryStatisticResource created on azure
// for more information of creating MySqlQueryStatisticResource, please refer to the document of MySqlQueryStatisticResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "testResourceGroupName";
string serverName = "testServerName";
string queryStatisticId = "66-636923268000000000-636923277000000000-avg-duration";
ResourceIdentifier mySqlQueryStatisticResourceId = MySqlQueryStatisticResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, queryStatisticId);
MySqlQueryStatisticResource mySqlQueryStatistic = client.GetMySqlQueryStatisticResource(mySqlQueryStatisticResourceId);

// invoke the operation
MySqlQueryStatisticResource result = await mySqlQueryStatistic.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MySqlQueryStatisticData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
