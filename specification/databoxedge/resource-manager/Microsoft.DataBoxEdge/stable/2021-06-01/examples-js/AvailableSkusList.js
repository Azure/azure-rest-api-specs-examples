const { DataBoxEdgeManagementClient } = require("@azure/arm-databoxedge");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the available Skus and information related to them.
 *
 * @summary List all the available Skus and information related to them.
 * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/AvailableSkusList.json
 */
async function availableSkus() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const credential = new DefaultAzureCredential();
  const client = new DataBoxEdgeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.availableSkus.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

availableSkus().catch(console.error);
