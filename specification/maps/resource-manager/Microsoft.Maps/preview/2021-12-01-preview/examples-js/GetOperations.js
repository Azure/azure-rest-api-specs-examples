const { AzureMapsManagementClient } = require("@azure/arm-maps");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List operations available for the Maps Resource Provider
 *
 * @summary List operations available for the Maps Resource Provider
 * x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/preview/2021-12-01-preview/examples/GetOperations.json
 */
async function getOperations() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new AzureMapsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.maps.listOperations()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getOperations().catch(console.error);
