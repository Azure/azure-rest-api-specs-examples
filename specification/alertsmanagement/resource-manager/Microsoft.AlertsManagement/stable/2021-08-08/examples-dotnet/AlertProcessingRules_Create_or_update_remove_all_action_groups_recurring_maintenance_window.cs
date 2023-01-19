using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AlertsManagement;
using Azure.ResourceManager.AlertsManagement.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/stable/2021-08-08/examples/AlertProcessingRules_Create_or_update_remove_all_action_groups_recurring_maintenance_window.json
// this example is just showing the usage of "AlertProcessingRules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subId1";
string resourceGroupName = "alertscorrelationrg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this AlertProcessingRuleResource
AlertProcessingRuleCollection collection = resourceGroupResource.GetAlertProcessingRules();

// invoke the operation
string alertProcessingRuleName = "RemoveActionGroupsRecurringMaintenance";
AlertProcessingRuleData data = new AlertProcessingRuleData(new AzureLocation("Global"))
{
    Properties = new AlertProcessingRuleProperties(new string[]
{
"/subscriptions/subId1/resourceGroups/RGId1","/subscriptions/subId1/resourceGroups/RGId2"
}, new AlertProcessingRuleAction[]
{
new AlertProcessingRuleRemoveAllGroupsAction()
})
    {
        Conditions =
        {
        new AlertProcessingRuleCondition()
        {
        Field = AlertProcessingRuleField.TargetResourceType,
        Operator = AlertProcessingRuleOperator.EqualsValue,
        Values =
        {
        "microsoft.compute/virtualmachines"
        },
        }
        },
        Schedule = new AlertProcessingRuleSchedule()
        {
            TimeZone = "India Standard Time",
            Recurrences =
            {
            new AlertProcessingRuleWeeklyRecurrence(new AlertsManagementDayOfWeek[]
            {
            AlertsManagementDayOfWeek.Saturday,AlertsManagementDayOfWeek.Sunday
            })
            {
            StartOn = TimeSpan.Parse("22:00:00"),
            EndOn = TimeSpan.Parse("04:00:00"),
            }
            },
        },
        Description = "Remove all ActionGroups from all Vitual machine Alerts during the recurring maintenance",
        IsEnabled = true,
    },
    Tags =
    {
    },
};
ArmOperation<AlertProcessingRuleResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, alertProcessingRuleName, data);
AlertProcessingRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AlertProcessingRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
