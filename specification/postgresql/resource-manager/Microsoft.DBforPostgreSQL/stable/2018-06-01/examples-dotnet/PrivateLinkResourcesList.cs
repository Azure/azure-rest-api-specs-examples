using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PostgreSql;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2018-06-01/examples/PrivateLinkResourcesList.json
// this example is just showing the usage of "PrivateLinkResources_ListByServer" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlServerResource created on azure
// for more information of creating PostgreSqlServerResource, please refer to the document of PostgreSqlServerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default";
string serverName = "test-svr";
ResourceIdentifier postgreSqlServerResourceId = PostgreSqlServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
PostgreSqlServerResource postgreSqlServer = client.GetPostgreSqlServerResource(postgreSqlServerResourceId);

// get the collection of this PostgreSqlPrivateLinkResource
PostgreSqlPrivateLinkResourceCollection collection = postgreSqlServer.GetPostgreSqlPrivateLinkResources();

// invoke the operation and iterate over the result
await foreach (PostgreSqlPrivateLinkResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    PostgreSqlPrivateLinkResourceData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
