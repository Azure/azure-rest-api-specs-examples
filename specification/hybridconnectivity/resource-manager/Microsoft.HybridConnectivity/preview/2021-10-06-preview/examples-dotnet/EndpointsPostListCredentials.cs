using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridConnectivity;
using Azure.ResourceManager.HybridConnectivity.Models;

// Generated from example definition: specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/preview/2021-10-06-preview/examples/EndpointsPostListCredentials.json
// this example is just showing the usage of "Endpoints_ListCredentials" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EndpointResource created on azure
// for more information of creating EndpointResource, please refer to the document of EndpointResource
string scope = "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine";
string endpointName = "default";
ResourceIdentifier endpointResourceId = EndpointResource.CreateResourceIdentifier(scope, endpointName);
EndpointResource endpointResource = client.GetEndpointResource(endpointResourceId);

// invoke the operation
long? expiresin = 10800;
TargetResourceEndpointAccess result = await endpointResource.GetCredentialsAsync(expiresin: expiresin);

Console.WriteLine($"Succeeded: {result}");
