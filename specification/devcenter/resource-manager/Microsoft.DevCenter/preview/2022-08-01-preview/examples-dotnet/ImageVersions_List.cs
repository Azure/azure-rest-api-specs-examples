using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevCenter;

// Generated from example definition: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/ImageVersions_List.json
// this example is just showing the usage of "ImageVersions_ListByImage" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ImageResource created on azure
// for more information of creating ImageResource, please refer to the document of ImageResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "rg1";
string devCenterName = "Contoso";
string galleryName = "DefaultDevGallery";
string imageName = "Win11";
ResourceIdentifier imageResourceId = ImageResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, devCenterName, galleryName, imageName);
ImageResource image = client.GetImageResource(imageResourceId);

// get the collection of this ImageVersionResource
ImageVersionCollection collection = image.GetImageVersions();

// invoke the operation and iterate over the result
await foreach (ImageVersionResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ImageVersionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
