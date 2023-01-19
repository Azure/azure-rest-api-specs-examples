using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Media;
using Azure.ResourceManager.Media.Models;

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/accountFilters-get-by-name.json
// this example is just showing the usage of "AccountFilters_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MediaServicesAccountFilterResource created on azure
// for more information of creating MediaServicesAccountFilterResource, please refer to the document of MediaServicesAccountFilterResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contoso";
string accountName = "contosomedia";
string filterName = "accountFilterWithTrack";
ResourceIdentifier mediaServicesAccountFilterResourceId = MediaServicesAccountFilterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, filterName);
MediaServicesAccountFilterResource mediaServicesAccountFilter = client.GetMediaServicesAccountFilterResource(mediaServicesAccountFilterResourceId);

// invoke the operation
MediaServicesAccountFilterResource result = await mediaServicesAccountFilter.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MediaServicesAccountFilterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
