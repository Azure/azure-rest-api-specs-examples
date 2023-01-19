using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevTestLabs;
using Azure.ResourceManager.DevTestLabs.Models;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ArtifactSources_CreateOrUpdate.json
// this example is just showing the usage of "ArtifactSources_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevTestLabResource created on azure
// for more information of creating DevTestLabResource, please refer to the document of DevTestLabResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "resourceGroupName";
string labName = "{labName}";
ResourceIdentifier devTestLabResourceId = DevTestLabResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, labName);
DevTestLabResource devTestLab = client.GetDevTestLabResource(devTestLabResourceId);

// get the collection of this DevTestLabArtifactSourceResource
DevTestLabArtifactSourceCollection collection = devTestLab.GetDevTestLabArtifactSources();

// invoke the operation
string name = "{artifactSourceName}";
DevTestLabArtifactSourceData data = new DevTestLabArtifactSourceData(new AzureLocation("placeholder"))
{
    DisplayName = "{displayName}",
    Uri = new Uri("{artifactSourceUri}"),
    SourceType = new DevTestLabSourceControlType("{VsoGit|GitHub|StorageAccount}"),
    FolderPath = "{folderPath}",
    ArmTemplateFolderPath = "{armTemplateFolderPath}",
    BranchRef = "{branchRef}",
    SecurityToken = "{securityToken}",
    Status = new DevTestLabEnableStatus("{Enabled|Disabled}"),
    Tags =
    {
    ["tagName1"] = "tagValue1",
    },
};
ArmOperation<DevTestLabArtifactSourceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, name, data);
DevTestLabArtifactSourceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DevTestLabArtifactSourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
