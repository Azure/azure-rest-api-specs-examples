const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List disk images available for custom image creation.
 *
 * @summary List disk images available for custom image creation.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_ListVhds.json
 */
async function labsListVhds() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const name = "{labName}";
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.labs.listVhds(resourceGroupName, name)) {
    resArray.push(item);
  }
  console.log(resArray);
}

labsListVhds().catch(console.error);
