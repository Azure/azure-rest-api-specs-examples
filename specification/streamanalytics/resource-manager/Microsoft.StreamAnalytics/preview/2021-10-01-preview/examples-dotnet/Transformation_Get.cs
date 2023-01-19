using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.StreamAnalytics;

// Generated from example definition: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Transformation_Get.json
// this example is just showing the usage of "Transformations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StreamingJobTransformationResource created on azure
// for more information of creating StreamingJobTransformationResource, please refer to the document of StreamingJobTransformationResource
string subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
string resourceGroupName = "sjrg4423";
string jobName = "sj8374";
string transformationName = "transformation952";
ResourceIdentifier streamingJobTransformationResourceId = StreamingJobTransformationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, jobName, transformationName);
StreamingJobTransformationResource streamingJobTransformation = client.GetStreamingJobTransformationResource(streamingJobTransformationResourceId);

// invoke the operation
StreamingJobTransformationResource result = await streamingJobTransformation.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StreamingJobTransformationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
