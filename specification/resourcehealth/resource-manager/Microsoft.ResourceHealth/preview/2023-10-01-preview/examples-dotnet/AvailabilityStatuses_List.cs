using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ResourceHealth.Models;
using Azure.ResourceManager.ResourceHealth;

// Generated from example definition: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/AvailabilityStatuses_List.json
// this example is just showing the usage of "AvailabilityStatuses_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// invoke the operation and iterate over the result
ResourceIdentifier scope = null;
await foreach (ResourceHealthAvailabilityStatus item in client.GetAvailabilityStatusesAsync(scope))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
