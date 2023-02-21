using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.TrafficManager;
using Azure.ResourceManager.TrafficManager.Models;

// Generated from example definition: specification/trafficmanager/resource-manager/Microsoft.Network/preview/2022-04-01-preview/examples/Profile-GET-WithEndpoints.json
// this example is just showing the usage of "Profiles_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "azuresdkfornetautoresttrafficmanager1323";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this TrafficManagerProfileResource
TrafficManagerProfileCollection collection = resourceGroupResource.GetTrafficManagerProfiles();

// invoke the operation
string profileName = "azuresdkfornetautoresttrafficmanager3880";
bool result = await collection.ExistsAsync(profileName);

Console.WriteLine($"Succeeded: {result}");
