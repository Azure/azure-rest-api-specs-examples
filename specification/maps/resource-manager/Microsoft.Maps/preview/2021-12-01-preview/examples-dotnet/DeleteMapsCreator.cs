using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Maps;
using Azure.ResourceManager.Maps.Models;

// Generated from example definition: specification/maps/resource-manager/Microsoft.Maps/preview/2021-12-01-preview/examples/DeleteMapsCreator.json
// this example is just showing the usage of "Creators_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MapsCreatorResource created on azure
// for more information of creating MapsCreatorResource, please refer to the document of MapsCreatorResource
string subscriptionId = "21a9967a-e8a9-4656-a70b-96ff1c4d05a0";
string resourceGroupName = "myResourceGroup";
string accountName = "myMapsAccount";
string creatorName = "myCreator";
ResourceIdentifier mapsCreatorResourceId = MapsCreatorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, creatorName);
MapsCreatorResource mapsCreator = client.GetMapsCreatorResource(mapsCreatorResourceId);

// invoke the operation
await mapsCreator.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
