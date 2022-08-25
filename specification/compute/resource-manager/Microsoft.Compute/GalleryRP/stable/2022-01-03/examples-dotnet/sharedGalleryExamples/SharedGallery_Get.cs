using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;
// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-01-03/examples/sharedGalleryExamples/SharedGallery_Get.json
// this example is just showing the usage of "SharedGalleries_Get" operation, for the dependent resources, they will have to be created separately.
            
// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());
            
// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "{subscription-id}";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);
            
// get the collection of this SharedGalleryResource
AzureLocation location = new AzureLocation("myLocation");
SharedGalleryCollection collection = subscriptionResource.GetSharedGalleries(location);
            
// invoke the operation
string galleryUniqueName = "galleryUniqueName";
bool result = await collection.ExistsAsync(galleryUniqueName);
            
Console.WriteLine($"Succeeded: {result}");
