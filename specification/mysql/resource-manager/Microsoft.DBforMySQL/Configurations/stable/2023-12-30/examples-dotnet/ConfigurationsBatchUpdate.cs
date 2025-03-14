using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.MySql.FlexibleServers.Models;
using Azure.ResourceManager.MySql.FlexibleServers;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/Configurations/stable/2023-12-30/examples/ConfigurationsBatchUpdate.json
// this example is just showing the usage of "Configurations_BatchUpdate" operation, for the dependent resources, they will have to be created separately.

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
MySqlFlexibleServerConfigurationListForBatchUpdate mySqlFlexibleServerConfigurationListForBatchUpdate = new MySqlFlexibleServerConfigurationListForBatchUpdate
{
    Values = {new MySqlFlexibleServerConfigurationForBatchUpdate
    {
    Name = "event_scheduler",
    Value = "OFF",
    }, new MySqlFlexibleServerConfigurationForBatchUpdate
    {
    Name = "div_precision_increment",
    Value = "8",
    }},
    ResetAllToDefault = MySqlFlexibleServerConfigurationResetAllToDefault.False,
};
ArmOperation<MySqlFlexibleServerConfigurations> lro = await mySqlFlexibleServer.UpdateConfigurationsAsync(WaitUntil.Completed, mySqlFlexibleServerConfigurationListForBatchUpdate);
MySqlFlexibleServerConfigurations result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
