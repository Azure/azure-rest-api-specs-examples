using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DurableTask.Models;
using Azure.ResourceManager.DurableTask;

// Generated from example definition: 2026-02-01/PrivateEndpointConnections_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "Schedulers_DeletePrivateEndpointConnection" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DurableTaskPrivateEndpointConnectionResource created on azure
// for more information of creating DurableTaskPrivateEndpointConnectionResource, please refer to the document of DurableTaskPrivateEndpointConnectionResource
string subscriptionId = "851A7597-D699-45CC-899B-7487A5B3B775";
string resourceGroupName = "rgdurabletask";
string schedulerName = "testscheduler";
string privateEndpointConnectionName = "spzckqrbhfnabu";
ResourceIdentifier DurableTaskPrivateEndpointConnectionResourceId = DurableTaskPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, schedulerName, privateEndpointConnectionName);
DurableTaskPrivateEndpointConnectionResource privateEndpointConnection = client.GetDurableTaskPrivateEndpointConnectionResource(DurableTaskPrivateEndpointConnectionResourceId);

// invoke the operation
await privateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
