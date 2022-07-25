const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The services resource is the top-level resource that represents the Database Migration Service. The GET method retrieves information about a service instance.
 *
 * @summary The services resource is the top-level resource that represents the Database Migration Service. The GET method retrieves information about a service instance.
 * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2021-10-30-preview/examples/Services_Get.json
 */
async function servicesCreateOrUpdate() {
  const subscriptionId = "fc04246f-04c5-437e-ac5e-206a19e7193f";
  const groupName = "DmsSdkRg";
  const serviceName = "DmsSdkService";
  const credential = new DefaultAzureCredential();
  const client = new DataMigrationManagementClient(credential, subscriptionId);
  const result = await client.services.get(groupName, serviceName);
  console.log(result);
}

servicesCreateOrUpdate().catch(console.error);
