using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataFactory.Models;
using Azure.ResourceManager.DataFactory;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_EnableInteractiveQuery.json
// this example is just showing the usage of "IntegrationRuntime_EnableInteractiveQuery" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataFactoryIntegrationRuntimeResource created on azure
// for more information of creating DataFactoryIntegrationRuntimeResource, please refer to the document of DataFactoryIntegrationRuntimeResource
string subscriptionId = "12345678-1234-1234-1234-12345678abc";
string resourceGroupName = "exampleResourceGroup";
string factoryName = "exampleFactoryName";
string integrationRuntimeName = "exampleIntegrationRuntime";
ResourceIdentifier dataFactoryIntegrationRuntimeResourceId = DataFactoryIntegrationRuntimeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, factoryName, integrationRuntimeName);
DataFactoryIntegrationRuntimeResource dataFactoryIntegrationRuntime = client.GetDataFactoryIntegrationRuntimeResource(dataFactoryIntegrationRuntimeResourceId);

// invoke the operation
EnableInteractiveQueryContent content = new EnableInteractiveQueryContent
{
    AutoTerminationMinutes = 10,
};
ArmOperation<DataFactoryIntegrationRuntimeResource> lro = await dataFactoryIntegrationRuntime.EnableInteractiveQueryIntegrationRuntimeAsync(WaitUntil.Completed, content);
DataFactoryIntegrationRuntimeResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataFactoryIntegrationRuntimeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
