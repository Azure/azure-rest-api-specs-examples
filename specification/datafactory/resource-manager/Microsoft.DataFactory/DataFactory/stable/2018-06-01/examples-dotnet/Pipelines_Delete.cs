using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Core.Expressions.DataFactory;
using Azure.Identity;
using Azure.ResourceManager.DataFactory.Models;
using Azure.ResourceManager.DataFactory;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/DataFactory/stable/2018-06-01/examples/Pipelines_Delete.json
// this example is just showing the usage of "Pipelines_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataFactoryPipelineResource created on azure
// for more information of creating DataFactoryPipelineResource, please refer to the document of DataFactoryPipelineResource
string subscriptionId = "12345678-1234-1234-1234-12345678abc";
string resourceGroupName = "exampleResourceGroup";
string factoryName = "exampleFactoryName";
string pipelineName = "examplePipeline";
ResourceIdentifier dataFactoryPipelineResourceId = DataFactoryPipelineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, factoryName, pipelineName);
DataFactoryPipelineResource dataFactoryPipeline = client.GetDataFactoryPipelineResource(dataFactoryPipelineResourceId);

// invoke the operation
await dataFactoryPipeline.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
