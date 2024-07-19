using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.FrontDoor.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.FrontDoor;

// Generated from example definition: specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentUpdateProfile.json
// this example is just showing the usage of "NetworkExperimentProfiles_Update" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
FrontDoorNetworkExperimentProfilePatch patch = new FrontDoorNetworkExperimentProfilePatch()
{
    Tags =
    {
    ["key1"] = "value1",
    ["key2"] = "value2",
    },
    EnabledState = FrontDoorExperimentState.Enabled,
};
ArmOperation<FrontDoorNetworkExperimentProfileResource> lro = await frontDoorNetworkExperimentProfile.UpdateAsync(WaitUntil.Completed, patch);
FrontDoorNetworkExperimentProfileResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FrontDoorNetworkExperimentProfileData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
