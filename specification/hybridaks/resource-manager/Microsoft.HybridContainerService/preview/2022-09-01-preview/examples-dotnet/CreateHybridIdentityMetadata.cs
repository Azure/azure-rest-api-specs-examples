using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridContainerService;

// Generated from example definition: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/CreateHybridIdentityMetadata.json
// this example is just showing the usage of "HybridIdentityMetadata_Put" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridIdentityMetadataResource created on azure
// for more information of creating HybridIdentityMetadataResource, please refer to the document of HybridIdentityMetadataResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "testrg";
string resourceName = "ContosoTargetCluster";
string hybridIdentityMetadataResourceName = "default";
ResourceIdentifier hybridIdentityMetadataResourceId = HybridIdentityMetadataResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, hybridIdentityMetadataResourceName);
HybridIdentityMetadataResource hybridIdentityMetadata = client.GetHybridIdentityMetadataResource(hybridIdentityMetadataResourceId);

// invoke the operation
HybridIdentityMetadataData data = new HybridIdentityMetadataData()
{
    ResourceUid = "f8b82dff-38ef-4220-99ef-d3a3f86ddc6c",
    PublicKey = "8ec7d60c-9700-40b1-8e6e-e5b2f6f477f2",
};
ArmOperation<HybridIdentityMetadataResource> lro = await hybridIdentityMetadata.UpdateAsync(WaitUntil.Completed, data);
HybridIdentityMetadataResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HybridIdentityMetadataData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
