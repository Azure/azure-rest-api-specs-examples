using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.StreamAnalytics;
using Azure.ResourceManager.StreamAnalytics.Models;

// Generated from example definition: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Function_Delete.json
// this example is just showing the usage of "Functions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StreamingJobFunctionResource created on azure
// for more information of creating StreamingJobFunctionResource, please refer to the document of StreamingJobFunctionResource
string subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
string resourceGroupName = "sjrg1637";
string jobName = "sj8653";
string functionName = "function8197";
ResourceIdentifier streamingJobFunctionResourceId = StreamingJobFunctionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, jobName, functionName);
StreamingJobFunctionResource streamingJobFunction = client.GetStreamingJobFunctionResource(streamingJobFunctionResourceId);

// invoke the operation
await streamingJobFunction.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
