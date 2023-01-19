using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.StreamAnalytics;
using Azure.ResourceManager.StreamAnalytics.Models;

// Generated from example definition: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Output_Create_EventHub_JSON.json
// this example is just showing the usage of "Outputs_CreateOrReplace" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StreamingJobResource created on azure
// for more information of creating StreamingJobResource, please refer to the document of StreamingJobResource
string subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
string resourceGroupName = "sjrg6912";
string jobName = "sj3310";
ResourceIdentifier streamingJobResourceId = StreamingJobResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, jobName);
StreamingJobResource streamingJob = client.GetStreamingJobResource(streamingJobResourceId);

// get the collection of this StreamingJobOutputResource
StreamingJobOutputCollection collection = streamingJob.GetStreamingJobOutputs();

// invoke the operation
string outputName = "output5195";
StreamingJobOutputData data = new StreamingJobOutputData()
{
    Datasource = new EventHubOutputDataSource()
    {
        ServiceBusNamespace = "sdktest",
        SharedAccessPolicyName = "RootManageSharedAccessKey",
        SharedAccessPolicyKey = "sharedAccessPolicyKey=",
        EventHubName = "sdkeventhub",
        PartitionKey = "partitionKey",
    },
    Serialization = new JsonFormatSerialization()
    {
        Encoding = StreamAnalyticsDataSerializationEncoding.Utf8,
        Format = JsonOutputSerializationFormat.Array,
    },
    WatermarkSettings = new StreamingJobOutputWatermarkProperties()
    {
        WatermarkMode = StreamingJobOutputWatermarkMode.SendCurrentPartitionWatermark,
        MaxWatermarkDifferenceAcrossPartitions = "16:14:30",
    },
};
ArmOperation<StreamingJobOutputResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, outputName, data);
StreamingJobOutputResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StreamingJobOutputData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
