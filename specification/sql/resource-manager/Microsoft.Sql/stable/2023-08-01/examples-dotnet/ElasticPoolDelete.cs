using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/stable/2023-08-01/examples/ElasticPoolDelete.json
// this example is just showing the usage of "ElasticPools_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ElasticPoolResource created on azure
// for more information of creating ElasticPoolResource, please refer to the document of ElasticPoolResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "sqlcrudtest-3129";
string serverName = "sqlcrudtest-228";
string elasticPoolName = "sqlcrudtest-3851";
ResourceIdentifier elasticPoolResourceId = ElasticPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, elasticPoolName);
ElasticPoolResource elasticPool = client.GetElasticPoolResource(elasticPoolResourceId);

// invoke the operation
await elasticPool.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
