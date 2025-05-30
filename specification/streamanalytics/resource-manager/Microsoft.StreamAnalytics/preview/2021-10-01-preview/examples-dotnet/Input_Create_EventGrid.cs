using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.StreamAnalytics.Models;
using Azure.ResourceManager.StreamAnalytics;

// Generated from example definition: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Input_Create_EventGrid.json
// this example is just showing the usage of "Inputs_CreateOrReplace" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StreamingJobResource created on azure
// for more information of creating StreamingJobResource, please refer to the document of StreamingJobResource
string subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
string resourceGroupName = "sjrg3467";
string jobName = "sj9742";
ResourceIdentifier streamingJobResourceId = StreamingJobResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, jobName);
StreamingJobResource streamingJob = client.GetStreamingJobResource(streamingJobResourceId);

// get the collection of this StreamingJobInputResource
StreamingJobInputCollection collection = streamingJob.GetStreamingJobInputs();

// invoke the operation
string inputName = "input7970";
StreamingJobInputData input = new StreamingJobInputData
{
    Properties = new StreamInputProperties
    {
        Datasource = new EventGridStreamInputDataSource
        {
            Subscriber = new EventHubV2StreamInputDataSource
            {
                ServiceBusNamespace = "sdktest",
                SharedAccessPolicyName = "RootManageSharedAccessKey",
                SharedAccessPolicyKey = "someSharedAccessPolicyKey==",
                AuthenticationMode = StreamAnalyticsAuthenticationMode.Msi,
                EventHubName = "sdkeventhub",
                PartitionCount = 16,
                ConsumerGroupName = "sdkconsumergroup",
            },
            Schema = EventGridEventSchemaType.CloudEventSchema,
            StorageAccounts = {new StreamAnalyticsStorageAccount
            {
            AccountName = "myaccount",
            AccountKey = "myaccountkey",
            AuthenticationMode = StreamAnalyticsAuthenticationMode.Msi,
            }},
            EventTypes = { "Microsoft.Storage.BlobCreated" },
        },
    },
};
ArmOperation<StreamingJobInputResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, inputName, input);
StreamingJobInputResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StreamingJobInputData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
