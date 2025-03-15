using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SelfHelp.Models;
using Azure.ResourceManager.SelfHelp;

// Generated from example definition: specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/CheckNameAvailabilityForDiagnosticWhenNameIsNotAvailable.json
// this example is just showing the usage of "CheckNameAvailability_CheckAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// invoke the operation
ResourceIdentifier scope = new ResourceIdentifier("subscriptions/0d0fcd2e-c4fd-4349-8497-200edb3923c6");
SelfHelpNameAvailabilityContent content = new SelfHelpNameAvailabilityContent
{
    ResourceName = "sampleName",
    ResourceType = new ResourceType("Microsoft.Help/diagnostics"),
};
SelfHelpNameAvailabilityResult result = await client.CheckSelfHelpNameAvailabilityAsync(scope, content: content);

Console.WriteLine($"Succeeded: {result}");
