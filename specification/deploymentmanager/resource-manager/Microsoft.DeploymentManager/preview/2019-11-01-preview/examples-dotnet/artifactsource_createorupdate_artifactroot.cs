using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DeploymentManager;
using Azure.ResourceManager.DeploymentManager.Models;

// Generated from example definition: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/artifactsource_createorupdate_artifactroot.json
// this example is just showing the usage of "ArtifactSources_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArtifactSourceResource created on azure
// for more information of creating ArtifactSourceResource, please refer to the document of ArtifactSourceResource
string subscriptionId = "caac1590-e859-444f-a9e0-62091c0f5929";
string resourceGroupName = "myResourceGroup";
string artifactSourceName = "myArtifactSource";
ResourceIdentifier artifactSourceResourceId = ArtifactSourceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, artifactSourceName);
ArtifactSourceResource artifactSource = client.GetArtifactSourceResource(artifactSourceResourceId);

// invoke the operation
ArtifactSourceData data = new ArtifactSourceData(new AzureLocation("centralus"))
{
    SourceType = "AzureStorage",
    ArtifactRoot = "1.0.0.0",
    Authentication = new SasAuthentication()
    {
        SasUri = new Uri("https://mystorageaccount.blob.core.windows.net/myartifactsource?st=2018-07-07T14%3A10%3A00Z&se=2019-12-31T15%3A10%3A00Z&sp=rl&sv=2017-04-17&sr=c&sig=Yh2SoJ1NhhLRwCLln7de%2Fkabcdefghijklmno5sWEIk%3D"),
    },
    Tags =
    {
    },
};
ArmOperation<ArtifactSourceResource> lro = await artifactSource.UpdateAsync(WaitUntil.Completed, data);
ArtifactSourceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ArtifactSourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
