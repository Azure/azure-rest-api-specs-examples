const { MaintenanceManagementClient } = require("@azure/arm-maintenance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List configurationAssignments for resource.
 *
 * @summary List configurationAssignments for resource.
 * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2023-04-01/examples/ConfigurationAssignments_ListParent.json
 */
async function configurationAssignmentsListParent() {
  const subscriptionId =
    process.env["MAINTENANCE_SUBSCRIPTION_ID"] || "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = process.env["MAINTENANCE_RESOURCE_GROUP"] || "examplerg";
  const providerName = "Microsoft.Compute";
  const resourceParentType = "virtualMachineScaleSets";
  const resourceParentName = "smdtest1";
  const resourceType = "virtualMachines";
  const resourceName = "smdtestvm1";
  const credential = new DefaultAzureCredential();
  const client = new MaintenanceManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.configurationAssignments.listParent(
    resourceGroupName,
    providerName,
    resourceParentType,
    resourceParentName,
    resourceType,
    resourceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
