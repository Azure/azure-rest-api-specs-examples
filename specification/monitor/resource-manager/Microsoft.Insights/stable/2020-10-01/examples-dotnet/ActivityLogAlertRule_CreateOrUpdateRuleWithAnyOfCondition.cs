using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/ActivityLogAlertRule_CreateOrUpdateRuleWithAnyOfCondition.json
// this example is just showing the usage of "ActivityLogAlerts_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "187f412d-1758-44d9-b052-169e2564721d";
string resourceGroupName = "MyResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ActivityLogAlertResource
ActivityLogAlertCollection collection = resourceGroupResource.GetActivityLogAlerts();

// invoke the operation
string activityLogAlertName = "SampleActivityLogAlertRuleWithAnyOfCondition";
ActivityLogAlertData data = new ActivityLogAlertData(new AzureLocation("Global"))
{
    Scopes =
    {
    "subscriptions/187f412d-1758-44d9-b052-169e2564721d"
    },
    ConditionAllOf =
    {
    new ActivityLogAlertAnyOfOrLeafCondition()
    {
    Field = "category",
    EqualsValue = "ServiceHealth",
    },new ActivityLogAlertAnyOfOrLeafCondition()
    {
    AnyOf =
    {
    new AlertRuleLeafCondition()
    {
    Field = "properties.incidentType",
    EqualsValue = "Incident",
    },new AlertRuleLeafCondition()
    {
    Field = "properties.incidentType",
    EqualsValue = "Maintenance",
    }
    },
    }
    },
    ActionsActionGroups =
    {
    new ActivityLogAlertActionGroup(new ResourceIdentifier("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/actionGroups/SampleActionGroup"))
    {
    WebhookProperties =
    {
    ["sampleWebhookProperty"] = "SamplePropertyValue",
    },
    }
    },
    IsEnabled = true,
    Description = "Description of sample Activity Log Alert rule with 'anyOf' condition.",
    Tags =
    {
    },
};
ArmOperation<ActivityLogAlertResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, activityLogAlertName, data);
ActivityLogAlertResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ActivityLogAlertData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
