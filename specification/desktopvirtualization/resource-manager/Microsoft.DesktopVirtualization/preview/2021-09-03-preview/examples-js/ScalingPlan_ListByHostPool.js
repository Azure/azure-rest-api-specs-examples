const { DesktopVirtualizationAPIClient } = require("@azure/arm-desktopvirtualization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List scaling plan associated with hostpool.
 *
 * @summary List scaling plan associated with hostpool.
 * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/ScalingPlan_ListByHostPool.json
 */
async function scalingPlanListByHostPool() {
  const subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
  const resourceGroupName = "resourceGroup1";
  const hostPoolName = "hostPool1";
  const credential = new DefaultAzureCredential();
  const client = new DesktopVirtualizationAPIClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.scalingPlans.listByHostPool(resourceGroupName, hostPoolName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

scalingPlanListByHostPool().catch(console.error);
