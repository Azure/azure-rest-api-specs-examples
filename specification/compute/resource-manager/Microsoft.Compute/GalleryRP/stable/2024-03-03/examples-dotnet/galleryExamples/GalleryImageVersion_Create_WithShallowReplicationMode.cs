using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2024-03-03/examples/galleryExamples/GalleryImageVersion_Create_WithShallowReplicationMode.json
// this example is just showing the usage of "GalleryImageVersions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GalleryImageResource created on azure
// for more information of creating GalleryImageResource, please refer to the document of GalleryImageResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string galleryName = "myGalleryName";
string galleryImageName = "myGalleryImageName";
ResourceIdentifier galleryImageResourceId = GalleryImageResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, galleryName, galleryImageName);
GalleryImageResource galleryImage = client.GetGalleryImageResource(galleryImageResourceId);

// get the collection of this GalleryImageVersionResource
GalleryImageVersionCollection collection = galleryImage.GetGalleryImageVersions();

// invoke the operation
string galleryImageVersionName = "1.0.0";
GalleryImageVersionData data = new GalleryImageVersionData(new AzureLocation("West US"))
{
    PublishingProfile = new GalleryImageVersionPublishingProfile
    {
        TargetRegions = {new TargetRegion("West US")
        {
        RegionalReplicaCount = 1,
        IsExcludedFromLatest = false,
        }},
        ReplicationMode = GalleryReplicationMode.Shallow,
    },
    StorageProfile = new GalleryImageVersionStorageProfile
    {
        GallerySource = new GalleryArtifactVersionFullSource
        {
            Id = new ResourceIdentifier("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/images/{imageName}"),
        },
    },
    SafetyProfile = new GalleryImageVersionSafetyProfile
    {
        IsBlockedDeletionBeforeEndOfLife = false,
        AllowDeletionOfReplicatedLocations = false,
    },
};
ArmOperation<GalleryImageVersionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, galleryImageVersionName, data);
GalleryImageVersionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
GalleryImageVersionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
