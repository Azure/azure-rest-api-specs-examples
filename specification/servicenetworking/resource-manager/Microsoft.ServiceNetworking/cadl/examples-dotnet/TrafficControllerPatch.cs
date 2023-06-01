using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ServiceNetworking;
using Azure.ResourceManager.ServiceNetworking.Models;

// Generated from example definition: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/TrafficControllerPatch.json
// this example is just showing the usage of "TrafficControllerInterface_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TrafficControllerResource created on azure
// for more information of creating TrafficControllerResource, please refer to the document of TrafficControllerResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string trafficControllerName = "tc1";
ResourceIdentifier trafficControllerResourceId = TrafficControllerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, trafficControllerName);
TrafficControllerResource trafficController = client.GetTrafficControllerResource(trafficControllerResourceId);

// invoke the operation
TrafficControllerPatch patch = new TrafficControllerPatch()
{
    Tags =
    {
    ["key1"] = "value1",
    },
};
TrafficControllerResource result = await trafficController.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
TrafficControllerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
