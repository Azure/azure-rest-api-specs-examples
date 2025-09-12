using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ComputeSchedule.Models;
using Azure.ResourceManager.ComputeSchedule;

// Generated from example definition: 2025-04-15-preview/Occurrences_Delay_MaximumSet_Gen.json
// this example is just showing the usage of "Occurrences_Delay" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ScheduledActionOccurrenceResource created on azure
// for more information of creating ScheduledActionOccurrenceResource, please refer to the document of ScheduledActionOccurrenceResource
string subscriptionId = "CB26D7CB-3E27-465F-99C8-EAF7A4118245";
string resourceGroupName = "rgcomputeschedule";
string scheduledActionName = "myScheduledAction";
string occurrenceId = "CB26D7CB-3E27-465F-99C8-EAF7A4118245";
ResourceIdentifier scheduledActionOccurrenceResourceId = ScheduledActionOccurrenceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, scheduledActionName, occurrenceId);
ScheduledActionOccurrenceResource scheduledActionOccurrence = client.GetScheduledActionOccurrenceResource(scheduledActionOccurrenceResourceId);

// invoke the operation
OccurrenceDelayContent content = new OccurrenceDelayContent(DateTimeOffset.Parse("2025-05-22T17:00:00.000-07:00"), new ResourceIdentifier[] { new ResourceIdentifier("/subscriptions/CB26D7CB-3E27-465F-99C8-EAF7A4118245/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachines/myVm") });
ArmOperation<ScheduledActionResourceOperationResult> lro = await scheduledActionOccurrence.DelayAsync(WaitUntil.Completed, content);
ScheduledActionResourceOperationResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
