using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Maintenance.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Maintenance;

// Generated from example definition: specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2023-10-01-preview/examples/ConfigurationAssignmentsForSubscriptions_UpdateForResource.json
// this example is just showing the usage of "ConfigurationAssignmentsForSubscriptions_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
string configurationAssignmentName = "workervmConfiguration";
MaintenanceConfigurationAssignmentData data = new MaintenanceConfigurationAssignmentData
{
    MaintenanceConfigurationId = new ResourceIdentifier("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Maintenance/maintenanceConfigurations/configuration1"),
    Filter = new MaintenanceConfigurationAssignmentFilter
    {
        ResourceTypes = { new ResourceType("Microsoft.HybridCompute/machines"), new ResourceType("Microsoft.Compute/virtualMachines") },
        ResourceGroups = { "RG1", "RG2" },
        Locations = { new AzureLocation("Japan East"), new AzureLocation("UK South") },
        TagSettings = new VmTagSettings
        {
            Tags =
            {
            ["tag1"] = new string[]{"tag1Value1", "tag1Value2", "tag1Value3"},
            ["tag2"] = new string[]{"tag2Value1", "tag2Value2", "tag2Value3"}
            },
            FilterOperator = VmTagOperator.Any,
        },
    },
};
MaintenanceConfigurationAssignmentData result = await subscriptionResource.UpdateConfigurationAssignmentBySubscriptionAsync(configurationAssignmentName, data);

Console.WriteLine($"Succeeded: {result}");
