using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Maintenance.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Maintenance;

// Generated from example definition: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2021-05-01/examples/MaintenanceConfigurations_UpdateForResource.json
// this example is just showing the usage of "MaintenanceConfigurations_Update" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this MaintenanceConfigurationResource created on azure
// for more information of creating MaintenanceConfigurationResource, please refer to the document of MaintenanceConfigurationResource
string subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
string resourceGroupName = "examplerg";
string resourceName = "configuration1";
ResourceIdentifier maintenanceConfigurationResourceId = MaintenanceConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
MaintenanceConfigurationResource maintenanceConfiguration = client.GetMaintenanceConfigurationResource(maintenanceConfigurationResourceId);

// invoke the operation
MaintenanceConfigurationData data = new MaintenanceConfigurationData(new AzureLocation("westus2"))
{
    Namespace = "Microsoft.Maintenance",
    MaintenanceScope = MaintenanceScope.Host,
    Visibility = MaintenanceConfigurationVisibility.Custom,
    StartOn = DateTimeOffset.Parse("2020-04-30 08:00"),
    ExpireOn = DateTimeOffset.Parse("9999-12-31 00:00"),
    Duration = TimeSpan.Parse("05:00"),
    TimeZone = "Pacific Standard Time",
    RecurEvery = "Month Third Sunday",
};
MaintenanceConfigurationResource result = await maintenanceConfiguration.UpdateAsync(data);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MaintenanceConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
