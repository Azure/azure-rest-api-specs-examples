using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ComputeSchedule.Models;
using Azure.ResourceManager.ComputeSchedule;

// Generated from example definition: 2025-04-15-preview/ScheduledActions_Update_MaximumSet_Gen.json
// this example is just showing the usage of "ScheduledAction_Update" operation, for the dependent resources, they will have to be created separately.

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
ScheduledActionPatch patch = new ScheduledActionPatch
{
    Tags =
    {
    ["key9989"] = "tryjidk"
    },
    Properties = new ScheduledActionPatchProperties
    {
        ResourceType = ScheduledActionResourceType.VirtualMachine,
        ActionType = ScheduledActionType.Start,
        StartOn = DateTimeOffset.Parse("2025-04-17T00:23:58.149Z"),
        EndOn = DateTimeOffset.Parse("2025-04-17T00:23:58.149Z"),
        Schedule = new ScheduledActionsSchedule(XmlConvert.ToTimeSpan("19:00:00"), "bni", new ScheduledActionsScheduleWeekDay[] { ScheduledActionsScheduleWeekDay.Monday }, new ScheduledActionsScheduleMonth[] { ScheduledActionsScheduleMonth.January }, new int[] { 15 })
        {
            ExecutionParameters = new ScheduledActionExecutionParameterDetail
            {
                OptimizationPreference = ScheduledActionOptimizationPreference.Cost,
                RetryPolicy = new UserRequestRetryPolicy
                {
                    RetryCount = 17,
                    RetryWindowInMinutes = 29,
                },
            },
            DeadlineType = ScheduledActionDeadlineType.Unknown,
        },
        NotificationSettings = {new NotificationSettings("wbhryycyolvnypjxzlawwvb", NotificationType.Email, NotificationLanguage.EnUs)
        {
        IsDisabled = true,
        }},
        Disabled = true,
    },
};
ScheduledActionResource result = await scheduledAction.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ScheduledActionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
