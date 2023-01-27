using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Media;
using Azure.ResourceManager.Media.Models;

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/Encoding/stable/2022-07-01/examples/transforms-delete.json
// this example is just showing the usage of "Transforms_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MediaTransformResource created on azure
// for more information of creating MediaTransformResource, please refer to the document of MediaTransformResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contosoresources";
string accountName = "contosomedia";
string transformName = "sampleTransform";
ResourceIdentifier mediaTransformResourceId = MediaTransformResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, transformName);
MediaTransformResource mediaTransform = client.GetMediaTransformResource(mediaTransformResourceId);

// invoke the operation
await mediaTransform.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
