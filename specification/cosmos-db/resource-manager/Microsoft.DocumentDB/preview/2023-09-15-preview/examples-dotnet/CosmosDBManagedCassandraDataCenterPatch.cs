using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CosmosDB;
using Azure.ResourceManager.CosmosDB.Models;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-09-15-preview/examples/CosmosDBManagedCassandraDataCenterPatch.json
// this example is just showing the usage of "CassandraDataCenters_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CassandraDataCenterResource created on azure
// for more information of creating CassandraDataCenterResource, please refer to the document of CassandraDataCenterResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "cassandra-prod-rg";
string clusterName = "cassandra-prod";
string dataCenterName = "dc1";
ResourceIdentifier cassandraDataCenterResourceId = CassandraDataCenterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, dataCenterName);
CassandraDataCenterResource cassandraDataCenter = client.GetCassandraDataCenterResource(cassandraDataCenterResourceId);

// invoke the operation
CassandraDataCenterData data = new CassandraDataCenterData()
{
    Properties = new CassandraDataCenterProperties()
    {
        DataCenterLocation = new AzureLocation("West US 2"),
        DelegatedSubnetId = new ResourceIdentifier("/subscriptions/536e130b-d7d6-4ac7-98a5-de20d69588d2/resourceGroups/customer-vnet-rg/providers/Microsoft.Network/virtualNetworks/customer-vnet/subnets/dc1-subnet"),
        NodeCount = 9,
        Base64EncodedCassandraYamlFragment = "Y29tcGFjdGlvbl90aHJvdWdocHV0X21iX3Blcl9zZWM6IDMyCmNvbXBhY3Rpb25fbGFyZ2VfcGFydGl0aW9uX3dhcm5pbmdfdGhyZXNob2xkX21iOiAxMDA=",
    },
};
ArmOperation<CassandraDataCenterResource> lro = await cassandraDataCenter.UpdateAsync(WaitUntil.Completed, data);
CassandraDataCenterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CassandraDataCenterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
