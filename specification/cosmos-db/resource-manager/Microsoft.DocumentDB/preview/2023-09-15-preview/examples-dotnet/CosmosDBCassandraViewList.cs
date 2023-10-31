using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CosmosDB;
using Azure.ResourceManager.CosmosDB.Models;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-09-15-preview/examples/CosmosDBCassandraViewList.json
// this example is just showing the usage of "CassandraResources_ListCassandraViews" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CassandraKeyspaceResource created on azure
// for more information of creating CassandraKeyspaceResource, please refer to the document of CassandraKeyspaceResource
string subscriptionId = "subid";
string resourceGroupName = "rgName";
string accountName = "ddb1";
string keyspaceName = "keyspacename";
ResourceIdentifier cassandraKeyspaceResourceId = CassandraKeyspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, keyspaceName);
CassandraKeyspaceResource cassandraKeyspace = client.GetCassandraKeyspaceResource(cassandraKeyspaceResourceId);

// get the collection of this CassandraViewGetResultResource
CassandraViewGetResultCollection collection = cassandraKeyspace.GetCassandraViewGetResults();

// invoke the operation and iterate over the result
await foreach (CassandraViewGetResultResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    CassandraViewGetResultData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
