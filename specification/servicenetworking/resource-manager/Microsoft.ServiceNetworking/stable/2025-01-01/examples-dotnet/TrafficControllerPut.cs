using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ServiceNetworking;

// Generated from example definition: 2025-01-01/TrafficControllerPut.json
// this example is just showing the usage of "TrafficController_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this TrafficControllerResource
TrafficControllerCollection collection = resourceGroupResource.GetTrafficControllers();

// invoke the operation
string trafficControllerName = "tc1";
TrafficControllerData data = new TrafficControllerData(new AzureLocation("NorthCentralUS"))
{
    Tags =
    {
    ["key1"] = "value1"
    },
};
ArmOperation<TrafficControllerResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, trafficControllerName, data);
TrafficControllerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
TrafficControllerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
