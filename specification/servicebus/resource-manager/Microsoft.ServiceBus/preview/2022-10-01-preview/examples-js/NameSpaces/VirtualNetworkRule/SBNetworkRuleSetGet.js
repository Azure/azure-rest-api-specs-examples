const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets NetworkRuleSet for a Namespace.
 *
 * @summary Gets NetworkRuleSet for a Namespace.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/NameSpaces/VirtualNetworkRule/SBNetworkRuleSetGet.json
 */
async function nameSpaceNetworkRuleSetGet() {
  const subscriptionId = process.env["SERVICEBUS_SUBSCRIPTION_ID"] || "Subscription";
  const resourceGroupName = process.env["SERVICEBUS_RESOURCE_GROUP"] || "ResourceGroup";
  const namespaceName = "sdk-Namespace-6019";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.namespaces.getNetworkRuleSet(resourceGroupName, namespaceName);
  console.log(result);
}
