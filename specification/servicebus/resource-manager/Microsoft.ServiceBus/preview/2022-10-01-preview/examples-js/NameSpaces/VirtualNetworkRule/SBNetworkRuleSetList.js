const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets list of NetworkRuleSet for a Namespace.
 *
 * @summary Gets list of NetworkRuleSet for a Namespace.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/NameSpaces/VirtualNetworkRule/SBNetworkRuleSetList.json
 */
async function nameSpaceNetworkRuleSetList() {
  const subscriptionId = process.env["SERVICEBUS_SUBSCRIPTION_ID"] || "Subscription";
  const resourceGroupName = process.env["SERVICEBUS_RESOURCE_GROUP"] || "ResourceGroup";
  const namespaceName = "sdk-Namespace-6019";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.namespaces.listNetworkRuleSets(resourceGroupName, namespaceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
