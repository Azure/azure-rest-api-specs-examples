using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Media.Models;
using Azure.ResourceManager.Media;

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/Encoding/stable/2022-07-01/examples/jobs-list-all-filter-by-name-and-state.json
// this example is just showing the usage of "Jobs_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MediaTransformResource created on azure
// for more information of creating MediaTransformResource, please refer to the document of MediaTransformResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contosoresources";
string accountName = "contosomedia";
string transformName = "exampleTransform";
ResourceIdentifier mediaTransformResourceId = MediaTransformResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, transformName);
MediaTransformResource mediaTransform = client.GetMediaTransformResource(mediaTransformResourceId);

// get the collection of this MediaJobResource
MediaJobCollection collection = mediaTransform.GetMediaJobs();

// invoke the operation and iterate over the result
string filter = "name eq 'job3' and properties/state eq Microsoft.Media.JobState'finished'";
await foreach (MediaJobResource item in collection.GetAllAsync(filter: filter))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MediaJobData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
