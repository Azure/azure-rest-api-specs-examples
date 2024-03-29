const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The services resource is the top-level resource that represents the Database Migration Service. The skus action returns the list of SKUs that a service resource can be updated to.
 *
 * @summary The services resource is the top-level resource that represents the Database Migration Service. The skus action returns the list of SKUs that a service resource can be updated to.
 * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2021-10-30-preview/examples/Services_ListSkus.json
 */
async function servicesListSkus() {
  const subscriptionId = "fc04246f-04c5-437e-ac5e-206a19e7193f";
  const groupName = "DmsSdkRg";
  const serviceName = "DmsSdkService";
  const credential = new DefaultAzureCredential();
  const client = new DataMigrationManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.services.listSkus(groupName, serviceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

servicesListSkus().catch(console.error);
