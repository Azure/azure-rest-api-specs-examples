using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;
// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-01-03/examples/galleryExamples/GalleryApplicationVersion_Get_WithReplicationStatus.json
// this example is just showing the usage of "GalleryApplicationVersions_Get" operation, for the dependent resources, they will have to be created separately.
            
// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());
            
// this example assumes you already have this GalleryApplicationVersionResource created on azure
// for more information of creating GalleryApplicationVersionResource, please refer to the document of GalleryApplicationVersionResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string galleryName = "myGalleryName";
string galleryApplicationName = "myGalleryApplicationName";
string galleryApplicationVersionName = "1.0.0";
ResourceIdentifier galleryApplicationVersionResourceId = GalleryApplicationVersionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName);
GalleryApplicationVersionResource galleryApplicationVersion = client.GetGalleryApplicationVersionResource(galleryApplicationVersionResourceId);
            
// invoke the operation
ReplicationStatusType? expand = ReplicationStatusType.ReplicationStatus;
GalleryApplicationVersionResource result = await galleryApplicationVersion.GetAsync(expand: expand);
            
// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
GalleryApplicationVersionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
