const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a mobile network site.
 *
 * @summary Creates or updates a mobile network site.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SiteCreate.json
 */
async function createMobileNetworkSite() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const mobileNetworkName = "testMobileNetwork";
  const siteName = "testSite";
  const parameters = {
    location: "testLocation",
    networkFunctions: [
      {
        id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.HybridNetwork/networkFunctions/testNf",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.sites.beginCreateOrUpdateAndWait(
    resourceGroupName,
    mobileNetworkName,
    siteName,
    parameters
  );
  console.log(result);
}

createMobileNetworkSite().catch(console.error);
