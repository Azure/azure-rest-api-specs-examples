const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the specified mobile network site.
 *
 * @summary Gets information about the specified mobile network site.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SiteGet.json
 */
async function getMobileNetworkSite() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const mobileNetworkName = "testMobileNetwork";
  const siteName = "testSite";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.sites.get(resourceGroupName, mobileNetworkName, siteName);
  console.log(result);
}

getMobileNetworkSite().catch(console.error);
