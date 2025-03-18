using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.StreamAnalytics.Models;
using Azure.ResourceManager.StreamAnalytics;

// Generated from example definition: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Function_Update_AzureML.json
// this example is just showing the usage of "Functions_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StreamingJobFunctionResource created on azure
// for more information of creating StreamingJobFunctionResource, please refer to the document of StreamingJobFunctionResource
string subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
string resourceGroupName = "sjrg7";
string jobName = "sj9093";
string functionName = "function588";
ResourceIdentifier streamingJobFunctionResourceId = StreamingJobFunctionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, jobName, functionName);
StreamingJobFunctionResource streamingJobFunction = client.GetStreamingJobFunctionResource(streamingJobFunctionResourceId);

// invoke the operation
StreamingJobFunctionData data = new StreamingJobFunctionData
{
    Properties = new ScalarFunctionProperties
    {
        Binding = new EMachineLearningStudioFunctionBinding
        {
            BatchSize = 5000,
        },
    },
};
StreamingJobFunctionResource result = await streamingJobFunction.UpdateAsync(data);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StreamingJobFunctionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
