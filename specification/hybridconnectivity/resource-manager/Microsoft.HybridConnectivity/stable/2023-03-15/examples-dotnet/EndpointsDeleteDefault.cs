using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HybridConnectivity.Models;
using Azure.ResourceManager.HybridConnectivity;

// Generated from example definition: specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/stable/2023-03-15/examples/EndpointsDeleteDefault.json
// this example is just showing the usage of "Endpoints_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridConnectivityEndpointResource created on azure
// for more information of creating HybridConnectivityEndpointResource, please refer to the document of HybridConnectivityEndpointResource
string resourceUri = "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine";
string endpointName = "default";
ResourceIdentifier hybridConnectivityEndpointResourceId = HybridConnectivityEndpointResource.CreateResourceIdentifier(resourceUri, endpointName);
HybridConnectivityEndpointResource hybridConnectivityEndpoint = client.GetHybridConnectivityEndpointResource(hybridConnectivityEndpointResourceId);

// invoke the operation
await hybridConnectivityEndpoint.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
