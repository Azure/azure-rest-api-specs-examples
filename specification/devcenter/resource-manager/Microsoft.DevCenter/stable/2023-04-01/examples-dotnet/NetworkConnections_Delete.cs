using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevCenter;
using Azure.ResourceManager.DevCenter.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/NetworkConnections_Delete.json
// this example is just showing the usage of "NetworkConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevCenterNetworkConnectionResource created on azure
// for more information of creating DevCenterNetworkConnectionResource, please refer to the document of DevCenterNetworkConnectionResource
string subscriptionId = "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
string resourceGroupName = "rg1";
string networkConnectionName = "eastusnetwork";
ResourceIdentifier devCenterNetworkConnectionResourceId = DevCenterNetworkConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkConnectionName);
DevCenterNetworkConnectionResource devCenterNetworkConnection = client.GetDevCenterNetworkConnectionResource(devCenterNetworkConnectionResourceId);

// invoke the operation
await devCenterNetworkConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
