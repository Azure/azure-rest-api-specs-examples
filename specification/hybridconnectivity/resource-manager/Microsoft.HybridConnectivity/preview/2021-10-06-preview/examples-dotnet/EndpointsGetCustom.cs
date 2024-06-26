using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridConnectivity;
using Azure.ResourceManager.HybridConnectivity.Models;

// Generated from example definition: specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/preview/2021-10-06-preview/examples/EndpointsGetCustom.json
// this example is just showing the usage of "Endpoints_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this EndpointResource
string scope = "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", scope));
EndpointResourceCollection collection = client.GetEndpointResources(scopeId);

// invoke the operation
string endpointName = "custom";
NullableResponse<EndpointResource> response = await collection.GetIfExistsAsync(endpointName);
EndpointResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    EndpointResourceData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
