using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevTestLabs;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Artifacts_Get.json
// this example is just showing the usage of "Artifacts_Get" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this DevTestLabArtifactResource
DevTestLabArtifactCollection collection = devTestLabArtifactSource.GetDevTestLabArtifacts();

// invoke the operation
string name = "{artifactName}";
NullableResponse<DevTestLabArtifactResource> response = await collection.GetIfExistsAsync(name);
DevTestLabArtifactResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DevTestLabArtifactData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
