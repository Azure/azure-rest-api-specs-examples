using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Media;
using Azure.ResourceManager.Media.Models;

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2023-01-01/examples/accountFilters-delete.json
// this example is just showing the usage of "AccountFilters_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MediaServicesAccountFilterResource created on azure
// for more information of creating MediaServicesAccountFilterResource, please refer to the document of MediaServicesAccountFilterResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contosorg";
string accountName = "contosomedia";
string filterName = "accountFilterWithTimeWindowAndTrack";
ResourceIdentifier mediaServicesAccountFilterResourceId = MediaServicesAccountFilterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, filterName);
MediaServicesAccountFilterResource mediaServicesAccountFilter = client.GetMediaServicesAccountFilterResource(mediaServicesAccountFilterResourceId);

// invoke the operation
await mediaServicesAccountFilter.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
