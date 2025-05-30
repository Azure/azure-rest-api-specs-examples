using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DevTestLabs.Models;
using Azure.ResourceManager.DevTestLabs;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Costs_CreateOrUpdate.json
// this example is just showing the usage of "Costs_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevTestLabCostResource created on azure
// for more information of creating DevTestLabCostResource, please refer to the document of DevTestLabCostResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "resourceGroupName";
string labName = "{labName}";
string name = "targetCost";
ResourceIdentifier devTestLabCostResourceId = DevTestLabCostResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, labName, name);
DevTestLabCostResource devTestLabCost = client.GetDevTestLabCostResource(devTestLabCostResourceId);

// invoke the operation
DevTestLabCostData data = new DevTestLabCostData(default)
{
    TargetCost = new DevTestLabTargetCost
    {
        Status = DevTestLabTargetCostStatus.Enabled,
        Target = 100,
        CostThresholds = {new DevTestLabCostThreshold
        {
        ThresholdId = "00000000-0000-0000-0000-000000000001",
        ThresholdValue = 25,
        DisplayOnChart = DevTestLabCostThresholdStatus.Disabled,
        SendNotificationWhenExceeded = DevTestLabCostThresholdStatus.Disabled,
        }, new DevTestLabCostThreshold
        {
        ThresholdId = "00000000-0000-0000-0000-000000000002",
        ThresholdValue = 50,
        DisplayOnChart = DevTestLabCostThresholdStatus.Enabled,
        SendNotificationWhenExceeded = DevTestLabCostThresholdStatus.Enabled,
        }, new DevTestLabCostThreshold
        {
        ThresholdId = "00000000-0000-0000-0000-000000000003",
        ThresholdValue = 75,
        DisplayOnChart = DevTestLabCostThresholdStatus.Disabled,
        SendNotificationWhenExceeded = DevTestLabCostThresholdStatus.Disabled,
        }, new DevTestLabCostThreshold
        {
        ThresholdId = "00000000-0000-0000-0000-000000000004",
        ThresholdValue = 100,
        DisplayOnChart = DevTestLabCostThresholdStatus.Disabled,
        SendNotificationWhenExceeded = DevTestLabCostThresholdStatus.Disabled,
        }, new DevTestLabCostThreshold
        {
        ThresholdId = "00000000-0000-0000-0000-000000000005",
        ThresholdValue = 125,
        DisplayOnChart = DevTestLabCostThresholdStatus.Disabled,
        SendNotificationWhenExceeded = DevTestLabCostThresholdStatus.Disabled,
        }},
        CycleStartOn = DateTimeOffset.Parse("2020-12-01T00:00:00.000Z"),
        CycleEndOn = DateTimeOffset.Parse("2020-12-31T00:00:00.000Z"),
        CycleType = DevTestLabReportingCycleType.CalendarMonth,
    },
    CurrencyCode = "USD",
    StartOn = DateTimeOffset.Parse("2020-12-01T00:00:00Z"),
    EndOn = DateTimeOffset.Parse("2020-12-31T23:59:59Z"),
};
ArmOperation<DevTestLabCostResource> lro = await devTestLabCost.UpdateAsync(WaitUntil.Completed, data);
DevTestLabCostResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DevTestLabCostData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
