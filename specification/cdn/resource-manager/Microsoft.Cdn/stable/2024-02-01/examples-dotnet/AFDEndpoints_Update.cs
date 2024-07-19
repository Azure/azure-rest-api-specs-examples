using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Cdn.Models;
using Azure.ResourceManager.Cdn;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/AFDEndpoints_Update.json
// this example is just showing the usage of "FrontDoorEndpoints_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FrontDoorEndpointResource created on azure
// for more information of creating FrontDoorEndpointResource, please refer to the document of FrontDoorEndpointResource
string subscriptionId = "subid";
string resourceGroupName = "RG";
string profileName = "profile1";
string endpointName = "endpoint1";
ResourceIdentifier frontDoorEndpointResourceId = FrontDoorEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName, endpointName);
FrontDoorEndpointResource frontDoorEndpoint = client.GetFrontDoorEndpointResource(frontDoorEndpointResourceId);

// invoke the operation
FrontDoorEndpointPatch patch = new FrontDoorEndpointPatch()
{
    Tags =
    {
    },
    EnabledState = EnabledState.Enabled,
};
ArmOperation<FrontDoorEndpointResource> lro = await frontDoorEndpoint.UpdateAsync(WaitUntil.Completed, patch);
FrontDoorEndpointResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FrontDoorEndpointData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
