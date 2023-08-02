using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Core.Expressions.DataFactory;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataFactory;
using Azure.ResourceManager.DataFactory.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/PipelineRuns_QueryByFactory.json
// this example is just showing the usage of "PipelineRuns_QueryByFactory" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation and iterate over the result
RunFilterContent content = new RunFilterContent(DateTimeOffset.Parse("2018-06-16T00:36:44.3345758Z"), DateTimeOffset.Parse("2018-06-16T00:49:48.3686473Z"))
{
    Filters =
    {
    new RunQueryFilter(RunQueryFilterOperand.PipelineName,RunQueryFilterOperator.EqualsValue,new string[]
    {
    "examplePipeline"
    })
    },
};
await foreach (DataFactoryPipelineRunInfo item in dataFactory.GetPipelineRunsAsync(content))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
