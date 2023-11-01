using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ResourceHealth;
using Azure.ResourceManager.ResourceHealth.Models;

// Generated from example definition: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/ChildResources_List.json
// this example is just showing the usage of "ChildResources_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// invoke the operation and iterate over the result
string resourceUri = "resourceUri";
ResourceIdentifier scope = new ResourceIdentifier(string.Format("/{0}", resourceUri));
await foreach (ResourceHealthAvailabilityStatus item in client.GetAvailabilityStatusOfChildResourcesAsync(scope))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
