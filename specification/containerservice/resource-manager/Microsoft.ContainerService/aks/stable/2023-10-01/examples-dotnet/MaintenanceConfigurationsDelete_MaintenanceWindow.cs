using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerService.Models;
using Azure.ResourceManager.ContainerService;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-10-01/examples/MaintenanceConfigurationsDelete_MaintenanceWindow.json
// this example is just showing the usage of "MaintenanceConfigurations_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerServiceMaintenanceConfigurationResource created on azure
// for more information of creating ContainerServiceMaintenanceConfigurationResource, please refer to the document of ContainerServiceMaintenanceConfigurationResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string resourceName = "clustername1";
string configName = "aksManagedNodeOSUpgradeSchedule";
ResourceIdentifier containerServiceMaintenanceConfigurationResourceId = ContainerServiceMaintenanceConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, configName);
ContainerServiceMaintenanceConfigurationResource containerServiceMaintenanceConfiguration = client.GetContainerServiceMaintenanceConfigurationResource(containerServiceMaintenanceConfigurationResourceId);

// invoke the operation
await containerServiceMaintenanceConfiguration.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
