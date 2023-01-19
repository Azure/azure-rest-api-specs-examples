using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataFactory;
using Azure.ResourceManager.DataFactory.Models;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_CreateLinkedIntegrationRuntime.json
// this example is just showing the usage of "IntegrationRuntimes_CreateLinkedIntegrationRuntime" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FactoryIntegrationRuntimeResource created on azure
// for more information of creating FactoryIntegrationRuntimeResource, please refer to the document of FactoryIntegrationRuntimeResource
string subscriptionId = "12345678-1234-1234-1234-12345678abc";
string resourceGroupName = "exampleResourceGroup";
string factoryName = "exampleFactoryName";
string integrationRuntimeName = "exampleIntegrationRuntime";
ResourceIdentifier factoryIntegrationRuntimeResourceId = FactoryIntegrationRuntimeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, factoryName, integrationRuntimeName);
FactoryIntegrationRuntimeResource factoryIntegrationRuntime = client.GetFactoryIntegrationRuntimeResource(factoryIntegrationRuntimeResourceId);

// invoke the operation
CreateLinkedIntegrationRuntimeContent content = new CreateLinkedIntegrationRuntimeContent()
{
    Name = "bfa92911-9fb6-4fbe-8f23-beae87bc1c83",
    SubscriptionId = "061774c7-4b5a-4159-a55b-365581830283",
    DataFactoryName = "e9955d6d-56ea-4be3-841c-52a12c1a9981",
    DataFactoryLocation = new AzureLocation("West US"),
};
IntegrationRuntimeStatusResult result = await factoryIntegrationRuntime.CreateLinkedIntegrationRuntimeAsync(content);

Console.WriteLine($"Succeeded: {result}");
