const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the authorization rules for a namespace.
 *
 * @summary Gets the authorization rules for a namespace.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/SBNameSpaceAuthorizationRuleListAll.json
 */
async function nameSpaceAuthorizationRuleListAll() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = "ArunMonocle";
  const namespaceName = "sdk-Namespace-6914";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.namespaces.listAuthorizationRules(
    resourceGroupName,
    namespaceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

nameSpaceAuthorizationRuleListAll().catch(console.error);
