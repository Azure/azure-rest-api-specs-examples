using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Maintenance;
using Azure.ResourceManager.Maintenance.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2021-05-01/examples/MaintenanceConfigurations_DeleteForResource.json
// this example is just showing the usage of "MaintenanceConfigurations_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MaintenanceConfigurationResource created on azure
// for more information of creating MaintenanceConfigurationResource, please refer to the document of MaintenanceConfigurationResource
string subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
string resourceGroupName = "examplerg";
string resourceName = "example1";
ResourceIdentifier maintenanceConfigurationResourceId = MaintenanceConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
MaintenanceConfigurationResource maintenanceConfiguration = client.GetMaintenanceConfigurationResource(maintenanceConfigurationResourceId);

// invoke the operation
await maintenanceConfiguration.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
