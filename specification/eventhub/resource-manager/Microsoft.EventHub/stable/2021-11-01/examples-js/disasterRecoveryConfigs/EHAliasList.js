const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all Alias(Disaster Recovery configurations)
 *
 * @summary Gets all Alias(Disaster Recovery configurations)
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/disasterRecoveryConfigs/EHAliasList.json
 */
async function ehAliasList() {
  const subscriptionId = "exampleSubscriptionId";
  const resourceGroupName = "exampleResourceGroup";
  const namespaceName = "sdk-Namespace-8859";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.disasterRecoveryConfigs.list(resourceGroupName, namespaceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

ehAliasList().catch(console.error);
