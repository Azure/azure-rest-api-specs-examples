const { SpringAppDiscoveryManagementClient } = require("@azure/arm-springappdiscovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List springbootapps resource by subscription
 *
 * @summary List springbootapps resource by subscription
 * x-ms-original-file: specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/springbootapps_ListBySubscription_MaximumSet_Gen.json
 */
async function springbootappsListBySubscriptionMaximumSetGen() {
  const subscriptionId =
    process.env["SPRINGAPPDISCOVERY_SUBSCRIPTION_ID"] || "jnetwlorzmxpxmcucorv";
  const siteName = "pdfosfhtemfsaglvwjdyqlyeipucrd";
  const credential = new DefaultAzureCredential();
  const client = new SpringAppDiscoveryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.springbootapps.listBySubscription(siteName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
