using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Maintenance.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Maintenance;

// Generated from example definition: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2021-05-01/examples/ConfigurationAssignments_CreateOrUpdate.json
// this example is just showing the usage of "ConfigurationAssignments_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
string resourceGroupName = "examplerg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// invoke the operation
string providerName = "Microsoft.Compute";
string resourceType = "virtualMachineScaleSets";
string resourceName = "smdtest1";
string configurationAssignmentName = "workervmConfiguration";
MaintenanceConfigurationAssignmentData data = new MaintenanceConfigurationAssignmentData()
{
    MaintenanceConfigurationId = new ResourceIdentifier("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Maintenance/maintenanceConfigurations/configuration1"),
};
MaintenanceConfigurationAssignmentData result = await resourceGroupResource.CreateOrUpdateConfigurationAssignmentAsync(providerName, resourceType, resourceName, configurationAssignmentName, data);

Console.WriteLine($"Succeeded: {result}");
