const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the authorization rules for a namespace.
 *
 * @summary Gets the authorization rules for a namespace.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/disasterRecoveryConfigs/SBAliasAuthorizationRuleListAll.json
 */
async function nameSpaceAuthorizationRuleListAll() {
  const subscriptionId = "exampleSubscriptionId";
  const resourceGroupName = "exampleResourceGroup";
  const namespaceName = "sdk-Namespace-9080";
  const alias = "sdk-DisasterRecovery-4047";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.disasterRecoveryConfigs.listAuthorizationRules(
    resourceGroupName,
    namespaceName,
    alias
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

nameSpaceAuthorizationRuleListAll().catch(console.error);
