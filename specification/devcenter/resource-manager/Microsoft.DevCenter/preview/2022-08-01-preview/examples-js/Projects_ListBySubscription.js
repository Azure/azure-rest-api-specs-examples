const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all projects in the subscription.
 *
 * @summary Lists all projects in the subscription.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/Projects_ListBySubscription.json
 */
async function projectsListBySubscription() {
  const subscriptionId = "{subscriptionId}";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.projects.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

projectsListBySubscription().catch(console.error);
