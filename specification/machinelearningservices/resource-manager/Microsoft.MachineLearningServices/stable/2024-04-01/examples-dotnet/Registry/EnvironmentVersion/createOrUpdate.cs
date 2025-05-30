using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Registry/EnvironmentVersion/createOrUpdate.json
// this example is just showing the usage of "RegistryEnvironmentVersions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningRegistryEnvironmentContainerResource created on azure
// for more information of creating MachineLearningRegistryEnvironmentContainerResource, please refer to the document of MachineLearningRegistryEnvironmentContainerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string registryName = "my-aml-registry";
string environmentName = "string";
ResourceIdentifier machineLearningRegistryEnvironmentContainerResourceId = MachineLearningRegistryEnvironmentContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName, environmentName);
MachineLearningRegistryEnvironmentContainerResource machineLearningRegistryEnvironmentContainer = client.GetMachineLearningRegistryEnvironmentContainerResource(machineLearningRegistryEnvironmentContainerResourceId);

// get the collection of this MachineLearningRegistryEnvironmentVersionResource
MachineLearningRegistryEnvironmentVersionCollection collection = machineLearningRegistryEnvironmentContainer.GetMachineLearningRegistryEnvironmentVersions();

// invoke the operation
string version = "string";
MachineLearningEnvironmentVersionData data = new MachineLearningEnvironmentVersionData(new MachineLearningEnvironmentVersionProperties
{
    Image = "docker.io/tensorflow/serving:latest",
    CondaFile = "string",
    Build = new MachineLearningBuildContext(new Uri("https://storage-account.blob.core.windows.net/azureml/DockerBuildContext/95ddede6b9b8c4e90472db3acd0a8d28/"))
    {
        DockerfilePath = "prod/Dockerfile",
    },
    InferenceConfig = new MachineLearningInferenceContainerProperties
    {
        LivenessRoute = new MachineLearningInferenceContainerRoute("string", 1),
        ReadinessRoute = new MachineLearningInferenceContainerRoute("string", 1),
        ScoringRoute = new MachineLearningInferenceContainerRoute("string", 1),
    },
    IsAnonymous = false,
    Description = "string",
    Tags =
    {
    ["string"] = "string"
    },
    Properties =
    {
    ["string"] = "string"
    },
});
ArmOperation<MachineLearningRegistryEnvironmentVersionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, version, data);
MachineLearningRegistryEnvironmentVersionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineLearningEnvironmentVersionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
