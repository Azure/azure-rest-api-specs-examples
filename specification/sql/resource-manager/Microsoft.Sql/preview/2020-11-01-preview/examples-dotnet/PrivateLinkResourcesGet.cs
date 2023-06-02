using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/PrivateLinkResourcesGet.json
// this example is just showing the usage of "PrivateLinkResources_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlServerResource created on azure
// for more information of creating SqlServerResource, please refer to the document of SqlServerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default";
string serverName = "test-svr";
ResourceIdentifier sqlServerResourceId = SqlServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
SqlServerResource sqlServer = client.GetSqlServerResource(sqlServerResourceId);

// get the collection of this SqlPrivateLinkResource
SqlPrivateLinkResourceCollection collection = sqlServer.GetSqlPrivateLinkResources();

// invoke the operation
string groupName = "plr";
bool result = await collection.ExistsAsync(groupName);

Console.WriteLine($"Succeeded: {result}");
