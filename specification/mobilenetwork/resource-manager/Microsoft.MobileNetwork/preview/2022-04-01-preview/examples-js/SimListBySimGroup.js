const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the SIMs in a SIM group.
 *
 * @summary Gets all the SIMs in a SIM group.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/SimListBySimGroup.json
 */
async function listSiMSInASimGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const simGroupName = "testSimGroup";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sims.listBySimGroup(resourceGroupName, simGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listSiMSInASimGroup().catch(console.error);
