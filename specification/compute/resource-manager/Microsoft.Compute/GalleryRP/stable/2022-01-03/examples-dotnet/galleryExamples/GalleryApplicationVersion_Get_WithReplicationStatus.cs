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
ReplicationStatusType? expand = ReplicationStatusType.ReplicationStatus;
bool result = await collection.ExistsAsync(galleryApplicationVersionName, expand: expand);

Console.WriteLine($"Succeeded: {result}");
