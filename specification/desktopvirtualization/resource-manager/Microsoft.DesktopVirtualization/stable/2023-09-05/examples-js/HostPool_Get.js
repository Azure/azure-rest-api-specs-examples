const { DesktopVirtualizationAPIClient } = require("@azure/arm-desktopvirtualization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a host pool.
 *
 * @summary Get a host pool.
 * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/HostPool_Get.json
 */
async function hostPoolGet() {
  const subscriptionId =
    process.env["DESKTOPVIRTUALIZATION_SUBSCRIPTION_ID"] || "daefabc0-95b4-48b3-b645-8a753a63c4fa";
  const resourceGroupName = process.env["DESKTOPVIRTUALIZATION_RESOURCE_GROUP"] || "resourceGroup1";
  const hostPoolName = "hostPool1";
  const credential = new DefaultAzureCredential();
  const client = new DesktopVirtualizationAPIClient(credential, subscriptionId);
  const result = await client.hostPools.get(resourceGroupName, hostPoolName);
  console.log(result);
}
