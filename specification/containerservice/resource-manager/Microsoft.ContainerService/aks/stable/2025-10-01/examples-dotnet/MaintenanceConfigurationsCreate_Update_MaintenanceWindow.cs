using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerService.Models;
using Azure.ResourceManager.ContainerService;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/MaintenanceConfigurationsCreate_Update_MaintenanceWindow.json
// this example is just showing the usage of "MaintenanceConfigurations_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerServiceMaintenanceConfigurationResource created on azure
// for more information of creating ContainerServiceMaintenanceConfigurationResource, please refer to the document of ContainerServiceMaintenanceConfigurationResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string resourceName = "clustername1";
string configName = "aksManagedAutoUpgradeSchedule";
ResourceIdentifier containerServiceMaintenanceConfigurationResourceId = ContainerServiceMaintenanceConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, configName);
ContainerServiceMaintenanceConfigurationResource containerServiceMaintenanceConfiguration = client.GetContainerServiceMaintenanceConfigurationResource(containerServiceMaintenanceConfigurationResourceId);

// invoke the operation
ContainerServiceMaintenanceConfigurationData data = new ContainerServiceMaintenanceConfigurationData
{
    MaintenanceWindow = new ContainerServiceMaintenanceWindow(new ContainerServiceMaintenanceSchedule
    {
        RelativeMonthly = new ContainerServiceMaintenanceRelativeMonthlySchedule(3, ContainerServiceMaintenanceRelativeMonthlyScheduleWeekIndex.First, ContainerServiceWeekDay.Monday),
    }, 10, "08:30")
    {
        UtcOffset = "+05:30",
        StartDate = "2023-01-01",
        NotAllowedDates = { new ContainerServiceDateSpan(DateTimeOffset.Parse("2023-02-18"), DateTimeOffset.Parse("2023-02-25")), new ContainerServiceDateSpan(DateTimeOffset.Parse("2023-12-23"), DateTimeOffset.Parse("2024-01-05")) },
    },
};
ArmOperation<ContainerServiceMaintenanceConfigurationResource> lro = await containerServiceMaintenanceConfiguration.UpdateAsync(WaitUntil.Completed, data);
ContainerServiceMaintenanceConfigurationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerServiceMaintenanceConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
