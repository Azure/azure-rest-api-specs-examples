const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a service operation health status.
 *
 * @summary Gets a service operation health status.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/OperationsHealthListByLocation.json
 */
async function getManagementOperationsHealthInTheGivenLocation() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const locationName = "WestUS";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operationsHealthOperations.listByLocation(locationName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getManagementOperationsHealthInTheGivenLocation().catch(console.error);
