using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MachineLearning;
using Azure.ResourceManager.MachineLearning.Models;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2023-06-01-preview/examples/Job/CommandJob/createOrUpdate.json
// this example is just showing the usage of "Jobs_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningWorkspaceResource created on azure
// for more information of creating MachineLearningWorkspaceResource, please refer to the document of MachineLearningWorkspaceResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "my-aml-workspace";
ResourceIdentifier machineLearningWorkspaceResourceId = MachineLearningWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
MachineLearningWorkspaceResource machineLearningWorkspace = client.GetMachineLearningWorkspaceResource(machineLearningWorkspaceResourceId);

// get the collection of this MachineLearningJobResource
MachineLearningJobCollection collection = machineLearningWorkspace.GetMachineLearningJobs();

// invoke the operation
string id = "string";
MachineLearningJobData data = new MachineLearningJobData(new MachineLearningCommandJob("string", new ResourceIdentifier("string"))
{
    MlflowAutologger = MachineLearningFlowAutoLoggerState.Disabled,
    CodeId = new ResourceIdentifier("string"),
    Distribution = new TensorFlowDistributionConfiguration()
    {
        ParameterServerCount = 1,
        WorkerCount = 1,
    },
    EnvironmentVariables =
    {
    ["string"] = "string",
    },
    Inputs =
    {
    ["string"] = new MachineLearningLiteralJobInput("string")
    {
    Description = "string",
    },
    },
    Limits = new MachineLearningCommandJobLimits()
    {
        Timeout = XmlConvert.ToTimeSpan("PT5M"),
    },
    Outputs =
    {
    ["string"] = new MachineLearningUriFileJobOutput()
    {
    AssetName = "string",
    AssetVersion = "string",
    Mode = MachineLearningOutputDeliveryMode.Upload,
    Uri = new Uri("string"),
    Description = "string",
    },
    },
    QueueSettings = new JobQueueSettings()
    {
        JobTier = JobTier.Basic,
        Priority = 1,
    },
    Resources = new MachineLearningJobResourceConfiguration()
    {
        DockerArgs = "string",
        ShmSize = "2g",
        InstanceCount = 1,
        InstanceType = "string",
        Locations =
        {
        "string"
        },
        Properties =
        {
        ["string"] = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
        {
        ["c9ac10d0-915b-4de5-afe8-a4c78a37a558"] = null}),
        },
    },
    ComponentId = new ResourceIdentifier("string"),
    ComputeId = new ResourceIdentifier("string"),
    DisplayName = "string",
    ExperimentName = "string",
    Identity = new AmlToken(),
    IsArchived = false,
    NotificationSetting = new NotificationSetting()
    {
        EmailOn =
        {
        EmailNotificationEnableType.JobCancelled
        },
        Emails =
        {
        "string"
        },
    },
    Services =
    {
    ["string"] = new MachineLearningJobService()
    {
    Endpoint = "string",
    JobServiceType = "string",
    Nodes = new JobAllNodes(),
    Port = 1,
    Properties =
    {
    ["string"] = "string",
    },
    },
    },
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
ArmOperation<MachineLearningJobResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, id, data);
MachineLearningJobResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineLearningJobData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
