using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-09-01-preview/examples/CosmosDBManagedCassandraCommand.json
// this example is just showing the usage of "CassandraClusters_InvokeCommand" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CassandraClusterResource created on azure
// for more information of creating CassandraClusterResource, please refer to the document of CassandraClusterResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "cassandra-prod-rg";
string clusterName = "cassandra-prod";
ResourceIdentifier cassandraClusterResourceId = CassandraClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
CassandraClusterResource cassandraCluster = client.GetCassandraClusterResource(cassandraClusterResourceId);

// invoke the operation
CassandraCommandPostBody body = new CassandraCommandPostBody("nodetool", "10.0.1.12")
{
    Arguments =
    {
    ["status"] = "",
    },
};
ArmOperation<CassandraCommandOutput> lro = await cassandraCluster.InvokeCommandAsync(WaitUntil.Completed, body);
CassandraCommandOutput result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
