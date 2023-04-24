const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets NetworkRuleSet for a Namespace.
 *
 * @summary Gets NetworkRuleSet for a Namespace.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-10-01-preview/examples/NameSpaces/VirtualNetworkRule/EHNetworkRuleSetGet.json
 */
async function nameSpaceNetworkRuleSetGet() {
  const subscriptionId = process.env["EVENTHUB_SUBSCRIPTION_ID"] || "Subscription";
  const resourceGroupName = process.env["EVENTHUB_RESOURCE_GROUP"] || "ResourceGroup";
  const namespaceName = "sdk-Namespace-6019";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.namespaces.getNetworkRuleSet(resourceGroupName, namespaceName);
  console.log(result);
}
