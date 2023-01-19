using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.StreamAnalytics;
using Azure.ResourceManager.StreamAnalytics.Models;

// Generated from example definition: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Output_Test.json
// this example is just showing the usage of "Outputs_Test" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StreamingJobOutputResource created on azure
// for more information of creating StreamingJobOutputResource, please refer to the document of StreamingJobOutputResource
string subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
string resourceGroupName = "sjrg2157";
string jobName = "sj6458";
string outputName = "output1755";
ResourceIdentifier streamingJobOutputResourceId = StreamingJobOutputResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, jobName, outputName);
StreamingJobOutputResource streamingJobOutput = client.GetStreamingJobOutputResource(streamingJobOutputResourceId);

// invoke the operation
ArmOperation<StreamAnalyticsResourceTestStatus> lro = await streamingJobOutput.TestAsync(WaitUntil.Completed);
StreamAnalyticsResourceTestStatus result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
