const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List service fabrics in a given user profile.
 *
 * @summary List service fabrics in a given user profile.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabrics_List.json
 */
async function serviceFabricsList() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const userName = "{userName}";
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.serviceFabrics.list(resourceGroupName, labName, userName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

serviceFabricsList().catch(console.error);
