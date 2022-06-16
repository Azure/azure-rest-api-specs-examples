const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an authorization rule for a namespace by rule name.
 *
 * @summary Gets an authorization rule for a namespace by rule name.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/SBNameSpaceAuthorizationRuleGet.json
 */
async function nameSpaceAuthorizationRuleGet() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = "ArunMonocle";
  const namespaceName = "sdk-Namespace-6914";
  const authorizationRuleName = "sdk-AuthRules-1788";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.namespaces.getAuthorizationRule(
    resourceGroupName,
    namespaceName,
    authorizationRuleName
  );
  console.log(result);
}

nameSpaceAuthorizationRuleGet().catch(console.error);
