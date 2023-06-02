using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CancelElasticPoolOperation.json
// this example is just showing the usage of "ElasticPoolOperations_Cancel" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ElasticPoolResource created on azure
// for more information of creating ElasticPoolResource, please refer to the document of ElasticPoolResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "sqlcrudtest-7398";
string serverName = "sqlcrudtest-6661";
string elasticPoolName = "testpool";
ResourceIdentifier elasticPoolResourceId = ElasticPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, elasticPoolName);
ElasticPoolResource elasticPool = client.GetElasticPoolResource(elasticPoolResourceId);

// invoke the operation
Guid operationId = Guid.Parse("f779414b-e748-4925-8cfe-c8598f7660ae");
await elasticPool.CancelElasticPoolOperationAsync(operationId);

Console.WriteLine($"Succeeded");
