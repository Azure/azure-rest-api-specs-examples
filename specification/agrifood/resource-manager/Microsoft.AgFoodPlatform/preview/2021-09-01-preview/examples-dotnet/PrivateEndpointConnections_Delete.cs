using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AgFoodPlatform;
using Azure.ResourceManager.AgFoodPlatform.Models;

// Generated from example definition: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/PrivateEndpointConnections_Delete.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AgFoodPlatformPrivateEndpointConnectionResource created on azure
// for more information of creating AgFoodPlatformPrivateEndpointConnectionResource, please refer to the document of AgFoodPlatformPrivateEndpointConnectionResource
string subscriptionId = "11111111-2222-3333-4444-555555555555";
string resourceGroupName = "examples-rg";
string farmBeatsResourceName = "examples-farmbeatsResourceName";
string privateEndpointConnectionName = "privateEndpointConnectionName";
ResourceIdentifier agFoodPlatformPrivateEndpointConnectionResourceId = AgFoodPlatformPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, farmBeatsResourceName, privateEndpointConnectionName);
AgFoodPlatformPrivateEndpointConnectionResource agFoodPlatformPrivateEndpointConnection = client.GetAgFoodPlatformPrivateEndpointConnectionResource(agFoodPlatformPrivateEndpointConnectionResourceId);

// invoke the operation
await agFoodPlatformPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
