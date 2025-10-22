using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.TrafficManager.Models;
using Azure.ResourceManager.TrafficManager;

// Generated from example definition: specification/trafficmanager/resource-manager/Microsoft.Network/preview/2024-04-01-preview/examples/NameAvailabilityTest_NameNotAvailable-POST-example-23.json
// this example is just showing the usage of "Profiles_CheckTrafficManagerRelativeDnsNameAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
TrafficManagerRelativeDnsNameAvailabilityContent content = new TrafficManagerRelativeDnsNameAvailabilityContent
{
    Name = "azsmnet4696",
    ResourceType = new ResourceType("microsoft.network/trafficmanagerprofiles"),
};
TrafficManagerNameAvailabilityResult result = await tenantResource.CheckTrafficManagerRelativeDnsNameAvailabilityAsync(content);

Console.WriteLine($"Succeeded: {result}");
