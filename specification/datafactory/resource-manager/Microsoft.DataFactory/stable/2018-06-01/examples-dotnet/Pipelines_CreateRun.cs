using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataFactory;
using Azure.ResourceManager.DataFactory.Models;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_CreateRun.json
// this example is just showing the usage of "Pipelines_CreateRun" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FactoryPipelineResource created on azure
// for more information of creating FactoryPipelineResource, please refer to the document of FactoryPipelineResource
string subscriptionId = "12345678-1234-1234-1234-12345678abc";
string resourceGroupName = "exampleResourceGroup";
string factoryName = "exampleFactoryName";
string pipelineName = "examplePipeline";
ResourceIdentifier factoryPipelineResourceId = FactoryPipelineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, factoryName, pipelineName);
FactoryPipelineResource factoryPipeline = client.GetFactoryPipelineResource(factoryPipelineResourceId);

// invoke the operation
IDictionary<string, BinaryData> parameterValueSpecification = new Dictionary<string, BinaryData>()
{
    ["OutputBlobNameList"] = BinaryData.FromObjectAsJson(new object[] { "exampleoutput.csv" }),
};
string referencePipelineRunId = null;
PipelineCreateRunResult result = await factoryPipeline.CreateRunAsync(parameterValueSpecification: parameterValueSpecification, referencePipelineRunId: referencePipelineRunId);

Console.WriteLine($"Succeeded: {result}");
