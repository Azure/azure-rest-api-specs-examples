const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List usages in cores for all skus used by a subscription in a given location, for a specific quota type.
 *
 * @summary List usages in cores for all skus used by a subscription in a given location, for a specific quota type.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/GetUsagesInLocation.json
 */
async function getUsagesInLocationForSubscription() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const location = "West US";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.getUsagesInLocation.list(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}
