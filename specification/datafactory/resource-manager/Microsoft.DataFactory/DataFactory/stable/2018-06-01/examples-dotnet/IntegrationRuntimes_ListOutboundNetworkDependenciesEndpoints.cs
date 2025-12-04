using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataFactory.Models;
using Azure.ResourceManager.DataFactory;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_ListOutboundNetworkDependenciesEndpoints.json
// this example is just showing the usage of "IntegrationRuntimes_ListOutboundNetworkDependenciesEndpoints" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataFactoryIntegrationRuntimeResource created on azure
// for more information of creating DataFactoryIntegrationRuntimeResource, please refer to the document of DataFactoryIntegrationRuntimeResource
string subscriptionId = "7ad7c73b-38b8-4df3-84ee-52ff91092f61";
string resourceGroupName = "exampleResourceGroup";
string factoryName = "exampleFactoryName";
string integrationRuntimeName = "exampleIntegrationRuntime";
ResourceIdentifier dataFactoryIntegrationRuntimeResourceId = DataFactoryIntegrationRuntimeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, factoryName, integrationRuntimeName);
DataFactoryIntegrationRuntimeResource dataFactoryIntegrationRuntime = client.GetDataFactoryIntegrationRuntimeResource(dataFactoryIntegrationRuntimeResourceId);

// invoke the operation and iterate over the result
await foreach (IntegrationRuntimeOutboundNetworkDependenciesCategoryEndpoint item in dataFactoryIntegrationRuntime.GetOutboundNetworkDependenciesAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
