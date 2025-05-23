using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/CosmosDBManagedCassandraClusterCreate.json
// this example is just showing the usage of "CassandraClusters_CreateUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "cassandra-prod-rg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this CassandraClusterResource
CassandraClusterCollection collection = resourceGroupResource.GetCassandraClusters();

// invoke the operation
string clusterName = "cassandra-prod";
CassandraClusterData data = new CassandraClusterData(new AzureLocation("West US"))
{
    Properties = new CassandraClusterProperties
    {
        DelegatedManagementSubnetId = new ResourceIdentifier("/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/management"),
        CassandraVersion = "3.11",
        ClusterNameOverride = "ClusterNameIllegalForAzureResource",
        AuthenticationMethod = CassandraAuthenticationMethod.Cassandra,
        InitialCassandraAdminPassword = "mypassword",
        ClientCertificates = {new CassandraCertificate
        {
        Pem = "-----BEGIN CERTIFICATE-----\n...Base64 encoded certificate...\n-----END CERTIFICATE-----",
        }},
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
        HoursBetweenBackups = 24,
    },
    Tags = { },
};
ArmOperation<CassandraClusterResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, clusterName, data);
CassandraClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CassandraClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
