const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the system topics under an Azure subscription.
 *
 * @summary List all the system topics under an Azure subscription.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/SystemTopics_ListBySubscription.json
 */
async function systemTopicsListBySubscription() {
  const subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.systemTopics.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

systemTopicsListBySubscription().catch(console.error);
