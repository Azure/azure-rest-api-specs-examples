using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridNetwork;
using Azure.ResourceManager.HybridNetwork.Models;

// Generated from example definition: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/PureProxyArtifact/ArtifactChangeState.json
// this example is just showing the usage of "ProxyArtifact_UpdateState" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArtifactStoreResource created on azure
// for more information of creating ArtifactStoreResource, please refer to the document of ArtifactStoreResource
string subscriptionId = "subid";
string resourceGroupName = "TestResourceGroup";
string publisherName = "TestPublisher";
string artifactStoreName = "TestArtifactStoreName";
ResourceIdentifier artifactStoreResourceId = ArtifactStoreResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, publisherName, artifactStoreName);
ArtifactStoreResource artifactStore = client.GetArtifactStoreResource(artifactStoreResourceId);

// invoke the operation
string artifactVersionName = "1.0.0";
string artifactName = "fedrbac";
ArtifactChangeState artifactChangeState = new ArtifactChangeState()
{
    ArtifactState = ArtifactState.Deprecated,
};
ArmOperation<ProxyArtifactVersionsListOverview> lro = await artifactStore.UpdateStateProxyArtifactAsync(WaitUntil.Completed, artifactVersionName, artifactName, artifactChangeState);
ProxyArtifactVersionsListOverview result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
