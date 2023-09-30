using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MachineLearning;
using Azure.ResourceManager.MachineLearning.Models;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2023-06-01-preview/examples/Workspace/FeaturesetVersion/createOrUpdate.json
// this example is just showing the usage of "FeaturesetVersions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningFeatureSetContainerResource created on azure
// for more information of creating MachineLearningFeatureSetContainerResource, please refer to the document of MachineLearningFeatureSetContainerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "my-aml-workspace";
string name = "string";
ResourceIdentifier machineLearningFeatureSetContainerResourceId = MachineLearningFeatureSetContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, name);
MachineLearningFeatureSetContainerResource machineLearningFeatureSetContainer = client.GetMachineLearningFeatureSetContainerResource(machineLearningFeatureSetContainerResourceId);

// get the collection of this MachineLearningFeatureSetVersionResource
MachineLearningFeatureSetVersionCollection collection = machineLearningFeatureSetContainer.GetMachineLearningFeatureSetVersions();

// invoke the operation
string version = "string";
MachineLearningFeatureSetVersionData data = new MachineLearningFeatureSetVersionData(new MachineLearningFeatureSetVersionProperties()
{
    Entities =
    {
    "string"
    },
    MaterializationSettings = new MaterializationSettings()
    {
        Notification = new NotificationSetting()
        {
            EmailOn =
            {
            EmailNotificationEnableType.JobFailed
            },
            Emails =
            {
            "string"
            },
        },
        ResourceInstanceType = "string",
        Schedule = new MachineLearningRecurrenceTrigger(MachineLearningRecurrenceFrequency.Day, 1)
        {
            Schedule = new MachineLearningRecurrenceSchedule(new int[]
{
1
}, new int[]
{
1
})
            {
                MonthDays =
                {
                1
                },
                WeekDays =
                {
                MachineLearningDayOfWeek.Monday
                },
            },
            EndTime = "string",
            StartTime = "string",
            TimeZone = "string",
        },
        SparkConfiguration =
        {
        ["string"] = "string",
        },
        StoreType = MaterializationStoreType.Online,
    },
    SpecificationPath = "string",
    Stage = "string",
    IsAnonymous = false,
    IsArchived = false,
    Description = "string",
    Properties =
    {
    ["string"] = "string",
    },
    Tags =
    {
    ["string"] = "string",
    },
});
ArmOperation<MachineLearningFeatureSetVersionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, version, data);
MachineLearningFeatureSetVersionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineLearningFeatureSetVersionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
