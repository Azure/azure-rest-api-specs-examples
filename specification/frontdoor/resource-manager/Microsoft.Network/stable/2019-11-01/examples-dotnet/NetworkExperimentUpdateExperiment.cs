using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.FrontDoor.Models;
using Azure.ResourceManager.FrontDoor;

// Generated from example definition: specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentUpdateExperiment.json
// this example is just showing the usage of "Experiments_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FrontDoorExperimentResource created on azure
// for more information of creating FrontDoorExperimentResource, please refer to the document of FrontDoorExperimentResource
string subscriptionId = "subid";
string resourceGroupName = "MyResourceGroup";
string profileName = "MyProfile";
string experimentName = "MyExperiment";
ResourceIdentifier frontDoorExperimentResourceId = FrontDoorExperimentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName, experimentName);
FrontDoorExperimentResource frontDoorExperiment = client.GetFrontDoorExperimentResource(frontDoorExperimentResourceId);

// invoke the operation
FrontDoorExperimentPatch patch = new FrontDoorExperimentPatch()
{
    Description = "string",
    EnabledState = FrontDoorExperimentState.Enabled,
};
ArmOperation<FrontDoorExperimentResource> lro = await frontDoorExperiment.UpdateAsync(WaitUntil.Completed, patch);
FrontDoorExperimentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FrontDoorExperimentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
