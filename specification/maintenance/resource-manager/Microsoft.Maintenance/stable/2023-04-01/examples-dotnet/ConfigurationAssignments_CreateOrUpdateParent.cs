using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Maintenance;
using Azure.ResourceManager.Maintenance.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2023-04-01/examples/ConfigurationAssignments_CreateOrUpdateParent.json
// this example is just showing the usage of "ConfigurationAssignments_CreateOrUpdateParent" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
string resourceGroupName = "examplerg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// invoke the operation
ResourceGroupResourceCreateOrUpdateConfigurationAssignmentByParentOptions options = new ResourceGroupResourceCreateOrUpdateConfigurationAssignmentByParentOptions(providerName: "Microsoft.Compute", resourceParentType: "virtualMachineScaleSets", resourceParentName: "smdtest1", resourceType: "virtualMachines", resourceName: "smdvm1", configurationAssignmentName: "workervmPolicy", data: new MaintenanceConfigurationAssignmentData()
{
    MaintenanceConfigurationId = new ResourceIdentifier("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Maintenance/maintenanceConfigurations/policy1"),
})
{ };
MaintenanceConfigurationAssignmentData result = await resourceGroupResource.CreateOrUpdateConfigurationAssignmentByParentAsync(options);

Console.WriteLine($"Succeeded: {result}");
