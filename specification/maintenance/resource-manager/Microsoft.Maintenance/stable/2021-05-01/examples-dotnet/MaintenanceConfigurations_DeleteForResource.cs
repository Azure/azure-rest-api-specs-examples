using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Maintenance.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Maintenance;

// Generated from example definition: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2021-05-01/examples/MaintenanceConfigurations_DeleteForResource.json
// this example is just showing the usage of "MaintenanceConfigurations_Delete" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this MaintenanceConfigurationResource created on azure
// for more information of creating MaintenanceConfigurationResource, please refer to the document of MaintenanceConfigurationResource
string subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
string resourceGroupName = "examplerg";
string resourceName = "example1";
ResourceIdentifier maintenanceConfigurationResourceId = MaintenanceConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
MaintenanceConfigurationResource maintenanceConfiguration = client.GetMaintenanceConfigurationResource(maintenanceConfigurationResourceId);

// invoke the operation
ArmOperation<MaintenanceConfigurationResource> lro = await maintenanceConfiguration.DeleteAsync(WaitUntil.Completed);
MaintenanceConfigurationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MaintenanceConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
