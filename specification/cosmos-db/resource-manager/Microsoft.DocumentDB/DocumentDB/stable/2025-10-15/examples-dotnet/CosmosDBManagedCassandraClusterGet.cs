using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/CosmosDBManagedCassandraClusterGet.json
// this example is just showing the usage of "CassandraClusters_Get" operation, for the dependent resources, they will have to be created separately.

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
CassandraClusterResource result = await cassandraCluster.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CassandraClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
