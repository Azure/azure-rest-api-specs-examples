using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MachineLearning;
using Azure.ResourceManager.MachineLearning.Models;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2023-06-01-preview/examples/Workspace/ModelVersion/package.json
// this example is just showing the usage of "ModelVersions_Package" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningModelVersionResource created on azure
// for more information of creating MachineLearningModelVersionResource, please refer to the document of MachineLearningModelVersionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "my-aml-workspace";
string name = "string";
string version = "string";
ResourceIdentifier machineLearningModelVersionResourceId = MachineLearningModelVersionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, name, version);
MachineLearningModelVersionResource machineLearningModelVersion = client.GetMachineLearningModelVersionResource(machineLearningModelVersionResourceId);

// invoke the operation
ModelPackageContent content = new ModelPackageContent(new AzureMLBatchInferencingServer()
{
    CodeConfiguration = new MachineLearningCodeConfiguration("string")
    {
        CodeId = new ResourceIdentifier("string"),
    },
}, "string")
{
    BaseEnvironmentSource = new BaseEnvironmentType(new ResourceIdentifier("string")),
    EnvironmentVariables =
    {
    ["string"] = "string",
    },
    Inputs =
    {
    new ModelPackageInput(PackageInputType.UriFile,new PackageInputPathUri()
    {
    Uri = new Uri("string"),
    })
    {
    Mode = PackageInputDeliveryMode.Download,
    MountPath = "string",
    }
    },
    ModelConfiguration = new ModelConfiguration()
    {
        Mode = new PackageInputDeliveryMode("ReadOnlyMount"),
        MountPath = "string",
    },
    Tags =
    {
    ["string"] = "string",
    },
};
ArmOperation<ModelPackageResult> lro = await machineLearningModelVersion.PackageAsync(WaitUntil.Completed, content);
ModelPackageResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
