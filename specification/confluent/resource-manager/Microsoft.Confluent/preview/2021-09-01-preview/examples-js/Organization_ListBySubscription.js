const { ConfluentManagementClient } = require("@azure/arm-confluent");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all organizations under the specified subscription.
 *
 * @summary List all organizations under the specified subscription.
 * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/preview/2021-09-01-preview/examples/Organization_ListBySubscription.json
 */
async function organizationListBySubscription() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new ConfluentManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.organization.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

organizationListBySubscription().catch(console.error);
