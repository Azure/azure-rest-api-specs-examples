const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets NetworkRuleSet for a Namespace.
 *
 * @summary Gets NetworkRuleSet for a Namespace.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/VirtualNetworkRule/SBNetworkRuleSetGet.json
 */
async function nameSpaceNetworkRuleSetGet() {
  const subscriptionId = "Subscription";
  const resourceGroupName = "ResourceGroup";
  const namespaceName = "sdk-Namespace-6019";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.namespaces.getNetworkRuleSet(resourceGroupName, namespaceName);
  console.log(result);
}

nameSpaceNetworkRuleSetGet().catch(console.error);
