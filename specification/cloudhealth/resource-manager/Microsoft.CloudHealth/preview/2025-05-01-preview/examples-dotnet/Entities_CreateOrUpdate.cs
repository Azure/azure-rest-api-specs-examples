using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CloudHealth.Models;
using Azure.ResourceManager.CloudHealth;

// Generated from example definition: 2025-05-01-preview/Entities_CreateOrUpdate.json
// this example is just showing the usage of "Entity_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HealthModelEntityResource created on azure
// for more information of creating HealthModelEntityResource, please refer to the document of HealthModelEntityResource
string subscriptionId = "4980D7D5-4E07-47AD-AD34-E76C6BC9F061";
string resourceGroupName = "rgopenapi";
string healthModelName = "myHealthModel";
string entityName = "uszrxbdkxesdrxhmagmzywebgbjj";
ResourceIdentifier healthModelEntityResourceId = HealthModelEntityResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, healthModelName, entityName);
HealthModelEntityResource healthModelEntity = client.GetHealthModelEntityResource(healthModelEntityResourceId);

// invoke the operation
HealthModelEntityData data = new HealthModelEntityData
{
    Properties = new HealthModelEntityProperties
    {
        DisplayName = "My entity",
        CanvasPosition = new EntityCoordinates(14F, 13F),
        Icon = new EntityIcon("Custom")
        {
            CustomData = "rcitntvapruccrhtxmkqjphbxunkz",
        },
        HealthObjective = 62F,
        Impact = EntityImpact.Standard,
        Labels =
        {
        ["key1376"] = "ixfvzsfnpvkkbrce"
        },
        Signals = new EntitySignalGroup
        {
            AzureResource = new AzureResourceSignalGroup("B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX", new ResourceIdentifier("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1"))
            {
                SignalAssignments = { new EntitySignalAssignment(new string[] { "sigdef1" }) },
            },
            AzureLogAnalytics = new LogAnalyticsSignalGroup("B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX", new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.OperationalInsights/workspaces/myworkspace"))
            {
                SignalAssignments = { new EntitySignalAssignment(new string[] { "B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX" }) },
            },
            AzureMonitorWorkspace = new AzureMonitorWorkspaceSignalGroup("B3P1X3e-FZtZ-4Ak-2VLHGQ-4m4-05DE-XNW5zW3P-46XY-DC3SSX", new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.OperationalInsights/workspaces/myworkspace"))
            {
                SignalAssignments = { new EntitySignalAssignment(new string[] { "sigdef2" }), new EntitySignalAssignment(new string[] { "sigdef3" }) },
            },
            Dependencies = new DependenciesSignalGroup(DependenciesAggregationType.WorstOf),
        },
        Alerts = new EntityAlerts
        {
            Unhealthy = new EntityAlertConfiguration(EntityAlertSeverity.Sev1)
            {
                Description = "Alert description",
                ActionGroupIds = { new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Insights/actionGroups/myactiongroup") },
            },
            Degraded = new EntityAlertConfiguration(EntityAlertSeverity.Sev4)
            {
                Description = "Alert description",
                ActionGroupIds = { new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Insights/actionGroups/myactiongroup") },
            },
        },
    },
};
ArmOperation<HealthModelEntityResource> lro = await healthModelEntity.UpdateAsync(WaitUntil.Completed, data);
HealthModelEntityResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HealthModelEntityData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
