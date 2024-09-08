using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/OnlineDeployment/KubernetesOnlineDeployment/createOrUpdate.json
// this example is just showing the usage of "OnlineDeployments_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningOnlineEndpointResource created on azure
// for more information of creating MachineLearningOnlineEndpointResource, please refer to the document of MachineLearningOnlineEndpointResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "my-aml-workspace";
string endpointName = "testEndpointName";
ResourceIdentifier machineLearningOnlineEndpointResourceId = MachineLearningOnlineEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, endpointName);
MachineLearningOnlineEndpointResource machineLearningOnlineEndpoint = client.GetMachineLearningOnlineEndpointResource(machineLearningOnlineEndpointResourceId);

// get the collection of this MachineLearningOnlineDeploymentResource
MachineLearningOnlineDeploymentCollection collection = machineLearningOnlineEndpoint.GetMachineLearningOnlineDeployments();

// invoke the operation
string deploymentName = "testDeploymentName";
MachineLearningOnlineDeploymentData data = new MachineLearningOnlineDeploymentData(new AzureLocation("string"), new MachineLearningKubernetesOnlineDeployment()
{
    ContainerResourceRequirements = new MachineLearningContainerResourceRequirements()
    {
        ContainerResourceRequests = new MachineLearningContainerResourceSettings()
        {
            Cpu = "\"1\"",
            Memory = "\"2Gi\"",
            Gpu = "\"1\"",
        },
        ContainerResourceLimits = new MachineLearningContainerResourceSettings()
        {
            Cpu = "\"1\"",
            Memory = "\"2Gi\"",
            Gpu = "\"1\"",
        },
    },
    ScaleSettings = new MachineLearningDefaultScaleSettings(),
    RequestSettings = new MachineLearningOnlineRequestSettings()
    {
        MaxQueueWait = XmlConvert.ToTimeSpan("PT5M"),
        RequestTimeout = XmlConvert.ToTimeSpan("PT5M"),
        MaxConcurrentRequestsPerInstance = 1,
    },
    ModelMountPath = "string",
    AppInsightsEnabled = false,
    LivenessProbe = new MachineLearningProbeSettings()
    {
        FailureThreshold = 1,
        SuccessThreshold = 1,
        Timeout = XmlConvert.ToTimeSpan("PT5M"),
        Period = XmlConvert.ToTimeSpan("PT5M"),
        InitialDelay = XmlConvert.ToTimeSpan("PT5M"),
    },
    InstanceType = "string",
    Model = "string",
    Description = "string",
    Properties =
    {
    ["string"] = "string",
    },
    CodeConfiguration = new MachineLearningCodeConfiguration("string")
    {
        CodeId = new ResourceIdentifier("string"),
    },
    EnvironmentId = "string",
    EnvironmentVariables =
    {
    ["string"] = "string",
    },
})
{
    Kind = "string",
    Identity = new ManagedServiceIdentity("SystemAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("string")] = new UserAssignedIdentity(),
        },
    },
    Sku = new MachineLearningSku("string")
    {
        Tier = MachineLearningSkuTier.Free,
        Size = "string",
        Family = "string",
        Capacity = 1,
    },
    Tags =
    {
    },
};
ArmOperation<MachineLearningOnlineDeploymentResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, deploymentName, data);
MachineLearningOnlineDeploymentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineLearningOnlineDeploymentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
