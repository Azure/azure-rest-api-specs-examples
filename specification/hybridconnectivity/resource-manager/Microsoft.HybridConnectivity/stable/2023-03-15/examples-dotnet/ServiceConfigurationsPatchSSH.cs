using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HybridConnectivity.Models;
using Azure.ResourceManager.HybridConnectivity;

// Generated from example definition: specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/stable/2023-03-15/examples/ServiceConfigurationsPatchSSH.json
// this example is just showing the usage of "ServiceConfigurations_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridConnectivityServiceConfigurationResource created on azure
// for more information of creating HybridConnectivityServiceConfigurationResource, please refer to the document of HybridConnectivityServiceConfigurationResource
string resourceUri = "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridConnectivity/endpoints/default";
string endpointName = "default";
string serviceConfigurationName = "SSH";
ResourceIdentifier hybridConnectivityServiceConfigurationResourceId = HybridConnectivityServiceConfigurationResource.CreateResourceIdentifier(resourceUri, endpointName, serviceConfigurationName);
HybridConnectivityServiceConfigurationResource hybridConnectivityServiceConfiguration = client.GetHybridConnectivityServiceConfigurationResource(hybridConnectivityServiceConfigurationResourceId);

// invoke the operation
HybridConnectivityServiceConfigurationPatch patch = new HybridConnectivityServiceConfigurationPatch()
{
    Port = 22L,
};
HybridConnectivityServiceConfigurationResource result = await hybridConnectivityServiceConfiguration.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HybridConnectivityServiceConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
