using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AlertsManagement;
using Azure.ResourceManager.AlertsManagement.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/stable/2021-08-08/examples/AlertProcessingRules_Create_or_update_add_two_action_groups_all_Sev0_Sev1_two_resource_groups.json
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
string alertProcessingRuleName = "AddActionGroupsBySeverity";
AlertProcessingRuleData data = new AlertProcessingRuleData(new AzureLocation("Global"))
{
    Properties = new AlertProcessingRuleProperties(new string[]
{
"/subscriptions/subId1/resourceGroups/RGId1","/subscriptions/subId1/resourceGroups/RGId2"
}, new AlertProcessingRuleAction[]
{
new AlertProcessingRuleAddGroupsAction(new ResourceIdentifier[]
{
new ResourceIdentifier("/subscriptions/subId1/resourcegroups/RGId1/providers/microsoft.insights/actiongroups/AGId1"),new ResourceIdentifier("/subscriptions/subId1/resourcegroups/RGId1/providers/microsoft.insights/actiongroups/AGId2")
})
})
    {
        Conditions =
        {
        new AlertProcessingRuleCondition()
        {
        Field = AlertProcessingRuleField.Severity,
        Operator = AlertProcessingRuleOperator.EqualsValue,
        Values =
        {
        "sev0","sev1"
        },
        }
        },
        Description = "Add AGId1 and AGId2 to all Sev0 and Sev1 alerts in these resourceGroups",
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
