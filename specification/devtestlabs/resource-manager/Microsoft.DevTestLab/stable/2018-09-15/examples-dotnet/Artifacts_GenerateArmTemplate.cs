using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DevTestLabs.Models;
using Azure.ResourceManager.DevTestLabs;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Artifacts_GenerateArmTemplate.json
// this example is just showing the usage of "Artifacts_GenerateArmTemplate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevTestLabArtifactResource created on azure
// for more information of creating DevTestLabArtifactResource, please refer to the document of DevTestLabArtifactResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "resourceGroupName";
string labName = "{labName}";
string artifactSourceName = "{artifactSourceName}";
string name = "{artifactName}";
ResourceIdentifier devTestLabArtifactResourceId = DevTestLabArtifactResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, labName, artifactSourceName, name);
DevTestLabArtifactResource devTestLabArtifact = client.GetDevTestLabArtifactResource(devTestLabArtifactResourceId);

// invoke the operation
DevTestLabArtifactGenerateArmTemplateContent content = new DevTestLabArtifactGenerateArmTemplateContent
{
    VmName = "{vmName}",
    Location = new AzureLocation("{location}"),
    FileUploadOptions = DevTestLabFileUploadOption.None,
};
DevTestLabArmTemplateInfo result = await devTestLabArtifact.GenerateArmTemplateAsync(content);

Console.WriteLine($"Succeeded: {result}");
