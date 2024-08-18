using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataFactory.Models;
using Azure.ResourceManager.DataFactory;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_Create.json
// this example is just showing the usage of "Pipelines_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataFactoryResource created on azure
// for more information of creating DataFactoryResource, please refer to the document of DataFactoryResource
string subscriptionId = "12345678-1234-1234-1234-12345678abc";
string resourceGroupName = "exampleResourceGroup";
string factoryName = "exampleFactoryName";
ResourceIdentifier dataFactoryResourceId = DataFactoryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, factoryName);
DataFactoryResource dataFactory = client.GetDataFactoryResource(dataFactoryResourceId);

// get the collection of this DataFactoryPipelineResource
DataFactoryPipelineCollection collection = dataFactory.GetDataFactoryPipelines();

// invoke the operation
string pipelineName = "examplePipeline";
DataFactoryPipelineData data = new DataFactoryPipelineData()
{
    Activities =
    {
    new ForEachActivity("ExampleForeachActivity",new DataFactoryExpression(DataFactoryExpressionType.Expression,"@pipeline().parameters.OutputBlobNameList"),new PipelineActivity[]
    {
    new CopyActivity("ExampleCopyActivity",new DataFactoryBlobSource(),new DataFactoryBlobSink())
    {
    Inputs =
    {
    new DatasetReference(DatasetReferenceType.DatasetReference,"exampleDataset")
    {
    Parameters =
    {
    ["MyFileName"] = BinaryData.FromString("\"examplecontainer.csv\""),
    ["MyFolderPath"] = BinaryData.FromString("\"examplecontainer\""),
    },
    }
    },
    Outputs =
    {
    new DatasetReference(DatasetReferenceType.DatasetReference,"exampleDataset")
    {
    Parameters =
    {
    ["MyFileName"] = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
    {
    ["type"] = "Expression",
    ["value"] = "@item()"}),
    ["MyFolderPath"] = BinaryData.FromString("\"examplecontainer\""),
    },
    }
    },
    DataIntegrationUnits = 32,
    }
    })
    {
    IsSequential = true,
    }
    },
    Parameters =
    {
    ["JobId"] = new EntityParameterSpecification(EntityParameterType.String),
    ["OutputBlobNameList"] = new EntityParameterSpecification(EntityParameterType.Array),
    },
    Variables =
    {
    ["TestVariableArray"] = new PipelineVariableSpecification(PipelineVariableType.Array),
    },
    RunDimensions =
    {
    ["JobId"] = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
    {
    ["type"] = "Expression",
    ["value"] = "@pipeline().parameters.JobId"}),
    },
    ElapsedTimeMetricDuration = BinaryData.FromString("\"0.00:10:00\""),
};
ArmOperation<DataFactoryPipelineResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, pipelineName, data);
DataFactoryPipelineResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataFactoryPipelineData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
