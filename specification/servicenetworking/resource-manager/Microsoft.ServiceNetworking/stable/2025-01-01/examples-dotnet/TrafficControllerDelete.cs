using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceNetworking.Models;
using Azure.ResourceManager.ServiceNetworking;

// Generated from example definition: 2025-01-01/TrafficControllerDelete.json
// this example is just showing the usage of "TrafficController_Delete" operation, for the dependent resources, they will have to be created separately.

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
await trafficController.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
