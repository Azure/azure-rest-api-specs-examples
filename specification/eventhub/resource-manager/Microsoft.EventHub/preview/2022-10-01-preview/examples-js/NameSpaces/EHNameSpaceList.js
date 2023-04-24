const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the available Namespaces within a subscription, irrespective of the resource groups.
 *
 * @summary Lists all the available Namespaces within a subscription, irrespective of the resource groups.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-10-01-preview/examples/NameSpaces/EHNameSpaceList.json
 */
async function namespacesListBySubscription() {
  const subscriptionId = process.env["EVENTHUB_SUBSCRIPTION_ID"] || "SampleSubscription";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.namespaces.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
