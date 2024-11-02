using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-09-01-preview/examples/CosmosDBSqlContainerRetrieveThroughputDistribution.json
// this example is just showing the usage of "SqlResources_SqlContainerRetrieveThroughputDistribution" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBSqlContainerThroughputSettingResource created on azure
// for more information of creating CosmosDBSqlContainerThroughputSettingResource, please refer to the document of CosmosDBSqlContainerThroughputSettingResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string accountName = "ddb1";
string databaseName = "databaseName";
string containerName = "containerName";
ResourceIdentifier cosmosDBSqlContainerThroughputSettingResourceId = CosmosDBSqlContainerThroughputSettingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, databaseName, containerName);
CosmosDBSqlContainerThroughputSettingResource cosmosDBSqlContainerThroughputSetting = client.GetCosmosDBSqlContainerThroughputSettingResource(cosmosDBSqlContainerThroughputSettingResourceId);

// invoke the operation
RetrieveThroughputParameters retrieveThroughputParameters = new RetrieveThroughputParameters(new AzureLocation("placeholder"), new RetrieveThroughputPropertiesResource(new WritableSubResource[]
{
new WritableSubResource()
{
Id = new ResourceIdentifier("0"),
},new WritableSubResource()
{
Id = new ResourceIdentifier("1"),
}
}));
ArmOperation<PhysicalPartitionThroughputInfoResult> lro = await cosmosDBSqlContainerThroughputSetting.SqlContainerRetrieveThroughputDistributionAsync(WaitUntil.Completed, retrieveThroughputParameters);
PhysicalPartitionThroughputInfoResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
