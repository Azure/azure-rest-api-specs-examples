using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;
// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-01-03/examples/galleryExamples/GalleryApplicationVersion_Create.json
// this example is just showing the usage of "GalleryApplicationVersions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.
            
// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());
            
// this example assumes you already have this GalleryApplicationResource created on azure
// for more information of creating GalleryApplicationResource, please refer to the document of GalleryApplicationResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string galleryName = "myGalleryName";
string galleryApplicationName = "myGalleryApplicationName";
ResourceIdentifier galleryApplicationResourceId = GalleryApplicationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, galleryName, galleryApplicationName);
GalleryApplicationResource galleryApplication = client.GetGalleryApplicationResource(galleryApplicationResourceId);
            
// get the collection of this GalleryApplicationVersionResource
GalleryApplicationVersionCollection collection = galleryApplication.GetGalleryApplicationVersions();
            
// invoke the operation
string galleryApplicationVersionName = "1.0.0";
GalleryApplicationVersionData data = new GalleryApplicationVersionData(new AzureLocation("West US"))
{
    PublishingProfile = new GalleryApplicationVersionPublishingProfile(new UserArtifactSource("https://mystorageaccount.blob.core.windows.net/mycontainer/package.zip?{sasKey}"))
    {
        ManageActions = new UserArtifactManagement("powershell -command \"Expand-Archive -Path package.zip -DestinationPath C:\\package\"", "del C:\\package "),
        TargetRegions =
                    {
                    new TargetRegion("West US")
                    {
                    RegionalReplicaCount = 1,
                    StorageAccountType = ImageStorageAccountType.StandardLrs,
                    }
                    },
        ReplicaCount = 1,
        EndOfLifeOn = DateTimeOffset.Parse("2019-07-01T07:00:00Z"),
        StorageAccountType = ImageStorageAccountType.StandardLrs,
    },
};
ArmOperation<GalleryApplicationVersionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, galleryApplicationVersionName, data);
GalleryApplicationVersionResource result = lro.Value;
            
// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
GalleryApplicationVersionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
