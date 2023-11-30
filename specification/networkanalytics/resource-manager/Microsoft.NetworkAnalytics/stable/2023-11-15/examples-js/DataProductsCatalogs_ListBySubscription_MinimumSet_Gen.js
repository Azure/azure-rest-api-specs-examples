const { MicrosoftNetworkAnalytics } = require("@azure/arm-networkanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List data catalog by subscription.
 *
 * @summary List data catalog by subscription.
 * x-ms-original-file: specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/DataProductsCatalogs_ListBySubscription_MinimumSet_Gen.json
 */
async function dataProductsCatalogsListBySubscriptionMinimumSetGen() {
  const subscriptionId =
    process.env["NETWORKANALYTICS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-00000000000";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftNetworkAnalytics(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.dataProductsCatalogs.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
