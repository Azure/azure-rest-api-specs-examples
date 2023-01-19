using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.StreamAnalytics;
using Azure.ResourceManager.StreamAnalytics.Models;

// Generated from example definition: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Input_Get_Stream_EventHub_JSON.json
// this example is just showing the usage of "Inputs_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StreamingJobInputResource created on azure
// for more information of creating StreamingJobInputResource, please refer to the document of StreamingJobInputResource
string subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
string resourceGroupName = "sjrg3139";
string jobName = "sj197";
string inputName = "input7425";
ResourceIdentifier streamingJobInputResourceId = StreamingJobInputResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, jobName, inputName);
StreamingJobInputResource streamingJobInput = client.GetStreamingJobInputResource(streamingJobInputResourceId);

// invoke the operation
StreamingJobInputResource result = await streamingJobInput.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StreamingJobInputData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
