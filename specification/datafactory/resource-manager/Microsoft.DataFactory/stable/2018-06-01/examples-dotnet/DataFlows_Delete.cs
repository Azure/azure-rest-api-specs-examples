using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataFactory;
using Azure.ResourceManager.DataFactory.Models;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlows_Delete.json
// this example is just showing the usage of "DataFlows_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FactoryDataFlowResource created on azure
// for more information of creating FactoryDataFlowResource, please refer to the document of FactoryDataFlowResource
string subscriptionId = "12345678-1234-1234-1234-12345678abc";
string resourceGroupName = "exampleResourceGroup";
string factoryName = "exampleFactoryName";
string dataFlowName = "exampleDataFlow";
ResourceIdentifier factoryDataFlowResourceId = FactoryDataFlowResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, factoryName, dataFlowName);
FactoryDataFlowResource factoryDataFlow = client.GetFactoryDataFlowResource(factoryDataFlowResourceId);

// invoke the operation
await factoryDataFlow.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
