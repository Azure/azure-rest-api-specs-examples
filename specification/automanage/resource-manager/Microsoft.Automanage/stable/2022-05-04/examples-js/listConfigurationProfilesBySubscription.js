const { AutomanageClient } = require("@azure/arm-automanage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve a list of configuration profile within a subscription
 *
 * @summary Retrieve a list of configuration profile within a subscription
 * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listConfigurationProfilesBySubscription.json
 */
async function listConfigurationProfilesBySubscription() {
  const subscriptionId = "mySubscriptionId";
  const credential = new DefaultAzureCredential();
  const client = new AutomanageClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.configurationProfiles.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listConfigurationProfilesBySubscription().catch(console.error);
