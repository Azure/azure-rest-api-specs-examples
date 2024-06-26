const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The project resource is a nested resource representing a stored migration project. The PATCH method updates an existing project.
 *
 * @summary The project resource is a nested resource representing a stored migration project. The PATCH method updates an existing project.
 * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2021-10-30-preview/examples/Projects_Update.json
 */
async function projectsUpdate() {
  const subscriptionId = "fc04246f-04c5-437e-ac5e-206a19e7193f";
  const groupName = "DmsSdkRg";
  const serviceName = "DmsSdkService";
  const projectName = "DmsSdkProject";
  const parameters = {
    location: "southcentralus",
    sourcePlatform: "SQL",
    targetPlatform: "SQLDB",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataMigrationManagementClient(credential, subscriptionId);
  const result = await client.projects.update(groupName, serviceName, projectName, parameters);
  console.log(result);
}

projectsUpdate().catch(console.error);
