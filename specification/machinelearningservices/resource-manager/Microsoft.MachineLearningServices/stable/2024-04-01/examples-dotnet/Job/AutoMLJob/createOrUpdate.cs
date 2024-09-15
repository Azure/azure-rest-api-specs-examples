using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Job/AutoMLJob/createOrUpdate.json
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
MachineLearningJobData data = new MachineLearningJobData(new AutoMLJob(new ImageClassification(new MachineLearningTableJobInput(new Uri("string")), new ImageLimitSettings()
{
    MaxTrials = 2,
})
{
    ModelSettings = new ImageModelSettingsClassification()
    {
        ValidationCropSize = 2,
    },
    SearchSpace =
    {
    new ImageModelDistributionSettingsClassification()
    {
    ValidationCropSize = "choice(2, 360)",
    }
    },
    TargetColumnName = "string",
})
{
    Resources = new MachineLearningJobResourceConfiguration()
    {
        InstanceCount = 1,
        InstanceType = "string",
        Properties =
        {
        ["string"] = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
        {
        ["9bec0ab0-c62f-4fa9-a97c-7b24bbcc90ad"] = null}),
        },
    },
    EnvironmentId = "string",
    EnvironmentVariables =
    {
    ["string"] = "string",
    },
    Outputs =
    {
    ["string"] = new MachineLearningUriFileJobOutput()
    {
    Uri = new Uri("string"),
    Mode = MachineLearningOutputDeliveryMode.ReadWriteMount,
    Description = "string",
    },
    },
    DisplayName = "string",
    ExperimentName = "string",
    Services =
    {
    ["string"] = new MachineLearningJobService()
    {
    JobServiceType = "string",
    Port = 1,
    Endpoint = "string",
    Properties =
    {
    ["string"] = "string",
    },
    },
    },
    ComputeId = new ResourceIdentifier("string"),
    IsArchived = false,
    Identity = new AmlToken(),
    Description = "string",
    Tags =
    {
    ["string"] = "string",
    },
    Properties =
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
