using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.StreamAnalytics.Models;
using Azure.ResourceManager.StreamAnalytics;

// Generated from example definition: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Function_Create_AzureML.json
// this example is just showing the usage of "Functions_CreateOrReplace" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StreamingJobResource created on azure
// for more information of creating StreamingJobResource, please refer to the document of StreamingJobResource
string subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
string resourceGroupName = "sjrg7";
string jobName = "sj9093";
ResourceIdentifier streamingJobResourceId = StreamingJobResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, jobName);
StreamingJobResource streamingJob = client.GetStreamingJobResource(streamingJobResourceId);

// get the collection of this StreamingJobFunctionResource
StreamingJobFunctionCollection collection = streamingJob.GetStreamingJobFunctions();

// invoke the operation
string functionName = "function588";
StreamingJobFunctionData data = new StreamingJobFunctionData
{
    Properties = new ScalarFunctionProperties
    {
        Inputs = {new StreamingJobFunctionInput
        {
        DataType = "nvarchar(max)",
        }},
        OutputDataType = "nvarchar(max)",
        Binding = new EMachineLearningStudioFunctionBinding
        {
            Endpoint = "someAzureMLEndpointURL",
            ApiKey = "someApiKey==",
            Inputs = new MachineLearningStudioInputs
            {
                Name = "input1",
                ColumnNames = {new MachineLearningStudioInputColumn
                {
                Name = "tweet",
                DataType = "string",
                MapTo = 0,
                }},
            },
            Outputs = {new MachineLearningStudioOutputColumn
            {
            Name = "Sentiment",
            DataType = "string",
            }},
            BatchSize = 1000,
        },
    },
};
ArmOperation<StreamingJobFunctionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, functionName, data);
StreamingJobFunctionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StreamingJobFunctionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
