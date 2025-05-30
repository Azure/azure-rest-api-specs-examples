using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Schedule/createOrUpdate.json
// this example is just showing the usage of "Schedules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningScheduleResource created on azure
// for more information of creating MachineLearningScheduleResource, please refer to the document of MachineLearningScheduleResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "my-aml-workspace";
string name = "string";
ResourceIdentifier machineLearningScheduleResourceId = MachineLearningScheduleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, name);
MachineLearningScheduleResource machineLearningSchedule = client.GetMachineLearningScheduleResource(machineLearningScheduleResourceId);

// invoke the operation
MachineLearningScheduleData data = new MachineLearningScheduleData(new MachineLearningScheduleProperties(new CronTrigger("string")
{
    EndTime = "string",
    StartTime = "string",
    TimeZone = "string",
}, new MachineLearningEndpointScheduleAction(BinaryData.FromObjectAsJson(new Dictionary<string, object>
{
    ["9965593e-526f-4b89-bb36-761138cf2794"] = null
})))
{
    DisplayName = "string",
    IsEnabled = false,
    Description = "string",
    Tags =
    {
    ["string"] = "string"
    },
    Properties =
    {
    ["string"] = "string"
    },
});
ArmOperation<MachineLearningScheduleResource> lro = await machineLearningSchedule.UpdateAsync(WaitUntil.Completed, data);
MachineLearningScheduleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineLearningScheduleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
