using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-05-15-preview/examples/CosmosDBCassandraViewMigrateToAutoscale.json
// this example is just showing the usage of "CassandraResources_MigrateCassandraViewToAutoscale" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CassandraViewThroughputSettingResource created on azure
// for more information of creating CassandraViewThroughputSettingResource, please refer to the document of CassandraViewThroughputSettingResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string accountName = "ddb1";
string keyspaceName = "keyspacename";
string viewName = "viewname";
ResourceIdentifier cassandraViewThroughputSettingResourceId = CassandraViewThroughputSettingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, keyspaceName, viewName);
CassandraViewThroughputSettingResource cassandraViewThroughputSetting = client.GetCassandraViewThroughputSettingResource(cassandraViewThroughputSettingResourceId);

// invoke the operation
ArmOperation<CassandraViewThroughputSettingResource> lro = await cassandraViewThroughputSetting.MigrateCassandraViewToAutoscaleAsync(WaitUntil.Completed);
CassandraViewThroughputSettingResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ThroughputSettingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
