using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ResourceHealth.Models;
using Azure.ResourceManager.ResourceHealth;

// Generated from example definition: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/ChildAvailabilityStatus_GetByResource.json
// this example is just showing the usage of "ChildAvailabilityStatuses_GetByResource" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// invoke the operation
ResourceIdentifier scope = null;
string expand = "recommendedactions";
ResourceHealthAvailabilityStatus result = await client.GetAvailabilityStatusOfChildResourceAsync(scope, expand: expand);

Console.WriteLine($"Succeeded: {result}");
