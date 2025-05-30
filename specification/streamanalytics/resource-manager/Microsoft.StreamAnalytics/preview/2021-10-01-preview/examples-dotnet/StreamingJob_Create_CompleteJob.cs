using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.StreamAnalytics.Models;
using Azure.ResourceManager.StreamAnalytics;

// Generated from example definition: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/StreamingJob_Create_CompleteJob.json
// this example is just showing the usage of "StreamingJobs_CreateOrReplace" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
string resourceGroupName = "sjrg3276";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this StreamingJobResource
StreamingJobCollection collection = resourceGroupResource.GetStreamingJobs();

// invoke the operation
string jobName = "sj7804";
StreamingJobData data = new StreamingJobData(new AzureLocation("West US"))
{
    SkuName = StreamAnalyticsSkuName.Standard,
    EventsOutOfOrderPolicy = EventsOutOfOrderPolicy.Drop,
    OutputErrorPolicy = StreamingJobOutputErrorPolicy.Drop,
    EventsOutOfOrderMaxDelayInSeconds = 0,
    EventsLateArrivalMaxDelayInSeconds = 5,
    DataLocalion = new AzureLocation("en-US"),
    CompatibilityLevel = StreamingJobCompatibilityLevel.Level1_0,
    Inputs = {new StreamingJobInputData
    {
    Properties = new StreamInputProperties
    {
    Datasource = new BlobStreamInputDataSource
    {
    StorageAccounts = {new StreamAnalyticsStorageAccount
    {
    AccountName = "yourAccountName",
    AccountKey = "yourAccountKey==",
    }},
    Container = "containerName",
    PathPattern = "",
    },
    Serialization = new JsonFormatSerialization
    {
    Encoding = StreamAnalyticsDataSerializationEncoding.Utf8,
    },
    },
    Name = "inputtest",
    }},
    Transformation = new StreamingJobTransformationData
    {
        StreamingUnits = 1,
        Query = "Select Id, Name from inputtest",
        Name = "transformationtest",
    },
    Outputs = {new StreamingJobOutputData
    {
    Datasource = new SqlDatabaseOutputDataSource
    {
    Server = "serverName",
    Database = "databaseName",
    User = "<user>",
    Password = "userPassword",
    Table = "tableName",
    },
    Name = "outputtest",
    }},
    Functions = { },
    Externals = new StreamingJobExternal
    {
        StorageAccount = new StreamAnalyticsStorageAccount
        {
            AccountName = "mystorageaccount",
            AccountKey = "mykey",
        },
        Container = "mycontainer",
        Path = "UserCustomCode.zip",
        RefreshConfiguration = new StreamingJobRefreshConfiguration
        {
            PathPattern = "{date}\\\\{time}",
            DateFormat = "yyyy-dd-MM",
            TimeFormat = "HH",
            RefreshInterval = "00:01:00",
            RefreshType = DataRefreshType.Nonblocking,
        },
    },
    Tags =
    {
    ["key1"] = "value1",
    ["key3"] = "value3",
    ["randomKey"] = "randomValue"
    },
};
ArmOperation<StreamingJobResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, jobName, data);
StreamingJobResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StreamingJobData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
