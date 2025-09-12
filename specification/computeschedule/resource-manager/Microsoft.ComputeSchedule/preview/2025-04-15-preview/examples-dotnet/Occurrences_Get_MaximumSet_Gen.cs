using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ComputeSchedule.Models;
using Azure.ResourceManager.ComputeSchedule;

// Generated from example definition: 2025-04-15-preview/Occurrences_Get_MaximumSet_Gen.json
// this example is just showing the usage of "Occurrence_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ScheduledActionOccurrenceResource created on azure
// for more information of creating ScheduledActionOccurrenceResource, please refer to the document of ScheduledActionOccurrenceResource
string subscriptionId = "CB26D7CB-3E27-465F-99C8-EAF7A4118245";
string resourceGroupName = "rgcomputeschedule";
string scheduledActionName = "myScheduledAction";
string occurrenceId = "67b5bada-4772-43fc-8dbb-402476d98a45";
ResourceIdentifier scheduledActionOccurrenceResourceId = ScheduledActionOccurrenceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, scheduledActionName, occurrenceId);
ScheduledActionOccurrenceResource scheduledActionOccurrence = client.GetScheduledActionOccurrenceResource(scheduledActionOccurrenceResourceId);

// invoke the operation
ScheduledActionOccurrenceResource result = await scheduledActionOccurrence.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ScheduledActionOccurrenceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
