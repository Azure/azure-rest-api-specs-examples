const { ResourceConnectorManagementClient } = require("@azure/arm-resourceconnector");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all available Appliances operations.
 *
 * @summary Lists all available Appliances operations.
 * x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesListOperations.json
 */
async function listAppliancesOperations() {
  const credential = new DefaultAzureCredential();
  const client = new ResourceConnectorManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.appliances.listOperations()) {
    resArray.push(item);
  }
  console.log(resArray);
}
