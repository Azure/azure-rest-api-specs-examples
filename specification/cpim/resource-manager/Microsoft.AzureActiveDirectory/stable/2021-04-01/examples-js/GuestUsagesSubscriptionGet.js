const { ExternalIdentitiesConfigurationClient } = require("@azure/arm-azureadexternalidentities");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets Guest Usages resources under a subscription for the Microsoft.AzureActiveDirectory resource provider
 *
 * @summary Gets Guest Usages resources under a subscription for the Microsoft.AzureActiveDirectory resource provider
 * x-ms-original-file: specification/cpim/resource-manager/Microsoft.AzureActiveDirectory/stable/2021-04-01/examples/GuestUsagesSubscriptionGet.json
 */
async function guestUsagesSubscriptionList() {
  const subscriptionId = "c80fb759-c965-4c6a-9110-9b2b2d038882";
  const credential = new DefaultAzureCredential();
  const client = new ExternalIdentitiesConfigurationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.guestUsages.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

guestUsagesSubscriptionList().catch(console.error);
