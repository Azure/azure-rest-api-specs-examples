using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevTestLabs;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ArmTemplates_Get.json
// this example is just showing the usage of "ArmTemplates_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevTestLabArtifactSourceResource created on azure
// for more information of creating DevTestLabArtifactSourceResource, please refer to the document of DevTestLabArtifactSourceResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "resourceGroupName";
string labName = "{labName}";
string artifactSourceName = "{artifactSourceName}";
ResourceIdentifier devTestLabArtifactSourceResourceId = DevTestLabArtifactSourceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, labName, artifactSourceName);
DevTestLabArtifactSourceResource devTestLabArtifactSource = client.GetDevTestLabArtifactSourceResource(devTestLabArtifactSourceResourceId);

// get the collection of this DevTestLabArmTemplateResource
DevTestLabArmTemplateCollection collection = devTestLabArtifactSource.GetDevTestLabArmTemplates();

// invoke the operation
string name = "{armTemplateName}";
bool result = await collection.ExistsAsync(name);

Console.WriteLine($"Succeeded: {result}");
