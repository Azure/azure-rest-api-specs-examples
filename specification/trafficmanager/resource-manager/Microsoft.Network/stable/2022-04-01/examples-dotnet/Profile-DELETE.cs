using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.TrafficManager;
using Azure.ResourceManager.TrafficManager.Models;

// Generated from example definition: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Profile-DELETE.json
// this example is just showing the usage of "Profiles_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TrafficManagerProfileResource created on azure
// for more information of creating TrafficManagerProfileResource, please refer to the document of TrafficManagerProfileResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "azuresdkfornetautoresttrafficmanager1323";
string profileName = "azuresdkfornetautoresttrafficmanager3880";
ResourceIdentifier trafficManagerProfileResourceId = TrafficManagerProfileResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName);
TrafficManagerProfileResource trafficManagerProfile = client.GetTrafficManagerProfileResource(trafficManagerProfileResourceId);

// invoke the operation
await trafficManagerProfile.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
