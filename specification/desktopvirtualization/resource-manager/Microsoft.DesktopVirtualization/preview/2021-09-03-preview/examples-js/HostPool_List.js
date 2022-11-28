const { DesktopVirtualizationAPIClient } = require("@azure/arm-desktopvirtualization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List hostPools in subscription.
 *
 * @summary List hostPools in subscription.
 * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/HostPool_List.json
 */
async function hostPoolList() {
  const subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
  const credential = new DefaultAzureCredential();
  const client = new DesktopVirtualizationAPIClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.hostPools.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

hostPoolList().catch(console.error);
