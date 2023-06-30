using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.TrafficManager;
using Azure.ResourceManager.TrafficManager.Models;

// Generated from example definition: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/NameAvailabilityTest_NameAvailable-POST-example-21.json
// this example is just showing the usage of "Profiles_CheckTrafficManagerRelativeDnsNameAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
TrafficManagerRelativeDnsNameAvailabilityContent content = new TrafficManagerRelativeDnsNameAvailabilityContent()
{
    Name = "azsmnet5403",
    ResourceType = new ResourceType("microsoft.network/trafficmanagerprofiles"),
};
TrafficManagerNameAvailabilityResult result = await tenantResource.CheckTrafficManagerRelativeDnsNameAvailabilityAsync(content);

Console.WriteLine($"Succeeded: {result}");
