const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified mobile network site.
 *
 * @summary Deletes the specified mobile network site.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/SiteDelete.json
 */
async function deleteMobileNetworkSite() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const mobileNetworkName = "testMobileNetwork";
  const siteName = "testSite";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.sites.beginDeleteAndWait(
    resourceGroupName,
    mobileNetworkName,
    siteName
  );
  console.log(result);
}

deleteMobileNetworkSite().catch(console.error);
