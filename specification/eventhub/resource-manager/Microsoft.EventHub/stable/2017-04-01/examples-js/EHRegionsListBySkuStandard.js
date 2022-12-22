const { EventHubManagementClient } = require("@azure/arm-eventhub-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the available Regions for a given sku
 *
 * @summary Gets the available Regions for a given sku
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2017-04-01/examples/EHRegionsListBySkuStandard.json
 */
async function regionsListBySkuStandard() {
  const subscriptionId = process.env["EVENTHUB_SUBSCRIPTION_ID"] || "subscriptionid";
  const sku = "Standard";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.regions.listBySku(sku)) {
    resArray.push(item);
  }
  console.log(resArray);
}
