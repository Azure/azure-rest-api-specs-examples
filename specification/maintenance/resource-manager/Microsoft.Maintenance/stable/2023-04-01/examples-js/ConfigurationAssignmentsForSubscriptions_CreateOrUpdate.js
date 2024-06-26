const { MaintenanceManagementClient } = require("@azure/arm-maintenance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Register configuration for resource.
 *
 * @summary Register configuration for resource.
 * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2023-04-01/examples/ConfigurationAssignmentsForSubscriptions_CreateOrUpdate.json
 */
async function configurationAssignmentsForSubscriptionsCreateOrUpdate() {
  const subscriptionId =
    process.env["MAINTENANCE_SUBSCRIPTION_ID"] || "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const configurationAssignmentName = "workervmConfiguration";
  const configurationAssignment = {
    filter: {
      locations: ["Japan East", "UK South"],
      resourceGroups: ["RG1", "RG2"],
      resourceTypes: ["Microsoft.HybridCompute/machines", "Microsoft.Compute/virtualMachines"],
      tagSettings: {
        filterOperator: "Any",
        tags: {
          tag1: ["tag1Value1", "tag1Value2", "tag1Value3"],
          tag2: ["tag2Value1", "tag2Value2", "tag2Value3"],
        },
      },
    },
    maintenanceConfigurationId:
      "/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Maintenance/maintenanceConfigurations/configuration1",
  };
  const credential = new DefaultAzureCredential();
  const client = new MaintenanceManagementClient(credential, subscriptionId);
  const result = await client.configurationAssignmentsForSubscriptions.createOrUpdate(
    configurationAssignmentName,
    configurationAssignment
  );
  console.log(result);
}
