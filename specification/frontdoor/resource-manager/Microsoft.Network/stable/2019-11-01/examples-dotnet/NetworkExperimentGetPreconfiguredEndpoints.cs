using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.FrontDoor.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.FrontDoor;

// Generated from example definition: specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentGetPreconfiguredEndpoints.json
// this example is just showing the usage of "PreconfiguredEndpoints_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FrontDoorNetworkExperimentProfileResource created on azure
// for more information of creating FrontDoorNetworkExperimentProfileResource, please refer to the document of FrontDoorNetworkExperimentProfileResource
string subscriptionId = "subid";
string resourceGroupName = "MyResourceGroup";
string profileName = "MyProfile";
ResourceIdentifier frontDoorNetworkExperimentProfileResourceId = FrontDoorNetworkExperimentProfileResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName);
FrontDoorNetworkExperimentProfileResource frontDoorNetworkExperimentProfile = client.GetFrontDoorNetworkExperimentProfileResource(frontDoorNetworkExperimentProfileResourceId);

// invoke the operation and iterate over the result
await foreach (PreconfiguredEndpoint item in frontDoorNetworkExperimentProfile.GetPreconfiguredEndpointsAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
