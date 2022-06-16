const { ApplicationClient } = require("@azure/arm-managedapplications");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the applications within a subscription.
 *
 * @summary Gets all the applications within a subscription.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/listApplicationsBySubscription.json
 */
async function listsApplicationsBySubscription() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.applications.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listsApplicationsBySubscription().catch(console.error);
