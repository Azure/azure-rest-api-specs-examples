using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.StreamAnalytics;
using Azure.ResourceManager.StreamAnalytics.Models;

// Generated from example definition: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/StreamingJob_Delete.json
// this example is just showing the usage of "StreamingJobs_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StreamingJobResource created on azure
// for more information of creating StreamingJobResource, please refer to the document of StreamingJobResource
string subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
string resourceGroupName = "sjrg6936";
string jobName = "sj59";
ResourceIdentifier streamingJobResourceId = StreamingJobResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, jobName);
StreamingJobResource streamingJob = client.GetStreamingJobResource(streamingJobResourceId);

// invoke the operation
await streamingJob.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
