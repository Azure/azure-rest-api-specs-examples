const { DesktopVirtualizationAPIClient } = require("@azure/arm-desktopvirtualization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get an application group.
 *
 * @summary Get an application group.
 * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/ApplicationGroup_Get.json
 */
async function applicationGroupGet() {
  const subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
  const resourceGroupName = "resourceGroup1";
  const applicationGroupName = "applicationGroup1";
  const credential = new DefaultAzureCredential();
  const client = new DesktopVirtualizationAPIClient(credential, subscriptionId);
  const result = await client.applicationGroups.get(resourceGroupName, applicationGroupName);
  console.log(result);
}

applicationGroupGet().catch(console.error);
