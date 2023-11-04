using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridNetwork;
using Azure.ResourceManager.HybridNetwork.Models;

// Generated from example definition: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/ArtifactManifestCreate.json
// this example is just showing the usage of "ArtifactManifests_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArtifactStoreResource created on azure
// for more information of creating ArtifactStoreResource, please refer to the document of ArtifactStoreResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string publisherName = "TestPublisher";
string artifactStoreName = "TestArtifactStore";
ResourceIdentifier artifactStoreResourceId = ArtifactStoreResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, publisherName, artifactStoreName);
ArtifactStoreResource artifactStore = client.GetArtifactStoreResource(artifactStoreResourceId);

// get the collection of this ArtifactManifestResource
ArtifactManifestCollection collection = artifactStore.GetArtifactManifests();

// invoke the operation
string artifactManifestName = "TestManifest";
ArtifactManifestData data = new ArtifactManifestData(new AzureLocation("eastus"))
{
    Properties = new ArtifactManifestPropertiesFormat()
    {
        Artifacts =
        {
        new ManifestArtifactFormat()
        {
        ArtifactName = "fed-rbac",
        ArtifactType = ArtifactType.OCIArtifact,
        ArtifactVersion = "1.0.0",
        },new ManifestArtifactFormat()
        {
        ArtifactName = "nginx",
        ArtifactType = ArtifactType.OCIArtifact,
        ArtifactVersion = "v1",
        }
        },
    },
};
ArmOperation<ArtifactManifestResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, artifactManifestName, data);
ArtifactManifestResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ArtifactManifestData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
