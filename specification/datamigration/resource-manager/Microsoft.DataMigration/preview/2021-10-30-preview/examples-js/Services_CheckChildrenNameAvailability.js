const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This method checks whether a proposed nested resource name is valid and available.
 *
 * @summary This method checks whether a proposed nested resource name is valid and available.
 * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2021-10-30-preview/examples/Services_CheckChildrenNameAvailability.json
 */
async function servicesCheckChildrenNameAvailability() {
  const subscriptionId = "fc04246f-04c5-437e-ac5e-206a19e7193f";
  const groupName = "DmsSdkRg";
  const serviceName = "DmsSdkService";
  const parameters = { name: "Task1", type: "tasks" };
  const credential = new DefaultAzureCredential();
  const client = new DataMigrationManagementClient(credential, subscriptionId);
  const result = await client.services.checkChildrenNameAvailability(
    groupName,
    serviceName,
    parameters
  );
  console.log(result);
}

servicesCheckChildrenNameAvailability().catch(console.error);
