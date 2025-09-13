using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ComputeSchedule.Models;
using Azure.ResourceManager.ComputeSchedule;

// Generated from example definition: 2025-04-15-preview/ScheduledActions_CancelNextOccurrence_MaximumSet_Gen.json
// this example is just showing the usage of "ScheduledActions_CancelNextOccurrence" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ScheduledActionResource created on azure
// for more information of creating ScheduledActionResource, please refer to the document of ScheduledActionResource
string subscriptionId = "CB26D7CB-3E27-465F-99C8-EAF7A4118245";
string resourceGroupName = "rgcomputeschedule";
string scheduledActionName = "myScheduledAction";
ResourceIdentifier scheduledActionResourceId = ScheduledActionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, scheduledActionName);
ScheduledActionResource scheduledAction = client.GetScheduledActionResource(scheduledActionResourceId);

// invoke the operation
OccurrenceCancelContent content = new OccurrenceCancelContent(new ResourceIdentifier[] { new ResourceIdentifier("/subscriptions/1d04e8f1-ee04-4056-b0b2-718f5bb45b04/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachines/myVm") });
ScheduledActionResourceOperationResult result = await scheduledAction.CancelNextOccurrenceAsync(content);

Console.WriteLine($"Succeeded: {result}");
