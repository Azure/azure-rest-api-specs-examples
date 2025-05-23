using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.FrontDoor.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.FrontDoor;

// Generated from example definition: specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/CheckFrontdoorNameAvailability.json
// this example is just showing the usage of "FrontDoorNameAvailability_Check" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
FrontDoorNameAvailabilityContent content = new FrontDoorNameAvailabilityContent("sampleName", FrontDoorResourceType.MicrosoftNetworkFrontDoors);
FrontDoorNameAvailabilityResult result = await tenantResource.CheckFrontDoorNameAvailabilityAsync(content);

Console.WriteLine($"Succeeded: {result}");
