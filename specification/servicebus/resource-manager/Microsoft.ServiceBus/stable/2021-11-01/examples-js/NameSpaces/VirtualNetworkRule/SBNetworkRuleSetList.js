const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets list of NetworkRuleSet for a Namespace.
 *
 * @summary Gets list of NetworkRuleSet for a Namespace.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/VirtualNetworkRule/SBNetworkRuleSetList.json
 */
async function nameSpaceNetworkRuleSetList() {
  const subscriptionId = "Subscription";
  const resourceGroupName = "ResourceGroup";
  const namespaceName = "sdk-Namespace-6019";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.namespaces.listNetworkRuleSets(resourceGroupName, namespaceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

nameSpaceNetworkRuleSetList().catch(console.error);
