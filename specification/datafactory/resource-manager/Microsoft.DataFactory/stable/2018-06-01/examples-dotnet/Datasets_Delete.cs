using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Core.Expressions.DataFactory;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataFactory;
using Azure.ResourceManager.DataFactory.Models;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Datasets_Delete.json
// this example is just showing the usage of "Datasets_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataFactoryDatasetResource created on azure
// for more information of creating DataFactoryDatasetResource, please refer to the document of DataFactoryDatasetResource
string subscriptionId = "12345678-1234-1234-1234-12345678abc";
string resourceGroupName = "exampleResourceGroup";
string factoryName = "exampleFactoryName";
string datasetName = "exampleDataset";
ResourceIdentifier dataFactoryDatasetResourceId = DataFactoryDatasetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, factoryName, datasetName);
DataFactoryDatasetResource dataFactoryDataset = client.GetDataFactoryDatasetResource(dataFactoryDatasetResourceId);

// invoke the operation
await dataFactoryDataset.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
