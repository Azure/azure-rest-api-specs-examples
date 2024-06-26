const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The service tasks resource is a nested, proxy-only resource representing work performed by a DMS instance. The PUT method creates a new service task or updates an existing one, although since service tasks have no mutable custom properties, there is little reason to update an existing one.
 *
 * @summary The service tasks resource is a nested, proxy-only resource representing work performed by a DMS instance. The PUT method creates a new service task or updates an existing one, although since service tasks have no mutable custom properties, there is little reason to update an existing one.
 * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2021-10-30-preview/examples/ServiceTasks_CreateOrUpdate.json
 */
async function tasksCreateOrUpdate() {
  const subscriptionId = "fc04246f-04c5-437e-ac5e-206a19e7193f";
  const groupName = "DmsSdkRg";
  const serviceName = "DmsSdkService";
  const taskName = "DmsSdkTask";
  const parameters = {
    properties: {
      input: { serverVersion: "NA" },
      taskType: "Service.Check.OCI",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataMigrationManagementClient(credential, subscriptionId);
  const result = await client.serviceTasks.createOrUpdate(
    groupName,
    serviceName,
    taskName,
    parameters
  );
  console.log(result);
}

tasksCreateOrUpdate().catch(console.error);
