const { AzureMapsManagementClient } = require("@azure/arm-maps");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List operations available for the Maps Resource Provider
 *
 * @summary List operations available for the Maps Resource Provider
 * x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/examples/GetOperations.json
 */
async function getOperations() {
  const credential = new DefaultAzureCredential();
  const client = new AzureMapsManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.maps.listOperations()) {
    resArray.push(item);
  }
  console.log(resArray);
}
