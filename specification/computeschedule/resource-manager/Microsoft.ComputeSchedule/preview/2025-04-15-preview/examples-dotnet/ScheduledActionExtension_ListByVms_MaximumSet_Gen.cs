using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ComputeSchedule.Models;
using Azure.ResourceManager.ComputeSchedule;

// Generated from example definition: 2025-04-15-preview/ScheduledActionExtension_ListByVms_MaximumSet_Gen.json
// this example is just showing the usage of "ScheduledActionResources_GetAssociatedScheduledActions" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// invoke the operation and iterate over the result
ResourceIdentifier scope = null;
await foreach (ScheduledActionResources item in client.GetAssociatedScheduledActionsAsync(scope))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
