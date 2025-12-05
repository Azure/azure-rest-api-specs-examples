using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/CosmosDBManagedCassandraClusterPatch.json
// this example is just showing the usage of "CassandraClusters_Update" operation, for the dependent resources, they will have to be created separately.

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
CassandraClusterData data = new CassandraClusterData(default)
{
    Properties = new CassandraClusterProperties
    {
        AuthenticationMethod = CassandraAuthenticationMethod.None,
        ExternalGossipCertificates = {new CassandraCertificate
        {
        Pem = "-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----",
        }},
        ExternalSeedNodes = {new CassandraDataCenterSeedNode
        {
        IPAddress = "10.52.221.2",
        }, new CassandraDataCenterSeedNode
        {
        IPAddress = "10.52.221.3",
        }, new CassandraDataCenterSeedNode
        {
        IPAddress = "10.52.221.4",
        }},
        HoursBetweenBackups = 12,
    },
    Tags =
    {
    ["owner"] = "mike"
    },
};
ArmOperation<CassandraClusterResource> lro = await cassandraCluster.UpdateAsync(WaitUntil.Completed, data);
CassandraClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CassandraClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
