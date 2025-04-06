using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HybridConnectivity;

// Generated from example definition: 2024-12-01/ServiceConfigurationsList.json
// this example is just showing the usage of "ServiceConfigurationResource_ListByEndpointResource" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridConnectivityEndpointResource created on azure
// for more information of creating HybridConnectivityEndpointResource, please refer to the document of HybridConnectivityEndpointResource
string resourceUri = "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridConnectivity/endpoints/default";
string endpointName = "default";
ResourceIdentifier hybridConnectivityEndpointResourceId = HybridConnectivityEndpointResource.CreateResourceIdentifier(resourceUri, endpointName);
HybridConnectivityEndpointResource hybridConnectivityEndpoint = client.GetHybridConnectivityEndpointResource(hybridConnectivityEndpointResourceId);

// get the collection of this HybridConnectivityServiceConfigurationResource
HybridConnectivityServiceConfigurationCollection collection = hybridConnectivityEndpoint.GetHybridConnectivityServiceConfigurations();

// invoke the operation and iterate over the result
await foreach (HybridConnectivityServiceConfigurationResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    HybridConnectivityServiceConfigurationData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
