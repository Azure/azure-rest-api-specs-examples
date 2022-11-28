const { DesktopVirtualizationAPIClient } = require("@azure/arm-desktopvirtualization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Registration token of the host pool.
 *
 * @summary Registration token of the host pool.
 * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/HostPools_RetrieveRegistrationToken_Post.json
 */
async function hostPoolsRetrieveRegistrationTokenPost() {
  const subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
  const resourceGroupName = "resourceGroup1";
  const hostPoolName = "hostPool1";
  const credential = new DefaultAzureCredential();
  const client = new DesktopVirtualizationAPIClient(credential, subscriptionId);
  const result = await client.hostPools.retrieveRegistrationToken(resourceGroupName, hostPoolName);
  console.log(result);
}

hostPoolsRetrieveRegistrationTokenPost().catch(console.error);
