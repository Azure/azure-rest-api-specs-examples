using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Media;
using Azure.ResourceManager.Media.Models;

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/Encoding/stable/2022-07-01/examples/jobs-delete.json
// this example is just showing the usage of "Jobs_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MediaJobResource created on azure
// for more information of creating MediaJobResource, please refer to the document of MediaJobResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contosoresources";
string accountName = "contosomedia";
string transformName = "exampleTransform";
string jobName = "jobToDelete";
ResourceIdentifier mediaJobResourceId = MediaJobResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, transformName, jobName);
MediaJobResource mediaJob = client.GetMediaJobResource(mediaJobResourceId);

// invoke the operation
await mediaJob.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
