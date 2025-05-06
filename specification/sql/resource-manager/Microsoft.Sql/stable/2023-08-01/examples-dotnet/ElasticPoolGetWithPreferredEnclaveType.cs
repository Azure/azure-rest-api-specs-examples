using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/stable/2023-08-01/examples/ElasticPoolGetWithPreferredEnclaveType.json
// this example is just showing the usage of "ElasticPools_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ElasticPoolResource created on azure
// for more information of creating ElasticPoolResource, please refer to the document of ElasticPoolResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "sqlcrudtest-2369";
string serverName = "sqlcrudtest-8069";
string elasticPoolName = "sqlcrudtest-8102";
ResourceIdentifier elasticPoolResourceId = ElasticPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, elasticPoolName);
ElasticPoolResource elasticPool = client.GetElasticPoolResource(elasticPoolResourceId);

// invoke the operation
ElasticPoolResource result = await elasticPool.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ElasticPoolData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
