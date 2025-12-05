using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataFactory.Models;
using Azure.ResourceManager.DataFactory;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/DataFactory/stable/2018-06-01/examples/DataFlows_Delete.json
// this example is just showing the usage of "DataFlows_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataFactoryDataFlowResource created on azure
// for more information of creating DataFactoryDataFlowResource, please refer to the document of DataFactoryDataFlowResource
string subscriptionId = "12345678-1234-1234-1234-12345678abc";
string resourceGroupName = "exampleResourceGroup";
string factoryName = "exampleFactoryName";
string dataFlowName = "exampleDataFlow";
ResourceIdentifier dataFactoryDataFlowResourceId = DataFactoryDataFlowResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, factoryName, dataFlowName);
DataFactoryDataFlowResource dataFactoryDataFlow = client.GetDataFactoryDataFlowResource(dataFactoryDataFlowResourceId);

// invoke the operation
await dataFactoryDataFlow.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
