const { SpringAppDiscoveryManagementClient } = require("@azure/arm-springappdiscovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List springbootservers resource by subscription
 *
 * @summary List springbootservers resource by subscription
 * x-ms-original-file: specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/springbootservers_ListBySubscription_MinimumSet_Gen.json
 */
async function springbootserversListBySubscriptionMinimumSetGen() {
  const subscriptionId = process.env["SPRINGAPPDISCOVERY_SUBSCRIPTION_ID"] || "etmdxomjncqvygm";
  const siteName = "hlkrzldhyobavtabgpubtjbhlslnjmsvkthwcfboriwyxndacjypzbj";
  const credential = new DefaultAzureCredential();
  const client = new SpringAppDiscoveryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.springbootservers.listBySubscription(siteName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
