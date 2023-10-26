using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Kusto;
using Azure.ResourceManager.Kusto.Models;

// Generated from example definition: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDataConnectionsCosmosDbCreateOrUpdate.json
// this example is just showing the usage of "DataConnections_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this KustoDatabaseResource created on azure
// for more information of creating KustoDatabaseResource, please refer to the document of KustoDatabaseResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "kustorptest";
string clusterName = "kustoCluster";
string databaseName = "KustoDatabase1";
ResourceIdentifier kustoDatabaseResourceId = KustoDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, databaseName);
KustoDatabaseResource kustoDatabase = client.GetKustoDatabaseResource(kustoDatabaseResourceId);

// get the collection of this KustoDataConnectionResource
KustoDataConnectionCollection collection = kustoDatabase.GetKustoDataConnections();

// invoke the operation
string dataConnectionName = "dataConnectionTest";
KustoDataConnectionData data = new KustoCosmosDBDataConnection()
{
    TableName = "TestTable",
    MappingRuleName = "TestMapping",
    ManagedIdentityResourceId = new ResourceIdentifier("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1"),
    CosmosDBAccountResourceId = new ResourceIdentifier("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.DocumentDb/databaseAccounts/cosmosDbAccountTest1"),
    CosmosDBDatabase = "cosmosDbDatabaseTest",
    CosmosDBContainer = "cosmosDbContainerTest",
    RetrievalStartOn = DateTimeOffset.Parse("2022-07-29T12:00:00.6554616Z"),
    Location = new AzureLocation("westus"),
};
ArmOperation<KustoDataConnectionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, dataConnectionName, data);
KustoDataConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
KustoDataConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
