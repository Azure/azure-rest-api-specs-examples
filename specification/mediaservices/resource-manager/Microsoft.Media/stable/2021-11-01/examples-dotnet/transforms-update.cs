using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Media;
using Azure.ResourceManager.Media.Models;

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/transforms-update.json
// this example is just showing the usage of "Transforms_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MediaTransformResource created on azure
// for more information of creating MediaTransformResource, please refer to the document of MediaTransformResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contosoresources";
string accountName = "contosomedia";
string transformName = "transformToUpdate";
ResourceIdentifier mediaTransformResourceId = MediaTransformResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, transformName);
MediaTransformResource mediaTransform = client.GetMediaTransformResource(mediaTransformResourceId);

// invoke the operation
MediaTransformData data = new MediaTransformData()
{
    Description = "Example transform to illustrate update.",
    Outputs =
    {
    new MediaTransformOutput(new BuiltInStandardEncoderPreset(EncoderNamedPreset.H264MultipleBitrate720P))
    {
    RelativePriority = MediaJobPriority.High,
    }
    },
};
MediaTransformResource result = await mediaTransform.UpdateAsync(data);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MediaTransformData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
