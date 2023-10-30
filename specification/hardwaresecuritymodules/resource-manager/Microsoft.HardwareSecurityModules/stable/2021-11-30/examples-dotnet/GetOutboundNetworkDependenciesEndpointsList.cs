using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HardwareSecurityModules;
using Azure.ResourceManager.HardwareSecurityModules.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/stable/2021-11-30/examples/GetOutboundNetworkDependenciesEndpointsList.json
// this example is just showing the usage of "DedicatedHsm_ListOutboundNetworkDependenciesEndpoints" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DedicatedHsmResource created on azure
// for more information of creating DedicatedHsmResource, please refer to the document of DedicatedHsmResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "hsm-group";
string name = "hsm1";
ResourceIdentifier dedicatedHsmResourceId = DedicatedHsmResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
DedicatedHsmResource dedicatedHsm = client.GetDedicatedHsmResource(dedicatedHsmResourceId);

// invoke the operation and iterate over the result
await foreach (OutboundEnvironmentEndpoint item in dedicatedHsm.GetOutboundNetworkDependenciesEndpointsAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
