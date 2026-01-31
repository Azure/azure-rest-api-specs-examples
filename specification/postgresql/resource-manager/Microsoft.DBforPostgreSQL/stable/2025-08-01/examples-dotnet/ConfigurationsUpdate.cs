using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PostgreSql.FlexibleServers;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/ConfigurationsUpdate.json
// this example is just showing the usage of "Configurations_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlFlexibleServerConfigurationResource created on azure
// for more information of creating PostgreSqlFlexibleServerConfigurationResource, please refer to the document of PostgreSqlFlexibleServerConfigurationResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "exampleresourcegroup";
string serverName = "exampleserver";
string configurationName = "constraint_exclusion";
ResourceIdentifier postgreSqlFlexibleServerConfigurationResourceId = PostgreSqlFlexibleServerConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, configurationName);
PostgreSqlFlexibleServerConfigurationResource postgreSqlFlexibleServerConfiguration = client.GetPostgreSqlFlexibleServerConfigurationResource(postgreSqlFlexibleServerConfigurationResourceId);

// invoke the operation
PostgreSqlFlexibleServerConfigurationData data = new PostgreSqlFlexibleServerConfigurationData
{
    Value = "on",
    Source = "user-override",
};
ArmOperation<PostgreSqlFlexibleServerConfigurationResource> lro = await postgreSqlFlexibleServerConfiguration.UpdateAsync(WaitUntil.Completed, data);
PostgreSqlFlexibleServerConfigurationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PostgreSqlFlexibleServerConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
