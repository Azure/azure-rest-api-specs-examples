using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Kusto;
using Azure.ResourceManager.Kusto.Models;

// Generated from example definition: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoSuspendedDatabasesGet.json
// this example is just showing the usage of "Databases_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this KustoClusterResource created on azure
// for more information of creating KustoClusterResource, please refer to the document of KustoClusterResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "kustorptest";
string clusterName = "kustoCluster";
ResourceIdentifier kustoClusterResourceId = KustoClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
KustoClusterResource kustoCluster = client.GetKustoClusterResource(kustoClusterResourceId);

// get the collection of this KustoDatabaseResource
KustoDatabaseCollection collection = kustoCluster.GetKustoDatabases();

// invoke the operation
string databaseName = "KustoDatabase9";
NullableResponse<KustoDatabaseResource> response = await collection.GetIfExistsAsync(databaseName);
KustoDatabaseResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    KustoDatabaseData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
