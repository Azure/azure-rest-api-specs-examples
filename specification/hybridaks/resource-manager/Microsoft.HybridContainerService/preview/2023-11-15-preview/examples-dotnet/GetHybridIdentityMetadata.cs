using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridContainerService;

// Generated from example definition: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2023-11-15-preview/examples/GetHybridIdentityMetadata.json
// this example is just showing the usage of "HybridIdentityMetadata_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridIdentityMetadataResource created on azure
// for more information of creating HybridIdentityMetadataResource, please refer to the document of HybridIdentityMetadataResource
string connectedClusterResourceUri = "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster";
ResourceIdentifier hybridIdentityMetadataResourceId = HybridIdentityMetadataResource.CreateResourceIdentifier(connectedClusterResourceUri);
HybridIdentityMetadataResource hybridIdentityMetadata = client.GetHybridIdentityMetadataResource(hybridIdentityMetadataResourceId);

// invoke the operation
HybridIdentityMetadataResource result = await hybridIdentityMetadata.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HybridIdentityMetadataData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
